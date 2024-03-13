// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	abcitypes "github.com/cometbft/cometbft/abci/types"

	header "github.com/celestiaorg/go-header"

	mock "github.com/stretchr/testify/mock"

	types "github.com/rollkit/rollkit/types"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Store) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlock provides a mock function with given fields: ctx, height
func (_m *Store) GetBlock(ctx context.Context, height uint64) (*types.Block, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("no return value specified for GetBlock")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*types.Block, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *types.Block); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHash provides a mock function with given fields: ctx, hash
func (_m *Store) GetBlockByHash(ctx context.Context, hash header.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockByHash")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, header.Hash) (*types.Block, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, header.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, header.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockResponses provides a mock function with given fields: ctx, height
func (_m *Store) GetBlockResponses(ctx context.Context, height uint64) (*abcitypes.ResponseFinalizeBlock, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockResponses")
	}

	var r0 *abcitypes.ResponseFinalizeBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*abcitypes.ResponseFinalizeBlock, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *abcitypes.ResponseFinalizeBlock); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcitypes.ResponseFinalizeBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommit provides a mock function with given fields: ctx, height
func (_m *Store) GetCommit(ctx context.Context, height uint64) (*types.Commit, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("no return value specified for GetCommit")
	}

	var r0 *types.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*types.Commit, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *types.Commit); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommitByHash provides a mock function with given fields: ctx, hash
func (_m *Store) GetCommitByHash(ctx context.Context, hash header.Hash) (*types.Commit, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetCommitByHash")
	}

	var r0 *types.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, header.Hash) (*types.Commit, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, header.Hash) *types.Commit); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, header.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: ctx, key
func (_m *Store) GetMetadata(ctx context.Context, key string) ([]byte, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: ctx
func (_m *Store) GetState(ctx context.Context) (types.State, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 types.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.State, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.State); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(types.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Height provides a mock function with given fields:
func (_m *Store) Height() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Height")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// SaveBlock provides a mock function with given fields: ctx, block, commit
func (_m *Store) SaveBlock(ctx context.Context, block *types.Block, commit *types.Commit) error {
	ret := _m.Called(ctx, block, commit)

	if len(ret) == 0 {
		panic("no return value specified for SaveBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Block, *types.Commit) error); ok {
		r0 = rf(ctx, block, commit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveBlockResponses provides a mock function with given fields: ctx, height, responses
func (_m *Store) SaveBlockResponses(ctx context.Context, height uint64, responses *abcitypes.ResponseFinalizeBlock) error {
	ret := _m.Called(ctx, height, responses)

	if len(ret) == 0 {
		panic("no return value specified for SaveBlockResponses")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *abcitypes.ResponseFinalizeBlock) error); ok {
		r0 = rf(ctx, height, responses)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetHeight provides a mock function with given fields: ctx, height
func (_m *Store) SetHeight(ctx context.Context, height uint64) {
	_m.Called(ctx, height)
}

// SetMetadata provides a mock function with given fields: ctx, key, value
func (_m *Store) SetMetadata(ctx context.Context, key string, value []byte) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for SetMetadata")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateState provides a mock function with given fields: ctx, state
func (_m *Store) UpdateState(ctx context.Context, state types.State) error {
	ret := _m.Called(ctx, state)

	if len(ret) == 0 {
		panic("no return value specified for UpdateState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.State) error); ok {
		r0 = rf(ctx, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
