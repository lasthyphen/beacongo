// Code generated by MockGen. DO NOT EDIT.
// Source: utils/logging/logger.go

// Package logging is a generated GoMock package.
package logging

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// AssertDeferredNoError mocks base method.
func (m *MockLogger) AssertDeferredNoError(f func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AssertDeferredNoError", f)
}

// AssertDeferredNoError indicates an expected call of AssertDeferredNoError.
func (mr *MockLoggerMockRecorder) AssertDeferredNoError(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssertDeferredNoError", reflect.TypeOf((*MockLogger)(nil).AssertDeferredNoError), f)
}

// AssertDeferredTrue mocks base method.
func (m *MockLogger) AssertDeferredTrue(f func() bool, format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{f, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AssertDeferredTrue", varargs...)
}

// AssertDeferredTrue indicates an expected call of AssertDeferredTrue.
func (mr *MockLoggerMockRecorder) AssertDeferredTrue(f, format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{f, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssertDeferredTrue", reflect.TypeOf((*MockLogger)(nil).AssertDeferredTrue), varargs...)
}

// AssertNoError mocks base method.
func (m *MockLogger) AssertNoError(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AssertNoError", err)
}

// AssertNoError indicates an expected call of AssertNoError.
func (mr *MockLoggerMockRecorder) AssertNoError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssertNoError", reflect.TypeOf((*MockLogger)(nil).AssertNoError), err)
}

// AssertTrue mocks base method.
func (m *MockLogger) AssertTrue(b bool, format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{b, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AssertTrue", varargs...)
}

// AssertTrue indicates an expected call of AssertTrue.
func (mr *MockLoggerMockRecorder) AssertTrue(b, format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{b, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssertTrue", reflect.TypeOf((*MockLogger)(nil).AssertTrue), varargs...)
}

// Debug mocks base method.
func (m *MockLogger) Debug(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), varargs...)
}

// Error mocks base method.
func (m *MockLogger) Error(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), varargs...)
}

// Fatal mocks base method.
func (m *MockLogger) Fatal(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder) Fatal(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), varargs...)
}

// GetDisplayLevel mocks base method.
func (m *MockLogger) GetDisplayLevel() Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDisplayLevel")
	ret0, _ := ret[0].(Level)
	return ret0
}

// GetDisplayLevel indicates an expected call of GetDisplayLevel.
func (mr *MockLoggerMockRecorder) GetDisplayLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisplayLevel", reflect.TypeOf((*MockLogger)(nil).GetDisplayLevel))
}

// GetLogLevel mocks base method.
func (m *MockLogger) GetLogLevel() Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogLevel")
	ret0, _ := ret[0].(Level)
	return ret0
}

// GetLogLevel indicates an expected call of GetLogLevel.
func (mr *MockLoggerMockRecorder) GetLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogLevel", reflect.TypeOf((*MockLogger)(nil).GetLogLevel))
}

// Info mocks base method.
func (m *MockLogger) Info(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), varargs...)
}

// RecoverAndExit mocks base method.
func (m *MockLogger) RecoverAndExit(f, exit func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecoverAndExit", f, exit)
}

// RecoverAndExit indicates an expected call of RecoverAndExit.
func (mr *MockLoggerMockRecorder) RecoverAndExit(f, exit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverAndExit", reflect.TypeOf((*MockLogger)(nil).RecoverAndExit), f, exit)
}

// RecoverAndPanic mocks base method.
func (m *MockLogger) RecoverAndPanic(f func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecoverAndPanic", f)
}

// RecoverAndPanic indicates an expected call of RecoverAndPanic.
func (mr *MockLoggerMockRecorder) RecoverAndPanic(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverAndPanic", reflect.TypeOf((*MockLogger)(nil).RecoverAndPanic), f)
}

// SetContextualDisplayingEnabled mocks base method.
func (m *MockLogger) SetContextualDisplayingEnabled(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContextualDisplayingEnabled", arg0)
}

// SetContextualDisplayingEnabled indicates an expected call of SetContextualDisplayingEnabled.
func (mr *MockLoggerMockRecorder) SetContextualDisplayingEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContextualDisplayingEnabled", reflect.TypeOf((*MockLogger)(nil).SetContextualDisplayingEnabled), arg0)
}

// SetDisplayLevel mocks base method.
func (m *MockLogger) SetDisplayLevel(arg0 Level) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDisplayLevel", arg0)
}

