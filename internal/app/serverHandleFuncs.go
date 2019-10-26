package apiserver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (s *server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()

	decoder := json.NewDecoder(r.Body)
	user := new(model.User)
	err := decoder.Decode(user)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	s.userType = user.UserType
	if s.userType != userFreelancer && s.userType != userCustomer {
		s.userType = userFreelancer
	}
	fmt.Println(user)

	if err := user.Validate(); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	if err := user.BeforeCreate(); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.store.Mu.Lock()
	err = s.store.User().Create(user)
	s.store.Mu.Unlock()

	log.Println("USER ID: ", user.ID)

	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, err)
		return
	}

	f := model.Freelancer{
		AccountId:        user.ID,
		RegistrationDate: time.Now(),
	}

	s.store.Mu.Lock()
	err = s.store.Freelancer().Create(&f)
	s.store.Mu.Unlock()

	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, err)
		return
	}

	m := model.HireManager{
		AccountID:        user.ID,
		RegistrationDate: time.Now(),
	}

	s.store.Mu.Lock()
	err = s.store.Manager().Create(&m)
	s.store.Mu.Unlock()

	if err != nil {
		s.respond(w, r, http.StatusInternalServerError, err)
		return
	}

	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	session.Values["user_id"] = user.ID
	session.Values["user_type"] = s.userType

	if err := s.sessionStore.Save(r, w, session); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	// TODO: rename cookie2
	cookie := http.Cookie{Name: userTypeCookieName, Value: s.userType}
	cookie2 := http.Cookie{Name: hireManagerIdCookieName, Value: strconv.Itoa(1)} // m.Id
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)

	s.respond(w, r, http.StatusCreated, user)
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		s.store.Mu.Lock()
		u, err := s.store.User().Find(id.(int))
		s.store.Mu.Unlock()

		if err != nil {
			s.error(w,r,http.StatusNotFound, err)
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, &u)))
	})
}

func (s *server) HandleSessionCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)
	w.Header().Set("Content-Type", "application/json")


	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()
	decoder := json.NewDecoder(r.Body)
	user := new(model.User)
	err := decoder.Decode(user)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	log.Println("input user:", user)

	s.store.Mu.Lock()
	u, err := s.store.User().FindByEmail(user.Email)
	s.store.Mu.Unlock()

	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}

	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	if s.userType == "" {
		s.userType = userFreelancer
	}

	session.Values["user_id"] = u.ID
	session.Values["user_type"] = s.userType

	if err := s.sessionStore.Save(r, w, session); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, struct {}{})
}

func (s *server) HandleLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)


	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, err)
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		s.error(w, r, http.StatusExpectationFailed, err)
	}
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleSetUserType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		UserType string `json:"type"`
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()
	decoder := json.NewDecoder(r.Body)
	newInput := new(Input)
	err := decoder.Decode(newInput)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	if newInput.UserType != userCustomer && newInput.UserType != userFreelancer {
		s.error(w, r, http.StatusBadRequest, errors.New("user type may be only customer or freelancer"))
	}

	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		s.error(w, r,http.StatusUnauthorized, err)
		return
	}

	// TODO: delete cookie or not
	session.Values["user_type"] = newInput.UserType

	s.userType = newInput.UserType

	err = session.Save(r, w)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
}

func (s *server) HandleShowProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}
	s.respond(w, r, http.StatusOK, user)
}

