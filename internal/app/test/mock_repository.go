// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/repository.go

// Package mock_store is a generated GoMock package.
package test

import (
	model "github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserRepository) Create(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// Find mocks base method
func (m *MockUserRepository) Find(arg0 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserRepository)(nil).Find), arg0)
}

// FindByEmail mocks base method
func (m *MockUserRepository) FindByEmail(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserRepositoryMockRecorder) FindByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), arg0)
}

// Edit mocks base method
func (m *MockUserRepository) Edit(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockUserRepositoryMockRecorder) Edit(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockUserRepository)(nil).Edit), user)
}

// MockFreelancerRepository is a mock of FreelancerRepository interface
type MockFreelancerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFreelancerRepositoryMockRecorder
}

// MockFreelancerRepositoryMockRecorder is the mock recorder for MockFreelancerRepository
type MockFreelancerRepositoryMockRecorder struct {
	mock *MockFreelancerRepository
}

// NewMockFreelancerRepository creates a new mock instance
func NewMockFreelancerRepository(ctrl *gomock.Controller) *MockFreelancerRepository {
	mock := &MockFreelancerRepository{ctrl: ctrl}
	mock.recorder = &MockFreelancerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFreelancerRepository) EXPECT() *MockFreelancerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFreelancerRepository) Create(freelancer *model.Freelancer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", freelancer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockFreelancerRepositoryMockRecorder) Create(freelancer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFreelancerRepository)(nil).Create), freelancer)
}

// Find mocks base method
func (m *MockFreelancerRepository) Find(arg0 int64) (*model.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFreelancerRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFreelancerRepository)(nil).Find), arg0)
}

// FindByUser mocks base method
func (m *MockFreelancerRepository) FindByUser(arg0 int64) (*model.Freelancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0)
	ret0, _ := ret[0].(*model.Freelancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser
func (mr *MockFreelancerRepositoryMockRecorder) FindByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockFreelancerRepository)(nil).FindByUser), arg0)
}

// Edit mocks base method
func (m *MockFreelancerRepository) Edit(freelancer *model.Freelancer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", freelancer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockFreelancerRepositoryMockRecorder) Edit(freelancer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockFreelancerRepository)(nil).Edit), freelancer)
}

// MockManagerRepository is a mock of ManagerRepository interface
type MockManagerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockManagerRepositoryMockRecorder
}

// MockManagerRepositoryMockRecorder is the mock recorder for MockManagerRepository
type MockManagerRepositoryMockRecorder struct {
	mock *MockManagerRepository
}

// NewMockManagerRepository creates a new mock instance
func NewMockManagerRepository(ctrl *gomock.Controller) *MockManagerRepository {
	mock := &MockManagerRepository{ctrl: ctrl}
	mock.recorder = &MockManagerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerRepository) EXPECT() *MockManagerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockManagerRepository) Create(manager *model.HireManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", manager)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockManagerRepositoryMockRecorder) Create(manager interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManagerRepository)(nil).Create), manager)
}

// Find mocks base method
func (m *MockManagerRepository) Find(arg0 int64) (*model.HireManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.HireManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockManagerRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockManagerRepository)(nil).Find), arg0)
}

// FindByUser mocks base method
func (m *MockManagerRepository) FindByUser(arg0 int64) (*model.HireManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", arg0)
	ret0, _ := ret[0].(*model.HireManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser
func (mr *MockManagerRepositoryMockRecorder) FindByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockManagerRepository)(nil).FindByUser), arg0)
}

// Edit mocks base method
func (m *MockManagerRepository) Edit(manager *model.HireManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", manager)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockManagerRepositoryMockRecorder) Edit(manager interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockManagerRepository)(nil).Edit), manager)
}

// MockJobRepository is a mock of JobRepository interface
type MockJobRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobRepositoryMockRecorder
}

// MockJobRepositoryMockRecorder is the mock recorder for MockJobRepository
type MockJobRepositoryMockRecorder struct {
	mock *MockJobRepository
}

// NewMockJobRepository creates a new mock instance
func NewMockJobRepository(ctrl *gomock.Controller) *MockJobRepository {
	mock := &MockJobRepository{ctrl: ctrl}
	mock.recorder = &MockJobRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobRepository) EXPECT() *MockJobRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m_2 *MockJobRepository) Create(j *model.Job, m *model.HireManager) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", j, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockJobRepositoryMockRecorder) Create(j, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockJobRepository)(nil).Create), j, m)
}

// Find mocks base method
func (m *MockJobRepository) Find(arg0 int64) (*model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockJobRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockJobRepository)(nil).Find), arg0)
}

