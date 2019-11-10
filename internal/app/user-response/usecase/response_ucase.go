package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/app/freelancer"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/app/manager"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/app/user"
	user_job "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/user-job"
	user_response "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/user-response"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/pkg/errors"
	"time"
)

type ResponseUsecase struct {
	userRep			user.Repository
	managerRep		manager.Repository
	freelancerRep	freelancer.Repository
	jobRep			user_job.Repository
	responseRep		user_response.Repository
}

func NewResponseUsecase(u user.Repository,
	m manager.Repository,
	f freelancer.Repository,
	j user_job.Repository,
	r user_response.Repository) user_response.Usecase {
	return &ResponseUsecase{
		userRep:		u,
		managerRep:		m,
		freelancerRep:	f,
		jobRep:			j,
		responseRep:	r,
	}
}

func (u * ResponseUsecase) CreateResponse(user *model.User, jobId int64) error {
	if user.IsManager() {
		return errors.New("to response user need to be freelancer")
	}

	currFreelancer, err := u.freelancerRep.FindByUser(user.ID)
	if err != nil {
		return errors.Wrap(err, "freelancerRep.FindByUser: ")
	}

	// TODO: get files from request
	response := &model.Response{
		ID:               0,
		FreelancerId:     currFreelancer.ID,
		JobId:            jobId,
		Files:            "",
		Date:             time.Now(),
		StatusManager:    model.ResponseStatusReview,
		StatusFreelancer: model.ResponseStatusBlock,
	}

	if err := response.Validate(0); err != nil {
		return errors.Wrapf(err, "Validate: ")
	}

	if err := u.responseRep.Create(response); err != nil {
		return errors.Wrapf(err, "responseRep.Create(): ")
	}
	return nil
}

func (u * ResponseUsecase) GetResponses(user *model.User) (*[]model.Response, error){
	var responses []model.Response

	if user.IsManager() {
		currManager, err := u.managerRep.FindByUser(user.ID)
		if err != nil {
			err = errors.Wrapf(err, " GetManagerResponses<-Manager().FindByUser: ")
			return nil, err
		}

		responses, err = u.responseRep.ListForManager(currManager.ID)
		if err != nil {
			err = errors.Wrapf(err, "GetManagerResponses<-Responses().ListForManager: ")
			return nil, err
		}
	} else {
		currFreelancer, err := u.freelancerRep.FindByUser(user.ID)
		if err != nil {
			err = errors.Wrapf(err, " GetManagerResponses<-Manager().FindByUser: ")
			return nil, err
		}

		responses, err = u.responseRep.ListForFreelancer(currFreelancer.ID)
		if err != nil {
			err = errors.Wrapf(err, "GetManagerResponses<-Responses().ListForManager: ")
			return nil, err
		}
	}

	return &responses, nil
}

func (u * ResponseUsecase) AcceptResponse(user *model.User, responseId int64) error {
	response, err := u.responseRep.Find(responseId)
	if err != nil {
		return errors.Wrapf(err, "responseRep.Find(): ")
	}

	job, err := u.jobRep.Find(response.JobId)
	if err != nil {
		return errors.Wrapf(err, "jobRep<-Find(): ")
	}

	if user.IsManager() {
		currManager, err := u.managerRep.FindByUser(user.ID)
		if err != nil {
			return errors.Wrapf(err, "managerRep.FindByUser: ")
		}

		if job.HireManagerId != currManager.ID {
			return errors.New("current manager cant accept this response")
		}
		response.StatusManager = model.ResponseStatusAccepted
		response.StatusFreelancer = model.ResponseStatusReview
	} else {
		currFreelancer, err := u.freelancerRep.FindByUser(user.ID)
		if err != nil {
			return errors.Wrapf(err, "freelancerRep.FindByUser(): ")
		}

		if currFreelancer.ID != response.FreelancerId {
			return errors.New("current freelancer can't accept this response")
		}

		if response.StatusFreelancer == model.ResponseStatusBlock {
			return errors.New("freelancer can't accept response before manager")
		}

		response.StatusManager = model.ResponseStatusAccepted
	}
	return nil
}

func (u *ResponseUsecase) DenyResponse(user *model.User, responseId int64) error {
	response, err := u.responseRep.Find(responseId)
	if err != nil {
		return errors.Wrapf(err, "responseRep.Find(): ")
	}

	job, err := u.jobRep.Find(response.JobId)
	if err != nil {
		return errors.Wrapf(err, "jobRep<-Find(): ")
	}

	if user.IsManager() {
		currManager, err := u.managerRep.FindByUser(user.ID)
		if err != nil {
			return errors.Wrapf(err, "managerRep.FindByUser: ")
		}

		if job.HireManagerId != currManager.ID {
			return errors.New("current manager cant accept this response")
		}

		response.StatusManager = model.ResponseStatusDenied
		response.StatusFreelancer = model.ResponseStatusBlock
	} else {
		currFreelancer, err := u.freelancerRep.FindByUser(user.ID)
		if err != nil {
			return errors.Wrapf(err, "freelancerRep.FindByUser(): ")
		}

		if currFreelancer.ID != response.FreelancerId {
			return errors.New("current freelancer can't accept this response")
		}

		if response.StatusFreelancer == model.ResponseStatusBlock {
			return errors.New("freelancer can't accept response before manager")
		}

		response.StatusManager = model.ResponseStatusDenied
	}
	return nil
}