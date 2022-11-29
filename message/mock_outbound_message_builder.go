// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/message (interfaces: OutboundMsgBuilder)

// Package message is a generated GoMock package.
package message

import (
	reflect "reflect"
	time "time"

	ids "github.com/ava-labs/avalanchego/ids"
	ips "github.com/ava-labs/avalanchego/utils/ips"
	gomock "github.com/golang/mock/gomock"
)

// MockOutboundMsgBuilder is a mock of OutboundMsgBuilder interface.
type MockOutboundMsgBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockOutboundMsgBuilderMockRecorder
}

// MockOutboundMsgBuilderMockRecorder is the mock recorder for MockOutboundMsgBuilder.
type MockOutboundMsgBuilderMockRecorder struct {
	mock *MockOutboundMsgBuilder
}

// NewMockOutboundMsgBuilder creates a new mock instance.
func NewMockOutboundMsgBuilder(ctrl *gomock.Controller) *MockOutboundMsgBuilder {
	mock := &MockOutboundMsgBuilder{ctrl: ctrl}
	mock.recorder = &MockOutboundMsgBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutboundMsgBuilder) EXPECT() *MockOutboundMsgBuilderMockRecorder {
	return m.recorder
}

// Accepted mocks base method.
func (m *MockOutboundMsgBuilder) Accepted(arg0 ids.ID, arg1 uint32, arg2 []ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accepted", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accepted indicates an expected call of Accepted.
func (mr *MockOutboundMsgBuilderMockRecorder) Accepted(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accepted", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Accepted), arg0, arg1, arg2)
}

// AcceptedFrontier mocks base method.
func (m *MockOutboundMsgBuilder) AcceptedFrontier(arg0 ids.ID, arg1 uint32, arg2 []ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedFrontier indicates an expected call of AcceptedFrontier.
func (mr *MockOutboundMsgBuilderMockRecorder) AcceptedFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedFrontier", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).AcceptedFrontier), arg0, arg1, arg2)
}

// AcceptedStateSummary mocks base method.
func (m *MockOutboundMsgBuilder) AcceptedStateSummary(arg0 ids.ID, arg1 uint32, arg2 []ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedStateSummary", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptedStateSummary indicates an expected call of AcceptedStateSummary.
func (mr *MockOutboundMsgBuilderMockRecorder) AcceptedStateSummary(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedStateSummary", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).AcceptedStateSummary), arg0, arg1, arg2)
}

// Ancestors mocks base method.
func (m *MockOutboundMsgBuilder) Ancestors(arg0 ids.ID, arg1 uint32, arg2 [][]byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ancestors", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ancestors indicates an expected call of Ancestors.
func (mr *MockOutboundMsgBuilderMockRecorder) Ancestors(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ancestors", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Ancestors), arg0, arg1, arg2)
}

// AppGossip mocks base method.
func (m *MockOutboundMsgBuilder) AppGossip(arg0 ids.ID, arg1 []byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGossip", arg0, arg1)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppGossip indicates an expected call of AppGossip.
func (mr *MockOutboundMsgBuilderMockRecorder) AppGossip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGossip", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).AppGossip), arg0, arg1)
}

// AppRequest mocks base method.
func (m *MockOutboundMsgBuilder) AppRequest(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppRequest indicates an expected call of AppRequest.
func (mr *MockOutboundMsgBuilderMockRecorder) AppRequest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequest", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).AppRequest), arg0, arg1, arg2, arg3)
}

// AppResponse mocks base method.
func (m *MockOutboundMsgBuilder) AppResponse(arg0 ids.ID, arg1 uint32, arg2 []byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppResponse indicates an expected call of AppResponse.
func (mr *MockOutboundMsgBuilderMockRecorder) AppResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).AppResponse), arg0, arg1, arg2)
}

