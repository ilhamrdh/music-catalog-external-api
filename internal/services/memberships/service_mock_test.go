// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package memberships is a generated GoMock package.
package memberships

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	memberships "github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *Mockrepository) CreateUser(model memberships.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockrepositoryMockRecorder) CreateUser(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Mockrepository)(nil).CreateUser), model)
}

// GetUser mocks base method.
func (m *Mockrepository) GetUser(email, username string, id uint) (*memberships.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", email, username, id)
	ret0, _ := ret[0].(*memberships.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockrepositoryMockRecorder) GetUser(email, username, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*Mockrepository)(nil).GetUser), email, username, id)
}
