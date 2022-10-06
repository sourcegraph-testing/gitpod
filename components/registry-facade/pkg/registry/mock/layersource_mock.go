// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gitpod-io/gitpod/registry-facade/pkg/registry (interfaces: LayerSource)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	api "github.com/gitpod-io/gitpod/registry-facade/api"
	registry "github.com/gitpod-io/gitpod/registry-facade/pkg/registry"
	gomock "github.com/golang/mock/gomock"
	digest "github.com/opencontainers/go-digest"
)

// MockLayerSource is a mock of LayerSource interface.
type MockLayerSource struct {
	ctrl     *gomock.Controller
	recorder *MockLayerSourceMockRecorder
}

// MockLayerSourceMockRecorder is the mock recorder for MockLayerSource.
type MockLayerSourceMockRecorder struct {
	mock *MockLayerSource
}

// NewMockLayerSource creates a new mock instance.
func NewMockLayerSource(ctrl *gomock.Controller) *MockLayerSource {
	mock := &MockLayerSource{ctrl: ctrl}
	mock.recorder = &MockLayerSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLayerSource) EXPECT() *MockLayerSourceMockRecorder {
	return m.recorder
}

// Envs mocks base method.
func (m *MockLayerSource) Envs(arg0 context.Context, arg1 *api.ImageSpec) ([]registry.EnvModifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Envs", arg0, arg1)
	ret0, _ := ret[0].([]registry.EnvModifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Envs indicates an expected call of Envs.
func (mr *MockLayerSourceMockRecorder) Envs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Envs", reflect.TypeOf((*MockLayerSource)(nil).Envs), arg0, arg1)
}

// GetBlob mocks base method.
func (m *MockLayerSource) GetBlob(arg0 context.Context, arg1 *api.ImageSpec, arg2 digest.Digest) (bool, string, string, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[0].(string)
	ret2, _ := ret[1].(string)
	ret3, _ := ret[2].(io.ReadCloser)
	ret4, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetBlob indicates an expected call of GetBlob.
func (mr *MockLayerSourceMockRecorder) GetBlob(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlob", reflect.TypeOf((*MockLayerSource)(nil).GetBlob), arg0, arg1, arg2)
}

// GetLayer mocks base method.
func (m *MockLayerSource) GetLayer(arg0 context.Context, arg1 *api.ImageSpec) ([]registry.AddonLayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLayer", arg0, arg1)
	ret0, _ := ret[0].([]registry.AddonLayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLayer indicates an expected call of GetLayer.
func (mr *MockLayerSourceMockRecorder) GetLayer(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayer", reflect.TypeOf((*MockLayerSource)(nil).GetLayer), arg0, arg1)
}

// HasBlob mocks base method.
func (m *MockLayerSource) HasBlob(arg0 context.Context, arg1 *api.ImageSpec, arg2 digest.Digest) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBlob", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBlob indicates an expected call of HasBlob.
func (mr *MockLayerSourceMockRecorder) HasBlob(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBlob", reflect.TypeOf((*MockLayerSource)(nil).HasBlob), arg0, arg1, arg2)
}
