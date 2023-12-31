// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/v4/proto/eth/service (interfaces: EventsClient,Events_StreamEventsClient,Events_StreamEventsServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gateway "github.com/grpc-ecosystem/grpc-gateway/v2/proto/gateway"
	service "github.com/prysmaticlabs/prysm/v4/proto/eth/service"
	v1 "github.com/prysmaticlabs/prysm/v4/proto/eth/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockEventsClient is a mock of EventsClient interface.
type EventsClient struct {
	ctrl     *gomock.Controller
	recorder *EventsClientMockRecorder
}

// MockEventsClientMockRecorder is the mock recorder for MockEventsClient.
type EventsClientMockRecorder struct {
	mock *EventsClient
}

// NewMockEventsClient creates a new mock instance.
func NewMockEventsClient(ctrl *gomock.Controller) *EventsClient {
	mock := &EventsClient{ctrl: ctrl}
	mock.recorder = &EventsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *EventsClient) EXPECT() *EventsClientMockRecorder {
	return m.recorder
}

// StreamEvents mocks base method.
func (m *EventsClient) StreamEvents(arg0 context.Context, arg1 *v1.StreamEventsRequest, arg2 ...grpc.CallOption) (service.Events_StreamEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamEvents", varargs...)
	ret0, _ := ret[0].(service.Events_StreamEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamEvents indicates an expected call of StreamEvents.
func (mr *EventsClientMockRecorder) StreamEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamEvents", reflect.TypeOf((*EventsClient)(nil).StreamEvents), varargs...)
}

// MockEvents_StreamEventsClient is a mock of Events_StreamEventsClient interface.
type Events_StreamEventsClient struct {
	ctrl     *gomock.Controller
	recorder *Events_StreamEventsClientMockRecorder
}

// MockEvents_StreamEventsClientMockRecorder is the mock recorder for MockEvents_StreamEventsClient.
type Events_StreamEventsClientMockRecorder struct {
	mock *Events_StreamEventsClient
}

// NewMockEvents_StreamEventsClient creates a new mock instance.
func NewMockEvents_StreamEventsClient(ctrl *gomock.Controller) *Events_StreamEventsClient {
	mock := &Events_StreamEventsClient{ctrl: ctrl}
	mock.recorder = &Events_StreamEventsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Events_StreamEventsClient) EXPECT() *Events_StreamEventsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *Events_StreamEventsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *Events_StreamEventsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*Events_StreamEventsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *Events_StreamEventsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *Events_StreamEventsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*Events_StreamEventsClient)(nil).Context))
}

// Header mocks base method.
func (m *Events_StreamEventsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *Events_StreamEventsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*Events_StreamEventsClient)(nil).Header))
}

// Recv mocks base method.
func (m *Events_StreamEventsClient) Recv() (*gateway.EventSource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*gateway.EventSource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *Events_StreamEventsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*Events_StreamEventsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *Events_StreamEventsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *Events_StreamEventsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*Events_StreamEventsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *Events_StreamEventsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *Events_StreamEventsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*Events_StreamEventsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *Events_StreamEventsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *Events_StreamEventsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*Events_StreamEventsClient)(nil).Trailer))
}

// MockEvents_StreamEventsServer is a mock of Events_StreamEventsServer interface.
type Events_StreamEventsServer struct {
	ctrl     *gomock.Controller
	recorder *Events_StreamEventsServerMockRecorder
}

// MockEvents_StreamEventsServerMockRecorder is the mock recorder for MockEvents_StreamEventsServer.
type Events_StreamEventsServerMockRecorder struct {
	mock *Events_StreamEventsServer
}

// NewMockEvents_StreamEventsServer creates a new mock instance.
func NewMockEvents_StreamEventsServer(ctrl *gomock.Controller) *Events_StreamEventsServer {
	mock := &Events_StreamEventsServer{ctrl: ctrl}
	mock.recorder = &Events_StreamEventsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Events_StreamEventsServer) EXPECT() *Events_StreamEventsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *Events_StreamEventsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *Events_StreamEventsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*Events_StreamEventsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *Events_StreamEventsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *Events_StreamEventsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*Events_StreamEventsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *Events_StreamEventsServer) Send(arg0 *gateway.EventSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *Events_StreamEventsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*Events_StreamEventsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *Events_StreamEventsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *Events_StreamEventsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*Events_StreamEventsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *Events_StreamEventsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *Events_StreamEventsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*Events_StreamEventsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *Events_StreamEventsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *Events_StreamEventsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*Events_StreamEventsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *Events_StreamEventsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *Events_StreamEventsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*Events_StreamEventsServer)(nil).SetTrailer), arg0)
}
