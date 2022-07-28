// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/product-service/repository/product-option-repository.go

// Package service is a generated GoMock package.
package service

import (
	dto "chilindo/src/product-service/dto"
	entity "chilindo/src/product-service/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductOptionRepository is a mock of ProductOptionRepository interface.
type MockProductOptionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductOptionRepositoryMockRecorder
}

// MockProductOptionRepositoryMockRecorder is the mock recorder for MockProductOptionRepository.
type MockProductOptionRepositoryMockRecorder struct {
	mock *MockProductOptionRepository
}

// NewMockProductOptionRepository creates a new mock instance.
func NewMockProductOptionRepository(ctrl *gomock.Controller) *MockProductOptionRepository {
	mock := &MockProductOptionRepository{ctrl: ctrl}
	mock.recorder = &MockProductOptionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductOptionRepository) EXPECT() *MockProductOptionRepositoryMockRecorder {
	return m.recorder
}

// CreateOption mocks base method.
func (m *MockProductOptionRepository) CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOption", b)
	ret0, _ := ret[0].(*entity.ProductOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOption indicates an expected call of CreateOption.
func (mr *MockProductOptionRepositoryMockRecorder) CreateOption(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOption", reflect.TypeOf((*MockProductOptionRepository)(nil).CreateOption), b)
}

// DeleteOption mocks base method.
func (m *MockProductOptionRepository) DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOption", b)
	ret0, _ := ret[0].(*entity.ProductOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOption indicates an expected call of DeleteOption.
func (mr *MockProductOptionRepositoryMockRecorder) DeleteOption(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOption", reflect.TypeOf((*MockProductOptionRepository)(nil).DeleteOption), b)
}

// GetOptionByID mocks base method.
func (m *MockProductOptionRepository) GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptionByID", b)
	ret0, _ := ret[0].(*entity.ProductOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptionByID indicates an expected call of GetOptionByID.
func (mr *MockProductOptionRepositoryMockRecorder) GetOptionByID(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptionByID", reflect.TypeOf((*MockProductOptionRepository)(nil).GetOptionByID), b)
}

// GetOptions mocks base method.
func (m *MockProductOptionRepository) GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptions", b)
	ret0, _ := ret[0].(*[]entity.ProductOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptions indicates an expected call of GetOptions.
func (mr *MockProductOptionRepositoryMockRecorder) GetOptions(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptions", reflect.TypeOf((*MockProductOptionRepository)(nil).GetOptions), b)
}

// ProductOptionByID mocks base method.
func (m *MockProductOptionRepository) ProductOptionByID(b *dto.ProductDTO) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductOptionByID", b)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductOptionByID indicates an expected call of ProductOptionByID.
func (mr *MockProductOptionRepositoryMockRecorder) ProductOptionByID(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductOptionByID", reflect.TypeOf((*MockProductOptionRepository)(nil).ProductOptionByID), b)
}

// UpdateOption mocks base method.
func (m *MockProductOptionRepository) UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOption", b)
	ret0, _ := ret[0].(*entity.ProductOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOption indicates an expected call of UpdateOption.
func (mr *MockProductOptionRepositoryMockRecorder) UpdateOption(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOption", reflect.TypeOf((*MockProductOptionRepository)(nil).UpdateOption), b)
}