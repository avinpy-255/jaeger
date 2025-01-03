// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	spanstore "github.com/jaegertracing/jaeger/storage/spanstore"
	mock "github.com/stretchr/testify/mock"
)

// ArchiveFactory is an autogenerated mock type for the ArchiveFactory type
type ArchiveFactory struct {
	mock.Mock
}

// CreateArchiveSpanReader provides a mock function with no fields
func (_m *ArchiveFactory) CreateArchiveSpanReader() (spanstore.Reader, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateArchiveSpanReader")
	}

	var r0 spanstore.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func() (spanstore.Reader, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() spanstore.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spanstore.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateArchiveSpanWriter provides a mock function with no fields
func (_m *ArchiveFactory) CreateArchiveSpanWriter() (spanstore.Writer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateArchiveSpanWriter")
	}

	var r0 spanstore.Writer
	var r1 error
	if rf, ok := ret.Get(0).(func() (spanstore.Writer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() spanstore.Writer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spanstore.Writer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewArchiveFactory creates a new instance of ArchiveFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArchiveFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArchiveFactory {
	mock := &ArchiveFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