// Chits mocks base method.
func (m *MockOutboundMsgBuilder) Chits(arg0 ids.ID, arg1 uint32, arg2 []ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chits", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chits indicates an expected call of Chits.
func (mr *MockOutboundMsgBuilderMockRecorder) Chits(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chits", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Chits), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockOutboundMsgBuilder) Get(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOutboundMsgBuilderMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetAccepted mocks base method.
func (m *MockOutboundMsgBuilder) GetAccepted(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccepted", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccepted indicates an expected call of GetAccepted.
func (mr *MockOutboundMsgBuilderMockRecorder) GetAccepted(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccepted", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).GetAccepted), arg0, arg1, arg2, arg3)
}

// GetAcceptedFrontier mocks base method.
func (m *MockOutboundMsgBuilder) GetAcceptedFrontier(arg0 ids.ID, arg1 uint32, arg2 time.Duration) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcceptedFrontier indicates an expected call of GetAcceptedFrontier.
func (mr *MockOutboundMsgBuilderMockRecorder) GetAcceptedFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedFrontier", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).GetAcceptedFrontier), arg0, arg1, arg2)
}

// GetAcceptedStateSummary mocks base method.
func (m *MockOutboundMsgBuilder) GetAcceptedStateSummary(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []uint64) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedStateSummary", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcceptedStateSummary indicates an expected call of GetAcceptedStateSummary.
func (mr *MockOutboundMsgBuilderMockRecorder) GetAcceptedStateSummary(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedStateSummary", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).GetAcceptedStateSummary), arg0, arg1, arg2, arg3)
}

// GetAncestors mocks base method.
func (m *MockOutboundMsgBuilder) GetAncestors(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAncestors", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAncestors indicates an expected call of GetAncestors.
func (mr *MockOutboundMsgBuilderMockRecorder) GetAncestors(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAncestors", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).GetAncestors), arg0, arg1, arg2, arg3)
}

// GetStateSummaryFrontier mocks base method.
func (m *MockOutboundMsgBuilder) GetStateSummaryFrontier(arg0 ids.ID, arg1 uint32, arg2 time.Duration) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateSummaryFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateSummaryFrontier indicates an expected call of GetStateSummaryFrontier.
func (mr *MockOutboundMsgBuilderMockRecorder) GetStateSummaryFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateSummaryFrontier", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).GetStateSummaryFrontier), arg0, arg1, arg2)
}

// PeerList mocks base method.
func (m *MockOutboundMsgBuilder) PeerList(arg0 []ips.ClaimedIPPort, arg1 bool) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerList", arg0, arg1)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerList indicates an expected call of PeerList.
func (mr *MockOutboundMsgBuilderMockRecorder) PeerList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerList", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).PeerList), arg0, arg1)
}

// Ping mocks base method.
func (m *MockOutboundMsgBuilder) Ping() (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockOutboundMsgBuilderMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Ping))
}

// Pong mocks base method.
func (m *MockOutboundMsgBuilder) Pong(arg0 byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pong", arg0)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pong indicates an expected call of Pong.
func (mr *MockOutboundMsgBuilderMockRecorder) Pong(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pong", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Pong), arg0)
}

// PullQuery mocks base method.
func (m *MockOutboundMsgBuilder) PullQuery(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullQuery indicates an expected call of PullQuery.
func (mr *MockOutboundMsgBuilderMockRecorder) PullQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullQuery", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).PullQuery), arg0, arg1, arg2, arg3)
}

// PushQuery mocks base method.
func (m *MockOutboundMsgBuilder) PushQuery(arg0 ids.ID, arg1 uint32, arg2 time.Duration, arg3 []byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushQuery indicates an expected call of PushQuery.
func (mr *MockOutboundMsgBuilderMockRecorder) PushQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushQuery", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).PushQuery), arg0, arg1, arg2, arg3)
}

// Put mocks base method.
func (m *MockOutboundMsgBuilder) Put(arg0 ids.ID, arg1 uint32, arg2 []byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockOutboundMsgBuilderMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Put), arg0, arg1, arg2)
}

// StateSummaryFrontier mocks base method.
func (m *MockOutboundMsgBuilder) StateSummaryFrontier(arg0 ids.ID, arg1 uint32, arg2 []byte) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSummaryFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSummaryFrontier indicates an expected call of StateSummaryFrontier.
func (mr *MockOutboundMsgBuilderMockRecorder) StateSummaryFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSummaryFrontier", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).StateSummaryFrontier), arg0, arg1, arg2)
}

// Version mocks base method.
func (m *MockOutboundMsgBuilder) Version(arg0 uint32, arg1 uint64, arg2 ips.IPPort, arg3 string, arg4 uint64, arg5 []byte, arg6 []ids.ID) (OutboundMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(OutboundMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockOutboundMsgBuilderMockRecorder) Version(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockOutboundMsgBuilder)(nil).Version), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
