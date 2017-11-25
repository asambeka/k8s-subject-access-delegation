// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/core/v1 (interfaces: ServiceAccountInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	reflect "reflect"
)

// MockServiceAccountInterface is a mock of ServiceAccountInterface interface
type MockServiceAccountInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountInterfaceMockRecorder
}

// MockServiceAccountInterfaceMockRecorder is the mock recorder for MockServiceAccountInterface
type MockServiceAccountInterfaceMockRecorder struct {
	mock *MockServiceAccountInterface
}

// NewMockServiceAccountInterface creates a new mock instance
func NewMockServiceAccountInterface(ctrl *gomock.Controller) *MockServiceAccountInterface {
	mock := &MockServiceAccountInterface{ctrl: ctrl}
	mock.recorder = &MockServiceAccountInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceAccountInterface) EXPECT() *MockServiceAccountInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServiceAccountInterface) Create(arg0 *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServiceAccountInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceAccountInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockServiceAccountInterface) Delete(arg0 string, arg1 *v10.DeleteOptions) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockServiceAccountInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceAccountInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockServiceAccountInterface) DeleteCollection(arg0 *v10.DeleteOptions, arg1 v10.ListOptions) error {
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockServiceAccountInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockServiceAccountInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockServiceAccountInterface) Get(arg0 string, arg1 v10.GetOptions) (*v1.ServiceAccount, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServiceAccountInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceAccountInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockServiceAccountInterface) List(arg0 v10.ListOptions) (*v1.ServiceAccountList, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.ServiceAccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServiceAccountInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceAccountInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockServiceAccountInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1.ServiceAccount, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockServiceAccountInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockServiceAccountInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockServiceAccountInterface) Update(arg0 *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockServiceAccountInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceAccountInterface)(nil).Update), arg0)
}

// Watch mocks base method
func (m *MockServiceAccountInterface) Watch(arg0 v10.ListOptions) (watch.Interface, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockServiceAccountInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockServiceAccountInterface)(nil).Watch), arg0)
}
