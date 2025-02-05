// Code generated by mockery v2.2.0. DO NOT EDIT.

package storage

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	url "github.com/kmiku7/s5cmd/storage/url"
)

// MockStorage is an autogenerated mock type for the Storage type
type MockStorage struct {
	mock.Mock
}

// Copy provides a mock function with given fields: ctx, src, dst, metadata
func (_m *MockStorage) Copy(ctx context.Context, src *url.URL, dst *url.URL, metadata Metadata) error {
	ret := _m.Called(ctx, src, dst, metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL, *url.URL, Metadata) error); ok {
		r0 = rf(ctx, src, dst, metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, src
func (_m *MockStorage) Delete(ctx context.Context, src *url.URL) error {
	ret := _m.Called(ctx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL) error); ok {
		r0 = rf(ctx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, src, followSymlinks
func (_m *MockStorage) List(ctx context.Context, src *url.URL, followSymlinks bool) <-chan *Object {
	ret := _m.Called(ctx, src, followSymlinks)

	var r0 <-chan *Object
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL, bool) <-chan *Object); ok {
		r0 = rf(ctx, src, followSymlinks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *Object)
		}
	}

	return r0
}

// MultiDelete provides a mock function with given fields: ctx, urls
func (_m *MockStorage) MultiDelete(ctx context.Context, urls <-chan *url.URL) <-chan *Object {
	ret := _m.Called(ctx, urls)

	var r0 <-chan *Object
	if rf, ok := ret.Get(0).(func(context.Context, <-chan *url.URL) <-chan *Object); ok {
		r0 = rf(ctx, urls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *Object)
		}
	}

	return r0
}

// Stat provides a mock function with given fields: ctx, src
func (_m *MockStorage) Stat(ctx context.Context, src *url.URL) (*Object, error) {
	ret := _m.Called(ctx, src)

	var r0 *Object
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL) *Object); ok {
		r0 = rf(ctx, src)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *url.URL) error); ok {
		r1 = rf(ctx, src)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
