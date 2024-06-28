// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/ssclient/ssclient.go
//
// Generated by this command:
//
//	mockgen -typed -source=./internal/ssclient/ssclient.go -destination=./internal/ssclient/ssclientmock/mock_ssclient.go -package=ssclientmock Client
//

// Package ssclientmock is a generated GoMock package.
package ssclientmock

import (
	context "context"
	reflect "reflect"

	ssclient "github.com/artefactual/archivematica/hack/ccp/internal/ssclient"
	enums "github.com/artefactual/archivematica/hack/ccp/internal/ssclient/enums"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListLocations mocks base method.
func (m *MockClient) ListLocations(ctx context.Context, path string, purpose enums.LocationPurpose) ([]*ssclient.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", ctx, path, purpose)
	ret0, _ := ret[0].([]*ssclient.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockClientMockRecorder) ListLocations(ctx, path, purpose any) *MockClientListLocationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockClient)(nil).ListLocations), ctx, path, purpose)
	return &MockClientListLocationsCall{Call: call}
}

// MockClientListLocationsCall wrap *gomock.Call
type MockClientListLocationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientListLocationsCall) Return(arg0 []*ssclient.Location, arg1 error) *MockClientListLocationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientListLocationsCall) Do(f func(context.Context, string, enums.LocationPurpose) ([]*ssclient.Location, error)) *MockClientListLocationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientListLocationsCall) DoAndReturn(f func(context.Context, string, enums.LocationPurpose) ([]*ssclient.Location, error)) *MockClientListLocationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MoveFiles mocks base method.
func (m *MockClient) MoveFiles(ctx context.Context, src, dst *ssclient.Location, files [][2]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveFiles", ctx, src, dst, files)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveFiles indicates an expected call of MoveFiles.
func (mr *MockClientMockRecorder) MoveFiles(ctx, src, dst, files any) *MockClientMoveFilesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveFiles", reflect.TypeOf((*MockClient)(nil).MoveFiles), ctx, src, dst, files)
	return &MockClientMoveFilesCall{Call: call}
}

// MockClientMoveFilesCall wrap *gomock.Call
type MockClientMoveFilesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientMoveFilesCall) Return(arg0 error) *MockClientMoveFilesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientMoveFilesCall) Do(f func(context.Context, *ssclient.Location, *ssclient.Location, [][2]string) error) *MockClientMoveFilesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientMoveFilesCall) DoAndReturn(f func(context.Context, *ssclient.Location, *ssclient.Location, [][2]string) error) *MockClientMoveFilesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadDefaultLocation mocks base method.
func (m *MockClient) ReadDefaultLocation(ctx context.Context, purpose enums.LocationPurpose) (*ssclient.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDefaultLocation", ctx, purpose)
	ret0, _ := ret[0].(*ssclient.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDefaultLocation indicates an expected call of ReadDefaultLocation.
func (mr *MockClientMockRecorder) ReadDefaultLocation(ctx, purpose any) *MockClientReadDefaultLocationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDefaultLocation", reflect.TypeOf((*MockClient)(nil).ReadDefaultLocation), ctx, purpose)
	return &MockClientReadDefaultLocationCall{Call: call}
}

// MockClientReadDefaultLocationCall wrap *gomock.Call
type MockClientReadDefaultLocationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientReadDefaultLocationCall) Return(arg0 *ssclient.Location, arg1 error) *MockClientReadDefaultLocationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientReadDefaultLocationCall) Do(f func(context.Context, enums.LocationPurpose) (*ssclient.Location, error)) *MockClientReadDefaultLocationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientReadDefaultLocationCall) DoAndReturn(f func(context.Context, enums.LocationPurpose) (*ssclient.Location, error)) *MockClientReadDefaultLocationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadPipeline mocks base method.
func (m *MockClient) ReadPipeline(ctx context.Context, id uuid.UUID) (*ssclient.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPipeline", ctx, id)
	ret0, _ := ret[0].(*ssclient.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPipeline indicates an expected call of ReadPipeline.
func (mr *MockClientMockRecorder) ReadPipeline(ctx, id any) *MockClientReadPipelineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPipeline", reflect.TypeOf((*MockClient)(nil).ReadPipeline), ctx, id)
	return &MockClientReadPipelineCall{Call: call}
}

// MockClientReadPipelineCall wrap *gomock.Call
type MockClientReadPipelineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientReadPipelineCall) Return(arg0 *ssclient.Pipeline, arg1 error) *MockClientReadPipelineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientReadPipelineCall) Do(f func(context.Context, uuid.UUID) (*ssclient.Pipeline, error)) *MockClientReadPipelineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientReadPipelineCall) DoAndReturn(f func(context.Context, uuid.UUID) (*ssclient.Pipeline, error)) *MockClientReadPipelineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadProcessingLocation mocks base method.
func (m *MockClient) ReadProcessingLocation(ctx context.Context) (*ssclient.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProcessingLocation", ctx)
	ret0, _ := ret[0].(*ssclient.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProcessingLocation indicates an expected call of ReadProcessingLocation.
func (mr *MockClientMockRecorder) ReadProcessingLocation(ctx any) *MockClientReadProcessingLocationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProcessingLocation", reflect.TypeOf((*MockClient)(nil).ReadProcessingLocation), ctx)
	return &MockClientReadProcessingLocationCall{Call: call}
}

// MockClientReadProcessingLocationCall wrap *gomock.Call
type MockClientReadProcessingLocationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientReadProcessingLocationCall) Return(arg0 *ssclient.Location, arg1 error) *MockClientReadProcessingLocationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientReadProcessingLocationCall) Do(f func(context.Context) (*ssclient.Location, error)) *MockClientReadProcessingLocationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientReadProcessingLocationCall) DoAndReturn(f func(context.Context) (*ssclient.Location, error)) *MockClientReadProcessingLocationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}