// Edit mocks base method
func (m *MockJobRepository) Edit(job *model.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockJobRepositoryMockRecorder) Edit(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockJobRepository)(nil).Edit), job)
}

// List mocks base method
func (m *MockJobRepository) List() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockJobRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockJobRepository)(nil).List))
}

// MockResponseRepository is a mock of ResponseRepository interface
type MockResponseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResponseRepositoryMockRecorder
}

// MockResponseRepositoryMockRecorder is the mock recorder for MockResponseRepository
type MockResponseRepositoryMockRecorder struct {
	mock *MockResponseRepository
}

// NewMockResponseRepository creates a new mock instance
func NewMockResponseRepository(ctrl *gomock.Controller) *MockResponseRepository {
	mock := &MockResponseRepository{ctrl: ctrl}
	mock.recorder = &MockResponseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponseRepository) EXPECT() *MockResponseRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockResponseRepository) Create(response *model.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", response)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockResponseRepositoryMockRecorder) Create(response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResponseRepository)(nil).Create), response)
}

// Edit mocks base method
func (m *MockResponseRepository) Edit(response *model.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", response)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockResponseRepositoryMockRecorder) Edit(response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockResponseRepository)(nil).Edit), response)
}

// ListForFreelancer mocks base method
func (m *MockResponseRepository) ListForFreelancer(arg0 int64) ([]model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForFreelancer", arg0)
	ret0, _ := ret[0].([]model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForFreelancer indicates an expected call of ListForFreelancer
func (mr *MockResponseRepositoryMockRecorder) ListForFreelancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForFreelancer", reflect.TypeOf((*MockResponseRepository)(nil).ListForFreelancer), arg0)
}

// ListForManager mocks base method
func (m *MockResponseRepository) ListForManager(arg0 int64) ([]model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForManager", arg0)
	ret0, _ := ret[0].([]model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForManager indicates an expected call of ListForManager
func (mr *MockResponseRepositoryMockRecorder) ListForManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForManager", reflect.TypeOf((*MockResponseRepository)(nil).ListForManager), arg0)
}

// Find mocks base method
func (m *MockResponseRepository) Find(arg0 int64) (*model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockResponseRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockResponseRepository)(nil).Find), arg0)
}

// MockCompanyRepository is a mock of CompanyRepository interface
type MockCompanyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyRepositoryMockRecorder
}

// MockCompanyRepositoryMockRecorder is the mock recorder for MockCompanyRepository
type MockCompanyRepositoryMockRecorder struct {
	mock *MockCompanyRepository
}

// NewMockCompanyRepository creates a new mock instance
func NewMockCompanyRepository(ctrl *gomock.Controller) *MockCompanyRepository {
	mock := &MockCompanyRepository{ctrl: ctrl}
	mock.recorder = &MockCompanyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompanyRepository) EXPECT() *MockCompanyRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCompanyRepository) Create(company *model.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", company)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCompanyRepositoryMockRecorder) Create(company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCompanyRepository)(nil).Create), company)
}

// Find mocks base method
func (m *MockCompanyRepository) Find(arg0 int64) (*model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockCompanyRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCompanyRepository)(nil).Find), arg0)
}

// Edit mocks base method
func (m *MockCompanyRepository) Edit(company *model.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", company)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockCompanyRepositoryMockRecorder) Edit(company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockCompanyRepository)(nil).Edit), company)
}

// MockContractRepository is a mock of ContractRepository interface
type MockContractRepository struct {
	ctrl     *gomock.Controller
	recorder *MockContractRepositoryMockRecorder
}

// MockContractRepositoryMockRecorder is the mock recorder for MockContractRepository
type MockContractRepositoryMockRecorder struct {
	mock *MockContractRepository
}

// NewMockContractRepository creates a new mock instance
func NewMockContractRepository(ctrl *gomock.Controller) *MockContractRepository {
	mock := &MockContractRepository{ctrl: ctrl}
	mock.recorder = &MockContractRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractRepository) EXPECT() *MockContractRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockContractRepository) Create(contract *model.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", contract)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockContractRepositoryMockRecorder) Create(contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContractRepository)(nil).Create), contract)
}

// Edit mocks base method
func (m *MockContractRepository) Edit(contract *model.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", contract)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit
func (mr *MockContractRepositoryMockRecorder) Edit(contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockContractRepository)(nil).Edit), contract)
}

// List mocks base method
func (m *MockContractRepository) List(arg0 int64, arg1 string) ([]model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockContractRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContractRepository)(nil).List), arg0, arg1)
}

// Find mocks base method
func (m *MockContractRepository) Find(arg0 int64) (*model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockContractRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockContractRepository)(nil).Find), arg0)
}