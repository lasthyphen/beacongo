// Code generated by MockGen. DO NOT EDIT.
// Source: vms/platformvm/fx/fx.go

// Package fx is a generated GoMock package.
package fx

import (
	reflect "reflect"

	snow "github.com/lasthyphen/beacongo/snow"
	gomock "github.com/golang/mock/gomock"
)

// MockFx is a mock of Fx interface.
type MockFx struct {
	ctrl     *gomock.Controller
	recorder *MockFxMockRecorder
}

// MockFxMockRecorder is the mock recorder for MockFx.
type MockFxMockRecorder struct {
	mock *MockFx
}

// NewMockFx creates a new mock instance.
func NewMockFx(ctrl *gomock.Controller) *MockFx {
	mock := &MockFx{ctrl: ctrl}
	mock.recorder = &MockFxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFx) EXPECT() *MockFxMockRecorder {
	return m.recorder
}

// Bootstrapped mocks base method.
func (m *MockFx) Bootstrapped() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrapped")
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrapped indicates an expected call of Bootstrapped.
func (mr *MockFxMockRecorder) Bootstrapped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrapped", reflect.TypeOf((*MockFx)(nil).Bootstrapped))
}

// Bootstrapping mocks base method.
func (m *MockFx) Bootstrapping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrapping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrapping indicates an expected call of Bootstrapping.
func (mr *MockFxMockRecorder) Bootstrapping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrapping", reflect.TypeOf((*MockFx)(nil).Bootstrapping))
}

// CreateOutput mocks base method.
func (m *MockFx) CreateOutput(amount uint64, controlGroup interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOutput", amount, controlGroup)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOutput indicates an expected call of CreateOutput.
func (mr *MockFxMockRecorder) CreateOutput(amount, controlGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOutput", reflect.TypeOf((*MockFx)(nil).CreateOutput), amount, controlGroup)
}

// Initialize mocks base method.
func (m *MockFx) Initialize(vm interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", vm)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockFxMockRecorder) Initialize(vm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockFx)(nil).Initialize), vm)
}

// VerifyPermission mocks base method.
func (m *MockFx) VerifyPermission(tx, in, cred, controlGroup interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPermission", tx, in, cred, controlGroup)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyPermission indicates an expected call of VerifyPermission.
func (mr *MockFxMockRecorder) VerifyPermission(tx, in, cred, controlGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPermission", reflect.TypeOf((*MockFx)(nil).VerifyPermission), tx, in, cred, controlGroup)
}

// VerifyTransfer mocks base method.
func (m *MockFx) VerifyTransfer(tx, in, cred, utxo interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTransfer", tx, in, cred, utxo)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyTransfer indicates an expected call of VerifyTransfer.
func (mr *MockFxMockRecorder) VerifyTransfer(tx, in, cred, utxo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTransfer", reflect.TypeOf((*MockFx)(nil).VerifyTransfer), tx, in, cred, utxo)
}

// MockOwner is a mock of Owner interface.
type MockOwner struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerMockRecorder
}

// MockOwnerMockRecorder is the mock recorder for MockOwner.
type MockOwnerMockRecorder struct {
	mock *MockOwner
}

// NewMockOwner creates a new mock instance.
func NewMockOwner(ctrl *gomock.Controller) *MockOwner {
	mock := &MockOwner{ctrl: ctrl}
	mock.recorder = &MockOwnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOwner) EXPECT() *MockOwnerMockRecorder {
	return m.recorder
}

// InitCtx mocks base method.
func (m *MockOwner) InitCtx(ctx *snow.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitCtx", ctx)
}

// InitCtx indicates an expected call of InitCtx.
func (mr *MockOwnerMockRecorder) InitCtx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCtx", reflect.TypeOf((*MockOwner)(nil).InitCtx), ctx)
}

// Verify mocks base method.
func (m *MockOwner) Verify() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify")
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockOwnerMockRecorder) Verify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockOwner)(nil).Verify))
}

// MockOwned is a mock of Owned interface.
type MockOwned struct {
	ctrl     *gomock.Controller
	recorder *MockOwnedMockRecorder
}

// MockOwnedMockRecorder is the mock recorder for MockOwned.
type MockOwnedMockRecorder struct {
	mock *MockOwned
}

// NewMockOwned creates a new mock instance.
func NewMockOwned(ctrl *gomock.Controller) *MockOwned {
	mock := &MockOwned{ctrl: ctrl}
	mock.recorder = &MockOwnedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOwned) EXPECT() *MockOwnedMockRecorder {
	return m.recorder
}

// Owners mocks base method.
func (m *MockOwned) Owners() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owners")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Owners indicates an expected call of Owners.
func (mr *MockOwnedMockRecorder) Owners() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owners", reflect.TypeOf((*MockOwned)(nil).Owners))
}
