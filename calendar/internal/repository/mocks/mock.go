// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "http/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockStorage) CreateEvent(event models.Event) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", event)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockStorageMockRecorder) CreateEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockStorage)(nil).CreateEvent), event)
}

// GetEventById mocks base method.
func (m *MockStorage) GetEventById(arg0 string) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventById", arg0)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventById indicates an expected call of GetEventById.
func (mr *MockStorageMockRecorder) GetEventById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventById", reflect.TypeOf((*MockStorage)(nil).GetEventById), arg0)
}

// GetEvents mocks base method.
func (m *MockStorage) GetEvents() ([]models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents")
	ret0, _ := ret[0].([]models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockStorageMockRecorder) GetEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockStorage)(nil).GetEvents))
}

// GetUserByToken mocks base method.
func (m *MockStorage) GetUserByToken(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByToken", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByToken indicates an expected call of GetUserByToken.
func (mr *MockStorageMockRecorder) GetUserByToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByToken", reflect.TypeOf((*MockStorage)(nil).GetUserByToken), arg0)
}

// GetUserByUsername mocks base method.
func (m *MockStorage) GetUserByUsername(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStorageMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStorage)(nil).GetUserByUsername), arg0)
}

// UpdateEvent mocks base method.
func (m *MockStorage) UpdateEvent(event models.Event, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", event, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockStorageMockRecorder) UpdateEvent(event, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockStorage)(nil).UpdateEvent), event, id)
}

// UpdateUserTimezone mocks base method.
func (m *MockStorage) UpdateUserTimezone(token, timezone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserTimezone", token, timezone)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserTimezone indicates an expected call of UpdateUserTimezone.
func (mr *MockStorageMockRecorder) UpdateUserTimezone(token, timezone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserTimezone", reflect.TypeOf((*MockStorage)(nil).UpdateUserTimezone), token, timezone)
}

// UpdateUserToken mocks base method.
func (m *MockStorage) UpdateUserToken(token string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserToken", token, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserToken indicates an expected call of UpdateUserToken.
func (mr *MockStorageMockRecorder) UpdateUserToken(token, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserToken", reflect.TypeOf((*MockStorage)(nil).UpdateUserToken), token, id)
}
