// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/api_snapshot_simple_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/glooshot/pkg/api/v1"
)

// MockApiSimpleEmitter is a mock of ApiSimpleEmitter interface
type MockApiSimpleEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockApiSimpleEmitterMockRecorder
}

// MockApiSimpleEmitterMockRecorder is the mock recorder for MockApiSimpleEmitter
type MockApiSimpleEmitterMockRecorder struct {
	mock *MockApiSimpleEmitter
}

// NewMockApiSimpleEmitter creates a new mock instance
func NewMockApiSimpleEmitter(ctrl *gomock.Controller) *MockApiSimpleEmitter {
	mock := &MockApiSimpleEmitter{ctrl: ctrl}
	mock.recorder = &MockApiSimpleEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiSimpleEmitter) EXPECT() *MockApiSimpleEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockApiSimpleEmitter) Snapshots(ctx context.Context) (<-chan *v1.ApiSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", ctx)
	ret0, _ := ret[0].(<-chan *v1.ApiSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockApiSimpleEmitterMockRecorder) Snapshots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockApiSimpleEmitter)(nil).Snapshots), ctx)
}
