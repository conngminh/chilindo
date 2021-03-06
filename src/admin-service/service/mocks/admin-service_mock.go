// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/admin-service/service/admin-service.go

// Package service is a generated GoMock package.
package service

import (
	dto "chilindo/src/admin-service/dto"
	entity "chilindo/src/admin-service/entity"
	admin "chilindo/src/pkg/pb/admin"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAdminService is a mock of IAdminService interface.
type MockIAdminService struct {
	ctrl     *gomock.Controller
	recorder *MockIAdminServiceMockRecorder
}

// MockIAdminServiceMockRecorder is the mock recorder for MockIAdminService.
type MockIAdminServiceMockRecorder struct {
	mock *MockIAdminService
}

// NewMockIAdminService creates a new mock instance.
func NewMockIAdminService(ctrl *gomock.Controller) *MockIAdminService {
	mock := &MockIAdminService{ctrl: ctrl}
	mock.recorder = &MockIAdminServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAdminService) EXPECT() *MockIAdminServiceMockRecorder {
	return m.recorder
}

// CheckIsAuth mocks base method.
func (m *MockIAdminService) CheckIsAuth(req *admin.CheckIsAuthRequest) (*admin.CheckIsAuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsAuth", req)
	ret0, _ := ret[0].(*admin.CheckIsAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsAuth indicates an expected call of CheckIsAuth.
func (mr *MockIAdminServiceMockRecorder) CheckIsAuth(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsAuth", reflect.TypeOf((*MockIAdminService)(nil).CheckIsAuth), req)
}

// CreateAdmin mocks base method.
func (m *MockIAdminService) CreateAdmin(admin *entity.Admin) (*entity.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", admin)
	ret0, _ := ret[0].(*entity.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockIAdminServiceMockRecorder) CreateAdmin(admin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockIAdminService)(nil).CreateAdmin), admin)
}

// IsDuplicateUsername mocks base method.
func (m *MockIAdminService) IsDuplicateUsername(username string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDuplicateUsername", username)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDuplicateUsername indicates an expected call of IsDuplicateUsername.
func (mr *MockIAdminServiceMockRecorder) IsDuplicateUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDuplicateUsername", reflect.TypeOf((*MockIAdminService)(nil).IsDuplicateUsername), username)
}

// VerifyCredential mocks base method.
func (m *MockIAdminService) VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCredential", loginDTO)
	ret0, _ := ret[0].(*entity.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyCredential indicates an expected call of VerifyCredential.
func (mr *MockIAdminServiceMockRecorder) VerifyCredential(loginDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCredential", reflect.TypeOf((*MockIAdminService)(nil).VerifyCredential), loginDTO)
}
