package test

import (
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"testing"
	"time"
)

func testUser(t *testing.T) *model.User {
	t.Helper()
	return &model.User{
		ID: 1,
		FirstName: "masha",
		SecondName: "ivanova",
		UserName: "masha1996",
		Email: "masha@mail.ru",
		Password: "123456",
		EncryptPassword: "",
		Avatar: nil,
		UserType: "freelancer",
	}
}

// TODO: all helpers in another file test_helpers.go
func testManager(t *testing.T, user * model.User) *model.HireManager {
	t.Helper()
	return &model.HireManager{
		ID:					1,
		AccountID: 			user.ID,
		RegistrationDate:	time.Now(),
		Location:			"Moscow",
	}
}

func testFreelancer(t *testing.T, user *model.User) *model.Freelancer {
	t.Helper()
	return &model.Freelancer{
		ID:				   1,
		AccountId:         user.ID,
		RegistrationDate:  time.Now(),
		Country:           "Russia",
		City:              "Moscow",
		Address:           "moscow",
		Phone:             "111111111",
	}
}

func testJob(t *testing.T, manager * model.HireManager) *model.Job {
	t.Helper()
	return &model.Job{
		ID:				1,
		HireManagerId:	manager.ID,
		Title:          "title",
		Description:    "description",
		PaymentAmount:   11222,
		Country: 	"russia",
		City:	"moscow",
	}
}

func testResponse(t *testing.T, freelancer *model.Freelancer, job *model.Job) *model.Response {
	t.Helper()
	return &model.Response{
		ID:            1,
		FreelancerId:  freelancer.ID,
		JobId:         job.ID,
		Files:         "no files",
		Date:          time.Time{},
		StatusManager: model.ResponseStatusReview,
	}
}