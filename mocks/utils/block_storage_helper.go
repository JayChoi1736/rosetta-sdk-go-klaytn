// Code generated by mockery v1.0.0. DO NOT EDIT.

package utils

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/klaytn/rosetta-sdk-go-klaytn/types"
)

// BlockStorageHelper is an autogenerated mock type for the BlockStorageHelper type
type BlockStorageHelper struct {
	mock.Mock
}

// GetBlockLazy provides a mock function with given fields: ctx, blockIdentifier
func (_m *BlockStorageHelper) GetBlockLazy(ctx context.Context, blockIdentifier *types.PartialBlockIdentifier) (*types.BlockResponse, error) {
	ret := _m.Called(ctx, blockIdentifier)

	var r0 *types.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.PartialBlockIdentifier) *types.BlockResponse); ok {
		r0 = rf(ctx, blockIdentifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.PartialBlockIdentifier) error); ok {
		r1 = rf(ctx, blockIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
