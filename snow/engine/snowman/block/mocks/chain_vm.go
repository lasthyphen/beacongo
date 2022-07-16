// Code generated by MockGen. DO NOT EDIT.
// Source: snow/engine/snowman/block/vm.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	manager "github.com/lasthyphen/beacongo/database/manager"
	ids "github.com/lasthyphen/beacongo/ids"
	snow "github.com/lasthyphen/beacongo/snow"
	snowman "github.com/lasthyphen/beacongo/snow/consensus/snowman"
	common "github.com/lasthyphen/beacongo/snow/engine/common"
	version "github.com/lasthyphen/beacongo/version"
	gomock "github.com/golang/mock/gomock"
)

// MockChainVM is a mock of ChainVM interface.
type MockChainVM struct {
	ctrl     *gomock.Controller
	recorder *MockChainVMMockRecorder
}

// MockChainVMMockRecorder is the mock recorder for MockChainVM.
type MockChainVMMockRecorder struct {
	mock *MockChainVM
}

// NewMockChainVM creates a new mock instance.
func NewMockChainVM(ctrl *gomock.Controller) *MockChainVM {
	mock := &MockChainVM{ctrl: ctrl}
	mock.recorder = &MockChainVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainVM) EXPECT() *MockChainVMMockRecorder {
	return m.recorder
}

// AppGossip mocks base method.
func (m *MockChainVM) AppGossip(nodeID ids.NodeID, msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGossip", nodeID, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppGossip indicates an expected call of AppGossip.
func (mr *MockChainVMMockRecorder) AppGossip(nodeID, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGossip", reflect.TypeOf((*MockChainVM)(nil).AppGossip), nodeID, msg)
}

// AppRequest mocks base method.
func (m *MockChainVM) AppRequest(nodeID ids.NodeID, requestID uint32, deadline time.Time, request []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequest", nodeID, requestID, deadline, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequest indicates an expected call of AppRequest.
func (mr *MockChainVMMockRecorder) AppRequest(nodeID, requestID, deadline, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequest", reflect.TypeOf((*MockChainVM)(nil).AppRequest), nodeID, requestID, deadline, request)
}

// AppRequestFailed mocks base method.
func (m *MockChainVM) AppRequestFailed(nodeID ids.NodeID, requestID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequestFailed", nodeID, requestID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequestFailed indicates an expected call of AppRequestFailed.
func (mr *MockChainVMMockRecorder) AppRequestFailed(nodeID, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequestFailed", reflect.TypeOf((*MockChainVM)(nil).AppRequestFailed), nodeID, requestID)
}

// AppResponse mocks base method.
func (m *MockChainVM) AppResponse(nodeID ids.NodeID, requestID uint32, response []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", nodeID, requestID, response)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppResponse indicates an expected call of AppResponse.
func (mr *MockChainVMMockRecorder) AppResponse(nodeID, requestID, response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*MockChainVM)(nil).AppResponse), nodeID, requestID, response)
}

// BuildBlock mocks base method.
func (m *MockChainVM) BuildBlock() (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildBlock")
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildBlock indicates an expected call of BuildBlock.
func (mr *MockChainVMMockRecorder) BuildBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildBlock", reflect.TypeOf((*MockChainVM)(nil).BuildBlock))
}

// Connected mocks base method.
func (m *MockChainVM) Connected(id ids.NodeID, nodeVersion version.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected", id, nodeVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connected indicates an expected call of Connected.
func (mr *MockChainVMMockRecorder) Connected(id, nodeVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockChainVM)(nil).Connected), id, nodeVersion)
}

// CreateHandlers mocks base method.
func (m *MockChainVM) CreateHandlers() (map[string]*common.HTTPHandler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHandlers")
	ret0, _ := ret[0].(map[string]*common.HTTPHandler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHandlers indicates an expected call of CreateHandlers.
func (mr *MockChainVMMockRecorder) CreateHandlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHandlers", reflect.TypeOf((*MockChainVM)(nil).CreateHandlers))
}

// CreateStaticHandlers mocks base method.
func (m *MockChainVM) CreateStaticHandlers() (map[string]*common.HTTPHandler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStaticHandlers")
	ret0, _ := ret[0].(map[string]*common.HTTPHandler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStaticHandlers indicates an expected call of CreateStaticHandlers.
func (mr *MockChainVMMockRecorder) CreateStaticHandlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStaticHandlers", reflect.TypeOf((*MockChainVM)(nil).CreateStaticHandlers))
}

// Disconnected mocks base method.
func (m *MockChainVM) Disconnected(id ids.NodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnected", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockChainVMMockRecorder) Disconnected(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockChainVM)(nil).Disconnected), id)
}

// GetBlock mocks base method.
func (m *MockChainVM) GetBlock(arg0 ids.ID) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockChainVMMockRecorder) GetBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockChainVM)(nil).GetBlock), arg0)
}

// HealthCheck mocks base method.
func (m *MockChainVM) HealthCheck() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockChainVMMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockChainVM)(nil).HealthCheck))
}

// Initialize mocks base method.
func (m *MockChainVM) Initialize(ctx *snow.Context, dbManager manager.Manager, genesisBytes, upgradeBytes, configBytes []byte, toEngine chan<- common.Message, fxs []*common.Fx, appSender common.AppSender) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockChainVMMockRecorder) Initialize(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockChainVM)(nil).Initialize), ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)
}

// LastAccepted mocks base method.
func (m *MockChainVM) LastAccepted() (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastAccepted")
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastAccepted indicates an expected call of LastAccepted.
func (mr *MockChainVMMockRecorder) LastAccepted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastAccepted", reflect.TypeOf((*MockChainVM)(nil).LastAccepted))
}

// ParseBlock mocks base method.
func (m *MockChainVM) ParseBlock(arg0 []byte) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseBlock indicates an expected call of ParseBlock.
func (mr *MockChainVMMockRecorder) ParseBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBlock", reflect.TypeOf((*MockChainVM)(nil).ParseBlock), arg0)
}

// SetPreference mocks base method.
func (m *MockChainVM) SetPreference(arg0 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPreference", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPreference indicates an expected call of SetPreference.
func (mr *MockChainVMMockRecorder) SetPreference(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPreference", reflect.TypeOf((*MockChainVM)(nil).SetPreference), arg0)
}

// SetState mocks base method.
func (m *MockChainVM) SetState(state snow.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", state)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState.
func (mr *MockChainVMMockRecorder) SetState(state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockChainVM)(nil).SetState), state)
}

// Shutdown mocks base method.
func (m *MockChainVM) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockChainVMMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockChainVM)(nil).Shutdown))
}

// Version mocks base method.
func (m *MockChainVM) Version() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockChainVMMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockChainVM)(nil).Version))
}

// MockGetter is a mock of Getter interface.
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetterMockRecorder
}

// MockGetterMockRecorder is the mock recorder for MockGetter.
type MockGetterMockRecorder struct {
	mock *MockGetter
}

// NewMockGetter creates a new mock instance.
func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &MockGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetter) EXPECT() *MockGetterMockRecorder {
	return m.recorder
}

// GetBlock mocks base method.
func (m *MockGetter) GetBlock(arg0 ids.ID) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockGetterMockRecorder) GetBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockGetter)(nil).GetBlock), arg0)
}

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// ParseBlock mocks base method.
func (m *MockParser) ParseBlock(arg0 []byte) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseBlock indicates an expected call of ParseBlock.
func (mr *MockParserMockRecorder) ParseBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBlock", reflect.TypeOf((*MockParser)(nil).ParseBlock), arg0)
}
