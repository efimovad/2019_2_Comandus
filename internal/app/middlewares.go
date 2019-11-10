package apiserver

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
)

func (s *server) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		u, err := s.store.User().Find(int64(id.(int)))

		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, &u)))
	})
}

func (s *server) RequestIDMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := strconv.Itoa(rand.Int())
		ctx := r.Context()
		ctx = context.WithValue(ctx, "rIDKey", reqID)
		w.Header().Set("Request-ID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) AccessLogMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info(r.URL.Path,
			zap.String("method:", r.Method),
			zap.String("remote_addr:", r.RemoteAddr),
			zap.String("url:", r.URL.Path),
		)
		next.ServeHTTP(w, r)
	})
}

func (s *server) CORSMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Lol")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", s.clientUrl)
		if r.Method == http.MethodOptions{
			s.respond(w , r , http.StatusOK, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *server) CheckTokenMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/token" {
			sess, err := s.sessionStore.Get(r, sessionName)
			if err != nil {
				err = errors.Wrapf(err, "CheckTokenMiddleware<-sessionStore.Get :")
				s.error(w, r, http.StatusUnauthorized, err)
				return
			}

			isEqual, err := s.token.Check(sess, r.Header.Get("csrf-token"))
			if !isEqual {
				err = errors.Wrapf(err, "CheckTokenMiddleware<-Check:")
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}