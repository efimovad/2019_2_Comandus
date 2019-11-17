// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/user-job/usecase.go

// Package test is a generated GoMock package.
package test

import (
	model "github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJobUsecase is a mock of Usecase interface
type MockJobUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockJobUsecaseMockRecorder
}

// MockJobUsecaseMockRecorder is the mock recorder for MockJobUsecase
type MockJobUsecaseMockRecorder struct {
	mock *MockJobUsecase
}

// NewMockJobUsecase creates a new mock instance
func NewMockJobUsecase(ctrl *gomock.Controller) *MockJobUsecase {
	mock := &MockJobUsecase{ctrl: ctrl}
	mock.recorder = &MockJobUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobUsecase) EXPECT() *MockJobUsecaseMockRecorder {
	return m.recorder
}

// CreateJob mocks base method
func (m *MockJobUsecase) CreateJob(user *model.User, job *model.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", user, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateJob indicates an expected call of CreateJob
func (mr *MockJobUsecaseMockRecorder) CreateJob(user, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockJobUsecase)(nil).CreateJob), user, job)
}

// FindJob mocks base method
func (m *MockJobUsecase) FindJob(id int64) (*model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindJob", id)
	ret0, _ := ret[0].(*model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindJob indicates an expected call of FindJob
func (mr *MockJobUsecaseMockRecorder) FindJob(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindJob", reflect.TypeOf((*MockJobUsecase)(nil).FindJob), id)
}

// GetAllJobs mocks base method
func (m *MockJobUsecase) GetAllJobs() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJobs")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs
func (mr *MockJobUsecaseMockRecorder) GetAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockJobUsecase)(nil).GetAllJobs))
}

// EditJob mocks base method
func (m *MockJobUsecase) EditJob(user *model.User, job *model.Job, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditJob", user, job, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditJob indicates an expected call of EditJob
func (mr *MockJobUsecaseMockRecorder) EditJob(user, job, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditJob", reflect.TypeOf((*MockJobUsecase)(nil).EditJob), user, job, id)
}