func (s *server) HandleEditProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)
	w.Header().Set("Content-Type", "application/json")


	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}

	// TODO: validate edited user

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(user)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		s.error(w, r, http.StatusBadRequest, errors.New("invalid format of data"))
		return
	}

	s.store.Mu.Lock()
	err = s.store.User().Edit(user)
	s.store.Mu.Unlock()

	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	log.Println("edited profile: ", user)
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleEditPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type BodyPassword struct {
		Password string
		NewPassword string
		NewPasswordConfirmation string
	}

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()

	bodyPassword := new(BodyPassword)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(bodyPassword)

	if bodyPassword.NewPassword != bodyPassword.NewPasswordConfirmation {
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("new Passwords are different"))
		return
	}

	if !user.ComparePassword(bodyPassword.Password) {
		s.error(w, r, http.StatusBadRequest, fmt.Errorf("wrong password"))
		return
	}

	newEncryptPassword, err := model.EncryptString(bodyPassword.NewPasswordConfirmation)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("error in updating password"))
		return
	}
	user.EncryptPassword = newEncryptPassword

	s.store.Mu.Lock()
	err = s.store.User().Edit(user)
	s.store.Mu.Unlock()

	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) GetUserFromRequest(r *http.Request) (*model.User, error, int) {
	session, err := s.sessionStore.Get(r, sessionName)
	if err != nil {
		sendErr := fmt.Errorf("user isn't authorized")
		return nil, sendErr, http.StatusUnauthorized
	}

	uidInterface := session.Values["user_id"]
	uid := uidInterface.(int)

	s.store.Mu.Lock()
	user, err := s.store.User().Find(uid)
	s.store.Mu.Unlock()

	if err != nil {
		sendErr := fmt.Errorf("can't find user with id:" + strconv.Itoa(int(uid)))
		return nil, sendErr, http.StatusBadRequest
	}
	return user, nil, http.StatusOK
}

// TODO: fix after creating Notifications table
func (s *server) HandleEditNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, sendErr, codeStatus := s.GetUserFromRequest(r)
	if sendErr != nil {
		s.error(w, r, codeStatus, sendErr)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()

	userNotification := s.usersdb.Notifications[user.ID]
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(userNotification)
	log.Println(user)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		sendErr := fmt.Errorf("invalid format of data")
		s.error(w, r, http.StatusBadRequest, sendErr)
		return
	}
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleUploadAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("error retrieving the file"))
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, errors.New("error retrieving the file"))
		return
	}
	defer file.Close()

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}

	image := bytes.NewBuffer(nil)
	_, err = io.Copy(image, file)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	user.Avatar = image.Bytes()

	s.store.Mu.Lock()
	err = s.store.User().Edit(user)
	s.usersdb.Mu.Unlock()

	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleDownloadAvatar(w http.ResponseWriter, r *http.Request) {

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}


	if user.Avatar != nil {
		image := user.Avatar
		if _, err := w.Write(image); err != nil {
			s.error(w,r,http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Content-Length", strconv.Itoa(len(image)))
		w.Header().Set("Content-Type", "multipart/form-data")

	} else {
		filename := "internal/store/avatars/default.png"

		var openFile *os.File
		openFile, err = os.Open(filename)
		if err != nil {
			s.error(w, r, http.StatusNotFound, errors.New("cant open file"))
			return
		}

		defer func(){
			if err := openFile.Close(); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
			}
		}()

		fileHeader := make([]byte, 100000) // max image size!!!
		if _, err := openFile.Read(fileHeader); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		fileContentType := http.DetectContentType(fileHeader)
		fileStat, _ := openFile.Stat()
		fileSize := strconv.FormatInt(fileStat.Size(), 10)

		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", fileContentType)
		w.Header().Set("Content-Length", fileSize)

		if _, err := openFile.Seek(0, 0); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if _, err := io.Copy(w, openFile); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}
	s.respond(w,r,http.StatusOK, struct{}{})
}

func (s *server) HandleRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}

	hireManager, err := s.store.Manager().Find(user.ID)
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
	}

	// TODO: rewrite after Roles and Companies db interfaces realization
	company := s.usersdb.Companies[hireManager.ID]
	var roles []*model.Role
	clientRole := &model.Role{
		Role:   "client",
		Label:  company.CompanyName,
		Avatar: "/default.png",
	}
	freelanceRole := &model.Role{
		Role:   "freelancer",
		Label:  user.FirstName + " " + user.SecondName,
		Avatar: "/default.png",
	}
	roles = append(roles, clientRole)
	roles = append(roles, freelanceRole)
	s.respond(w, r, http.StatusOK, roles)
}

