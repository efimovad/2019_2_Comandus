// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/user/usecase.go

// Package mock_user is a generated GoMock package.
package test

import (
	model "github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUsecase) CreateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUsecaseMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsecase)(nil).CreateUser), arg0)
}

// VerifyUser mocks base method
func (m *MockUsecase) VerifyUser(arg0 *model.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyUser indicates an expected call of VerifyUser
func (mr *MockUsecaseMockRecorder) VerifyUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockUsecase)(nil).VerifyUser), arg0)
}

// EditUser mocks base method
func (m *MockUsecase) EditUser(newUser, oldUser *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", newUser, oldUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser
func (mr *MockUsecaseMockRecorder) EditUser(newUser, oldUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockUsecase)(nil).EditUser), newUser, oldUser)
}

// EditUserPassword mocks base method
func (m *MockUsecase) EditUserPassword(passwords *model.BodyPassword, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUserPassword", passwords, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUserPassword indicates an expected call of EditUserPassword
func (mr *MockUsecaseMockRecorder) EditUserPassword(passwords, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUserPassword", reflect.TypeOf((*MockUsecase)(nil).EditUserPassword), passwords, user)
}

// GetAvatar mocks base method
func (m *MockUsecase) GetAvatar(user *model.User) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar", user)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatar indicates an expected call of GetAvatar
func (mr *MockUsecaseMockRecorder) GetAvatar(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockUsecase)(nil).GetAvatar), user)
}

// Find mocks base method
func (m *MockUsecase) Find(arg0 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUsecaseMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUsecase)(nil).Find), arg0)
}

// SetUserType mocks base method
func (m *MockUsecase) SetUserType(user *model.User, userType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserType", user, userType)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserType indicates an expected call of SetUserType
func (mr *MockUsecaseMockRecorder) SetUserType(user, userType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserType", reflect.TypeOf((*MockUsecase)(nil).SetUserType), user, userType)
}

// GetRoles mocks base method
func (m *MockUsecase) GetRoles(user *model.User) ([]*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", user)
	ret0, _ := ret[0].([]*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles
func (mr *MockUsecaseMockRecorder) GetRoles(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockUsecase)(nil).GetRoles), user)
}