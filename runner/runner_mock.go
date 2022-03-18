// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package runner

import mock "github.com/stretchr/testify/mock"

// MockRunnerService is an autogenerated mock type for the RunnerService type
type MockRunnerService struct {
	mock.Mock
}

// BuildCatalog provides a mock function with given fields:
func (_m *MockRunnerService) BuildCatalog() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCatalog provides a mock function with given fields:
func (_m *MockRunnerService) GetCatalog() map[string]*Catalog {
	ret := _m.Called()

	var r0 map[string]*Catalog
	if rf, ok := ret.Get(0).(func() map[string]*Catalog); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*Catalog)
		}
	}

	return r0
}

// IsCatalogReady provides a mock function with given fields:
func (_m *MockRunnerService) IsCatalogReady() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
