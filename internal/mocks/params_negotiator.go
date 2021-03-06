// Code generated by MockGen. DO NOT EDIT.
// Source: ../handshake/params_negotiator_base.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
)

// MockParamsNegotiator is a mock of ParamsNegotiator interface
type MockParamsNegotiator struct {
	ctrl     *gomock.Controller
	recorder *MockParamsNegotiatorMockRecorder
}

// MockParamsNegotiatorMockRecorder is the mock recorder for MockParamsNegotiator
type MockParamsNegotiatorMockRecorder struct {
	mock *MockParamsNegotiator
}

// NewMockParamsNegotiator creates a new mock instance
func NewMockParamsNegotiator(ctrl *gomock.Controller) *MockParamsNegotiator {
	mock := &MockParamsNegotiator{ctrl: ctrl}
	mock.recorder = &MockParamsNegotiatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParamsNegotiator) EXPECT() *MockParamsNegotiatorMockRecorder {
	return m.recorder
}

// GetSendStreamFlowControlWindow mocks base method
func (m *MockParamsNegotiator) GetSendStreamFlowControlWindow() protocol.ByteCount {
	ret := m.ctrl.Call(m, "GetSendStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetSendStreamFlowControlWindow indicates an expected call of GetSendStreamFlowControlWindow
func (mr *MockParamsNegotiatorMockRecorder) GetSendStreamFlowControlWindow() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSendStreamFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetSendStreamFlowControlWindow))
}

// GetSendConnectionFlowControlWindow mocks base method
func (m *MockParamsNegotiator) GetSendConnectionFlowControlWindow() protocol.ByteCount {
	ret := m.ctrl.Call(m, "GetSendConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetSendConnectionFlowControlWindow indicates an expected call of GetSendConnectionFlowControlWindow
func (mr *MockParamsNegotiatorMockRecorder) GetSendConnectionFlowControlWindow() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSendConnectionFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetSendConnectionFlowControlWindow))
}

// GetReceiveStreamFlowControlWindow mocks base method
func (m *MockParamsNegotiator) GetReceiveStreamFlowControlWindow() protocol.ByteCount {
	ret := m.ctrl.Call(m, "GetReceiveStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetReceiveStreamFlowControlWindow indicates an expected call of GetReceiveStreamFlowControlWindow
func (mr *MockParamsNegotiatorMockRecorder) GetReceiveStreamFlowControlWindow() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiveStreamFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetReceiveStreamFlowControlWindow))
}

// GetReceiveConnectionFlowControlWindow mocks base method
func (m *MockParamsNegotiator) GetReceiveConnectionFlowControlWindow() protocol.ByteCount {
	ret := m.ctrl.Call(m, "GetReceiveConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetReceiveConnectionFlowControlWindow indicates an expected call of GetReceiveConnectionFlowControlWindow
func (mr *MockParamsNegotiatorMockRecorder) GetReceiveConnectionFlowControlWindow() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiveConnectionFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetReceiveConnectionFlowControlWindow))
}

// GetMaxOutgoingStreams mocks base method
func (m *MockParamsNegotiator) GetMaxOutgoingStreams() uint32 {
	ret := m.ctrl.Call(m, "GetMaxOutgoingStreams")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetMaxOutgoingStreams indicates an expected call of GetMaxOutgoingStreams
func (mr *MockParamsNegotiatorMockRecorder) GetMaxOutgoingStreams() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxOutgoingStreams", reflect.TypeOf((*MockParamsNegotiator)(nil).GetMaxOutgoingStreams))
}

// GetMaxIncomingStreams mocks base method
func (m *MockParamsNegotiator) GetMaxIncomingStreams() uint32 {
	ret := m.ctrl.Call(m, "GetMaxIncomingStreams")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetMaxIncomingStreams indicates an expected call of GetMaxIncomingStreams
func (mr *MockParamsNegotiatorMockRecorder) GetMaxIncomingStreams() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxIncomingStreams", reflect.TypeOf((*MockParamsNegotiator)(nil).GetMaxIncomingStreams))
}

// GetRemoteIdleTimeout mocks base method
func (m *MockParamsNegotiator) GetRemoteIdleTimeout() time.Duration {
	ret := m.ctrl.Call(m, "GetRemoteIdleTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetRemoteIdleTimeout indicates an expected call of GetRemoteIdleTimeout
func (mr *MockParamsNegotiatorMockRecorder) GetRemoteIdleTimeout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteIdleTimeout", reflect.TypeOf((*MockParamsNegotiator)(nil).GetRemoteIdleTimeout))
}

// OmitConnectionID mocks base method
func (m *MockParamsNegotiator) OmitConnectionID() bool {
	ret := m.ctrl.Call(m, "OmitConnectionID")
	ret0, _ := ret[0].(bool)
	return ret0
}

// OmitConnectionID indicates an expected call of OmitConnectionID
func (mr *MockParamsNegotiatorMockRecorder) OmitConnectionID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OmitConnectionID", reflect.TypeOf((*MockParamsNegotiator)(nil).OmitConnectionID))
}
