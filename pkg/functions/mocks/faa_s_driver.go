// Code generated by mockery v1.0.0
package mocks

import functions "gitlab.eng.vmware.com/serverless/serverless/pkg/functions"
import mock "github.com/stretchr/testify/mock"

// FaaSDriver is an autogenerated mock type for the FaaSDriver type
type FaaSDriver struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, exec
func (_m *FaaSDriver) Create(name string, exec *functions.Exec) error {
	ret := _m.Called(name, exec)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *functions.Exec) error); ok {
		r0 = rf(name, exec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: name
func (_m *FaaSDriver) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRunnable provides a mock function with given fields: name
func (_m *FaaSDriver) GetRunnable(name string) functions.Runnable {
	ret := _m.Called(name)

	var r0 functions.Runnable
	if rf, ok := ret.Get(0).(func(string) functions.Runnable); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(functions.Runnable)
		}
	}

	return r0
}