// SetDisplayLevel indicates an expected call of SetDisplayLevel.
func (mr *MockLoggerMockRecorder) SetDisplayLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDisplayLevel", reflect.TypeOf((*MockLogger)(nil).SetDisplayLevel), arg0)
}

// SetDisplayingEnabled mocks base method.
func (m *MockLogger) SetDisplayingEnabled(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDisplayingEnabled", arg0)
}

// SetDisplayingEnabled indicates an expected call of SetDisplayingEnabled.
func (mr *MockLoggerMockRecorder) SetDisplayingEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDisplayingEnabled", reflect.TypeOf((*MockLogger)(nil).SetDisplayingEnabled), arg0)
}

// SetLogLevel mocks base method.
func (m *MockLogger) SetLogLevel(arg0 Level) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", arg0)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockLoggerMockRecorder) SetLogLevel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockLogger)(nil).SetLogLevel), arg0)
}

// SetLoggingEnabled mocks base method.
func (m *MockLogger) SetLoggingEnabled(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLoggingEnabled", arg0)
}

// SetLoggingEnabled indicates an expected call of SetLoggingEnabled.
func (mr *MockLoggerMockRecorder) SetLoggingEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoggingEnabled", reflect.TypeOf((*MockLogger)(nil).SetLoggingEnabled), arg0)
}

// SetPrefix mocks base method.
func (m *MockLogger) SetPrefix(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPrefix", arg0)
}

// SetPrefix indicates an expected call of SetPrefix.
func (mr *MockLoggerMockRecorder) SetPrefix(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrefix", reflect.TypeOf((*MockLogger)(nil).SetPrefix), arg0)
}

// Stop mocks base method.
func (m *MockLogger) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockLoggerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockLogger)(nil).Stop))
}

// StopOnPanic mocks base method.
func (m *MockLogger) StopOnPanic() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopOnPanic")
}

// StopOnPanic indicates an expected call of StopOnPanic.
func (mr *MockLoggerMockRecorder) StopOnPanic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopOnPanic", reflect.TypeOf((*MockLogger)(nil).StopOnPanic))
}

// Trace mocks base method.
func (m *MockLogger) Trace(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Trace", varargs...)
}

// Trace indicates an expected call of Trace.
func (mr *MockLoggerMockRecorder) Trace(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockLogger)(nil).Trace), varargs...)
}

// Verbo mocks base method.
func (m *MockLogger) Verbo(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Verbo", varargs...)
}

// Verbo indicates an expected call of Verbo.
func (mr *MockLoggerMockRecorder) Verbo(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verbo", reflect.TypeOf((*MockLogger)(nil).Verbo), varargs...)
}

// Warn mocks base method.
func (m *MockLogger) Warn(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), varargs...)
}

// Write mocks base method.
func (m *MockLogger) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockLoggerMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockLogger)(nil).Write), p)
}

// MockRotatingWriter is a mock of RotatingWriter interface.
type MockRotatingWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRotatingWriterMockRecorder
}

// MockRotatingWriterMockRecorder is the mock recorder for MockRotatingWriter.
type MockRotatingWriterMockRecorder struct {
	mock *MockRotatingWriter
}

// NewMockRotatingWriter creates a new mock instance.
func NewMockRotatingWriter(ctrl *gomock.Controller) *MockRotatingWriter {
	mock := &MockRotatingWriter{ctrl: ctrl}
	mock.recorder = &MockRotatingWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRotatingWriter) EXPECT() *MockRotatingWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRotatingWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRotatingWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRotatingWriter)(nil).Close))
}

// Flush mocks base method.
func (m *MockRotatingWriter) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockRotatingWriterMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockRotatingWriter)(nil).Flush))
}

// Initialize mocks base method.
func (m *MockRotatingWriter) Initialize(arg0 Config) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Initialize indicates an expected call of Initialize.
func (mr *MockRotatingWriterMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockRotatingWriter)(nil).Initialize), arg0)
}

// Rotate mocks base method.
func (m *MockRotatingWriter) Rotate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rotate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rotate indicates an expected call of Rotate.
func (mr *MockRotatingWriterMockRecorder) Rotate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rotate", reflect.TypeOf((*MockRotatingWriter)(nil).Rotate))
}

// Write mocks base method.
func (m *MockRotatingWriter) Write(b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockRotatingWriterMockRecorder) Write(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockRotatingWriter)(nil).Write), b)
}

// WriteString mocks base method.
func (m *MockRotatingWriter) WriteString(s string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", s)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString.
func (mr *MockRotatingWriterMockRecorder) WriteString(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockRotatingWriter)(nil).WriteString), s)
}