func (s *server) HandleGetAuthHistory(w http.ResponseWriter, r *http.Request) {
	// TODO: get auth history
}

func (s *server) HandleGetSecQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: get sec question
}

func (s *server) HandleEditSecQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: edit sec question
}

func (s *server) HandleCheckSecQuestion(w http.ResponseWriter, r *http.Request) {
	// TODO: check seq question
}

func (s *server)CORSMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions{
			w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Lol")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)
			s.respond(w , r , http.StatusOK, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *server) HandleCreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()

	decoder := json.NewDecoder(r.Body)
	job := new(model.Job)
	err := decoder.Decode(job)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}

	if s.userType != userCustomer {
		s.error(w, r, http.StatusInternalServerError, errors.New("current user is not a manager"))
		return
	}

	s.store.Mu.Lock()
	manager, err := s.store.Manager().FindByUser(user.ID)
	s.store.Mu.Unlock()

	if err != nil {
		log.Println("fail find manager", err)
		s.error(w, r, http.StatusNotFound, err)
	}

	s.store.Mu.Lock()
	err = s.store.Job().Create(job, manager)
	s.store.Mu.Unlock()

	if err != nil {
		log.Println("fail create job", err)
		s.error(w, r, http.StatusInternalServerError, err)
	}

	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleGetJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("wrong id"))
	}

	job, err := s.store.Job().Find(id)
	if err != nil {
		s.error(w, r, http.StatusNotFound, errors.New("job not found"))
	}

	s.respond(w, r, http.StatusOK, &job)
}


func (s *server) HandleEditFreelancer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)

	user, err, codeStatus := s.GetUserFromRequest(r)
	if err != nil {
		s.error(w, r, codeStatus, err)
		return
	}

	freelancer, err := s.store.Freelancer().FindByUser(user.ID)
	if err != nil {
		s.error(w,r,http.StatusNotFound, errors.New("no such freelancer"))
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
	}()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(freelancer)

	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("invalid format of data"))
		return
	}
	// TODO: validate freelancer

	err = s.store.Freelancer().Edit(freelancer)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusOK, struct{}{})
}

func (s *server) HandleGetFreelancer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("wrong id"))
	}

	freelancer, err := s.store.Freelancer().Find(id)
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
	}

	s.respond(w, r, http.StatusOK, &freelancer)
}


func (s *server) HandleGetAvatar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)

	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.Atoi(ids)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, errors.New("wrong id"))
		return
	}

	user, err := s.store.User().Find(id)
	if err != nil {
		s.error(w,r,http.StatusNotFound, errors.New("no such user in database"))
		return
	}

	if user.Avatar != nil {
		image := user.Avatar
		if _, err := w.Write(image); err != nil {
			s.error(w,r,http.StatusInternalServerError, err)
			return
		}

		w.Header().Set("Content-Type", "multipart/form-data")
		w.Header().Set("Content-Length", strconv.Itoa(len(image)))
	} else {
		filename := "internal/store/avatars/default.png"

		var openFile *os.File
		openFile, err = os.Open(filename)
		if err != nil {
			s.error(w, r, http.StatusNotFound, errors.New("cant open file"))
			return
		}
		defer func() {
			if err := openFile.Close(); err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
			}
		}()

		fileHeader := make([]byte, 100000)
		_, err := openFile.Read(fileHeader)
		if err !=  nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		fileContentType := http.DetectContentType(fileHeader)
		fileStat, _ := openFile.Stat()
		fileSize := strconv.FormatInt(fileStat.Size(), 10)

		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", fileContentType)
		w.Header().Set("Content-Length", fileSize)

		_, err = openFile.Seek(0, 0)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		_, err = io.Copy(w, openFile)
		if err != nil {
			s.error(w,r, http.StatusInternalServerError, err)
			return
		}
	}
	s.respond(w,r,http.StatusOK, struct{}{})
}
