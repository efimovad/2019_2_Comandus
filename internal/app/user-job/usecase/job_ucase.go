package jobUcase

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/app/manager/delivery/grpc/manager_grpc"
	user_job "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/user-job"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"time"
)

type JobUsecase struct {
	jobRep			user_job.Repository
}

func NewJobUsecase(j user_job.Repository) user_job.Usecase {
	return &JobUsecase{
		jobRep:			j,
	}
}

func (u *JobUsecase) getManagerByUserFromServer(id int64) (*manager_grpc.Manager, error) {
	conn, err := grpc.Dial(":8084", grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "grpc.Dial()")
	}
	defer func(){
		if err := conn.Close(); err != nil {
			// TODO: use zap logger
			log.Println("conn.Close()", err)
		}
	}()

	client := manager_grpc.NewManagerHandlerClient(conn)

	userReq := &manager_grpc.UserID{
		ID:		id,
	}

	currManager, err := client.FindByUser(context.Background(), userReq)
	if err != nil {
		return nil, errors.Wrap(err, "userRep.Find()")
	}

	return currManager, nil
}

func (u *JobUsecase) CreateJob(currUser * model.User, job *model.Job) error {
	if !currUser.IsManager() {
		return errors.New("current user is not a manager")
	}

	currManager, err := u.getManagerByUserFromServer(currUser.ID)
	if err != nil {
		return errors.Wrapf(err, "getManagerByUserFromServer()")
	}

	job.HireManagerId = currManager.ID
	job.Date = time.Now()
	if err = u.jobRep.Create(job); err != nil {
		return errors.Wrapf(err, "jobRep.Create()")
	}
	return nil
}

func (u *JobUsecase) FindJob(id int64) (*model.Job, error) {
	job, err := u.jobRep.Find(id)
	if err != nil {
		return nil, errors.Wrap(err, "jobRep.Find()")
	}
	return job, nil
}

func (u *JobUsecase) GetAllJobs() ([]model.Job, error) {
	jobs, err := u.jobRep.List()
	if err != nil {
		err = errors.Wrapf(err, "HandleGetJob<-Find")
		return nil, errors.Wrap(err, "jobRep.List()")
	}
	return jobs, nil
}

func (u *JobUsecase) EditJob(user *model.User, inputJob *model.Job, id int64) error {
	if !user.IsManager() {
		return errors.New("only manager can edit job")
	}

	job, err := u.jobRep.Find(id)
	if err != nil {
		return errors.Wrapf(err, "jobRep.Find(): ")
	}

	currManager, err := u.getManagerByUserFromServer(user.ID)
	if err != nil {
		return errors.Wrap(err, "getManagerByUserFromServer()")
	}

	if job.HireManagerId != currManager.ID {
		return errors.New("no access for current manager")
	}

	inputJob.ID = job.ID
	inputJob.HireManagerId = job.HireManagerId

	if err := u.jobRep.Edit(inputJob); err != nil {
		return errors.Wrapf(err, "jobRep.Edit()")
	}
	return nil
}