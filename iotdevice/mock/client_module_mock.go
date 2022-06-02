// Code generated by MockGen. DO NOT EDIT.
// Source: ./iotdevice/client_module.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	iotdevice "github.com/amenzhinsky/iothub/iotdevice"
	gomock "github.com/golang/mock/gomock"
)

// MockModuleClient is a mock of ModuleClient interface.
type MockModuleClient struct {
	ctrl     *gomock.Controller
	recorder *MockModuleClientMockRecorder
}

// MockModuleClientMockRecorder is the mock recorder for MockModuleClient.
type MockModuleClientMockRecorder struct {
	mock *MockModuleClient
}

// NewMockModuleClient creates a new mock instance.
func NewMockModuleClient(ctrl *gomock.Controller) *MockModuleClient {
	mock := &MockModuleClient{ctrl: ctrl}
	mock.recorder = &MockModuleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModuleClient) EXPECT() *MockModuleClientMockRecorder {
	return m.recorder
}

// Broker mocks base method.
func (m *MockModuleClient) Broker() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broker")
	ret0, _ := ret[0].(string)
	return ret0
}

// Broker indicates an expected call of Broker.
func (mr *MockModuleClientMockRecorder) Broker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broker", reflect.TypeOf((*MockModuleClient)(nil).Broker))
}

// Close mocks base method.
func (m *MockModuleClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockModuleClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModuleClient)(nil).Close))
}

// Connect mocks base method.
func (m *MockModuleClient) Connect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockModuleClientMockRecorder) Connect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockModuleClient)(nil).Connect), ctx)
}

// DeviceID mocks base method.
func (m *MockModuleClient) DeviceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeviceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// DeviceID indicates an expected call of DeviceID.
func (mr *MockModuleClientMockRecorder) DeviceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceID", reflect.TypeOf((*MockModuleClient)(nil).DeviceID))
}

// Gateway mocks base method.
func (m *MockModuleClient) Gateway() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gateway")
	ret0, _ := ret[0].(string)
	return ret0
}

// Gateway indicates an expected call of Gateway.
func (mr *MockModuleClientMockRecorder) Gateway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gateway", reflect.TypeOf((*MockModuleClient)(nil).Gateway))
}

// GenerationID mocks base method.
func (m *MockModuleClient) GenerationID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerationID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerationID indicates an expected call of GenerationID.
func (mr *MockModuleClientMockRecorder) GenerationID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerationID", reflect.TypeOf((*MockModuleClient)(nil).GenerationID))
}

// ModuleID mocks base method.
func (m *MockModuleClient) ModuleID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModuleID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModuleID indicates an expected call of ModuleID.
func (mr *MockModuleClientMockRecorder) ModuleID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModuleID", reflect.TypeOf((*MockModuleClient)(nil).ModuleID))
}

// RegisterMethod mocks base method.
func (m *MockModuleClient) RegisterMethod(ctx context.Context, name string, fn iotdevice.DirectMethodHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMethod", ctx, name, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMethod indicates an expected call of RegisterMethod.
func (mr *MockModuleClientMockRecorder) RegisterMethod(ctx, name, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMethod", reflect.TypeOf((*MockModuleClient)(nil).RegisterMethod), ctx, name, fn)
}

// RetrieveTwinState mocks base method.
func (m *MockModuleClient) RetrieveTwinState(ctx context.Context) (iotdevice.TwinState, iotdevice.TwinState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveTwinState", ctx)
	ret0, _ := ret[0].(iotdevice.TwinState)
	ret1, _ := ret[1].(iotdevice.TwinState)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetrieveTwinState indicates an expected call of RetrieveTwinState.
func (mr *MockModuleClientMockRecorder) RetrieveTwinState(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveTwinState", reflect.TypeOf((*MockModuleClient)(nil).RetrieveTwinState), ctx)
}

// SendEvent mocks base method.
func (m *MockModuleClient) SendEvent(ctx context.Context, payload []byte, opts ...iotdevice.SendOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendEvent", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEvent indicates an expected call of SendEvent.
func (mr *MockModuleClientMockRecorder) SendEvent(ctx, payload interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvent", reflect.TypeOf((*MockModuleClient)(nil).SendEvent), varargs...)
}

// SubscribeEvents mocks base method.
func (m *MockModuleClient) SubscribeEvents(ctx context.Context) (*iotdevice.EventSub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeEvents", ctx)
	ret0, _ := ret[0].(*iotdevice.EventSub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeEvents indicates an expected call of SubscribeEvents.
func (mr *MockModuleClientMockRecorder) SubscribeEvents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEvents", reflect.TypeOf((*MockModuleClient)(nil).SubscribeEvents), ctx)
}

// SubscribeTwinUpdates mocks base method.
func (m *MockModuleClient) SubscribeTwinUpdates(ctx context.Context) (*iotdevice.TwinStateSub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeTwinUpdates", ctx)
	ret0, _ := ret[0].(*iotdevice.TwinStateSub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeTwinUpdates indicates an expected call of SubscribeTwinUpdates.
func (mr *MockModuleClientMockRecorder) SubscribeTwinUpdates(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeTwinUpdates", reflect.TypeOf((*MockModuleClient)(nil).SubscribeTwinUpdates), ctx)
}

// UnregisterMethod mocks base method.
func (m *MockModuleClient) UnregisterMethod(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnregisterMethod", name)
}

// UnregisterMethod indicates an expected call of UnregisterMethod.
func (mr *MockModuleClientMockRecorder) UnregisterMethod(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterMethod", reflect.TypeOf((*MockModuleClient)(nil).UnregisterMethod), name)
}

// UnsubscribeEvents mocks base method.
func (m *MockModuleClient) UnsubscribeEvents(sub *iotdevice.EventSub) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribeEvents", sub)
}

// UnsubscribeEvents indicates an expected call of UnsubscribeEvents.
func (mr *MockModuleClientMockRecorder) UnsubscribeEvents(sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeEvents", reflect.TypeOf((*MockModuleClient)(nil).UnsubscribeEvents), sub)
}

// UnsubscribeTwinUpdates mocks base method.
func (m *MockModuleClient) UnsubscribeTwinUpdates(sub *iotdevice.TwinStateSub) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnsubscribeTwinUpdates", sub)
}

// UnsubscribeTwinUpdates indicates an expected call of UnsubscribeTwinUpdates.
func (mr *MockModuleClientMockRecorder) UnsubscribeTwinUpdates(sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeTwinUpdates", reflect.TypeOf((*MockModuleClient)(nil).UnsubscribeTwinUpdates), sub)
}

// UpdateTwinState mocks base method.
func (m *MockModuleClient) UpdateTwinState(ctx context.Context, s iotdevice.TwinState) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTwinState", ctx, s)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTwinState indicates an expected call of UpdateTwinState.
func (mr *MockModuleClientMockRecorder) UpdateTwinState(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTwinState", reflect.TypeOf((*MockModuleClient)(nil).UpdateTwinState), ctx, s)
}

// UploadFile mocks base method.
func (m *MockModuleClient) UploadFile(ctx context.Context, blobName string, file io.Reader, size int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, blobName, file, size)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockModuleClientMockRecorder) UploadFile(ctx, blobName, file, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockModuleClient)(nil).UploadFile), ctx, blobName, file, size)
}
