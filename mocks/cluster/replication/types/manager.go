//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by mockery v2.53.2. DO NOT EDIT.

package types

import (
	mock "github.com/stretchr/testify/mock"
	api "github.com/weaviate/weaviate/cluster/proto/api"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

type Manager_Expecter struct {
	mock *mock.Mock
}

func (_m *Manager) EXPECT() *Manager_Expecter {
	return &Manager_Expecter{mock: &_m.Mock}
}

// GetReplicationDetailsByReplicationId provides a mock function with given fields: id
func (_m *Manager) GetReplicationDetailsByReplicationId(id uint64) (api.ReplicationDetailsResponse, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetReplicationDetailsByReplicationId")
	}

	var r0 api.ReplicationDetailsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (api.ReplicationDetailsResponse, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) api.ReplicationDetailsResponse); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(api.ReplicationDetailsResponse)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_GetReplicationDetailsByReplicationId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReplicationDetailsByReplicationId'
type Manager_GetReplicationDetailsByReplicationId_Call struct {
	*mock.Call
}

// GetReplicationDetailsByReplicationId is a helper method to define mock.On call
//   - id uint64
func (_e *Manager_Expecter) GetReplicationDetailsByReplicationId(id interface{}) *Manager_GetReplicationDetailsByReplicationId_Call {
	return &Manager_GetReplicationDetailsByReplicationId_Call{Call: _e.mock.On("GetReplicationDetailsByReplicationId", id)}
}

func (_c *Manager_GetReplicationDetailsByReplicationId_Call) Run(run func(id uint64)) *Manager_GetReplicationDetailsByReplicationId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *Manager_GetReplicationDetailsByReplicationId_Call) Return(_a0 api.ReplicationDetailsResponse, _a1 error) *Manager_GetReplicationDetailsByReplicationId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_GetReplicationDetailsByReplicationId_Call) RunAndReturn(run func(uint64) (api.ReplicationDetailsResponse, error)) *Manager_GetReplicationDetailsByReplicationId_Call {
	_c.Call.Return(run)
	return _c
}

// ReplicationDeleteReplica provides a mock function with given fields: node, collection, shard
func (_m *Manager) ReplicationDeleteReplica(node string, collection string, shard string) error {
	ret := _m.Called(node, collection, shard)

	if len(ret) == 0 {
		panic("no return value specified for ReplicationDeleteReplica")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(node, collection, shard)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_ReplicationDeleteReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplicationDeleteReplica'
type Manager_ReplicationDeleteReplica_Call struct {
	*mock.Call
}

// ReplicationDeleteReplica is a helper method to define mock.On call
//   - node string
//   - collection string
//   - shard string
func (_e *Manager_Expecter) ReplicationDeleteReplica(node interface{}, collection interface{}, shard interface{}) *Manager_ReplicationDeleteReplica_Call {
	return &Manager_ReplicationDeleteReplica_Call{Call: _e.mock.On("ReplicationDeleteReplica", node, collection, shard)}
}

func (_c *Manager_ReplicationDeleteReplica_Call) Run(run func(node string, collection string, shard string)) *Manager_ReplicationDeleteReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Manager_ReplicationDeleteReplica_Call) Return(_a0 error) *Manager_ReplicationDeleteReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_ReplicationDeleteReplica_Call) RunAndReturn(run func(string, string, string) error) *Manager_ReplicationDeleteReplica_Call {
	_c.Call.Return(run)
	return _c
}

// ReplicationDisableReplica provides a mock function with given fields: node, collection, shard
func (_m *Manager) ReplicationDisableReplica(node string, collection string, shard string) error {
	ret := _m.Called(node, collection, shard)

	if len(ret) == 0 {
		panic("no return value specified for ReplicationDisableReplica")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(node, collection, shard)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_ReplicationDisableReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplicationDisableReplica'
type Manager_ReplicationDisableReplica_Call struct {
	*mock.Call
}

// ReplicationDisableReplica is a helper method to define mock.On call
//   - node string
//   - collection string
//   - shard string
func (_e *Manager_Expecter) ReplicationDisableReplica(node interface{}, collection interface{}, shard interface{}) *Manager_ReplicationDisableReplica_Call {
	return &Manager_ReplicationDisableReplica_Call{Call: _e.mock.On("ReplicationDisableReplica", node, collection, shard)}
}

func (_c *Manager_ReplicationDisableReplica_Call) Run(run func(node string, collection string, shard string)) *Manager_ReplicationDisableReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Manager_ReplicationDisableReplica_Call) Return(_a0 error) *Manager_ReplicationDisableReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_ReplicationDisableReplica_Call) RunAndReturn(run func(string, string, string) error) *Manager_ReplicationDisableReplica_Call {
	_c.Call.Return(run)
	return _c
}

// ReplicationReplicateReplica provides a mock function with given fields: sourceNode, sourceCollection, sourceShard, targetNode
func (_m *Manager) ReplicationReplicateReplica(sourceNode string, sourceCollection string, sourceShard string, targetNode string) error {
	ret := _m.Called(sourceNode, sourceCollection, sourceShard, targetNode)

	if len(ret) == 0 {
		panic("no return value specified for ReplicationReplicateReplica")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(sourceNode, sourceCollection, sourceShard, targetNode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_ReplicationReplicateReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplicationReplicateReplica'
type Manager_ReplicationReplicateReplica_Call struct {
	*mock.Call
}

// ReplicationReplicateReplica is a helper method to define mock.On call
//   - sourceNode string
//   - sourceCollection string
//   - sourceShard string
//   - targetNode string
func (_e *Manager_Expecter) ReplicationReplicateReplica(sourceNode interface{}, sourceCollection interface{}, sourceShard interface{}, targetNode interface{}) *Manager_ReplicationReplicateReplica_Call {
	return &Manager_ReplicationReplicateReplica_Call{Call: _e.mock.On("ReplicationReplicateReplica", sourceNode, sourceCollection, sourceShard, targetNode)}
}

func (_c *Manager_ReplicationReplicateReplica_Call) Run(run func(sourceNode string, sourceCollection string, sourceShard string, targetNode string)) *Manager_ReplicationReplicateReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *Manager_ReplicationReplicateReplica_Call) Return(_a0 error) *Manager_ReplicationReplicateReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_ReplicationReplicateReplica_Call) RunAndReturn(run func(string, string, string, string) error) *Manager_ReplicationReplicateReplica_Call {
	_c.Call.Return(run)
	return _c
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
