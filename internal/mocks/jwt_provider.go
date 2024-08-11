// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	utils "golang-e-wallet-rest-api/internal/pkgs/utils"

	mock "github.com/stretchr/testify/mock"
)

// JWTProvider is an autogenerated mock type for the JWTProvider type
type JWTProvider struct {
	mock.Mock
}

// CreateToken provides a mock function with given fields: userID
func (_m *JWTProvider) CreateToken(userID int64) (string, error) {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: token
func (_m *JWTProvider) VerifyToken(token string) (utils.JWTClaims, error) {
	ret := _m.Called(token)

	var r0 utils.JWTClaims
	if rf, ok := ret.Get(0).(func(string) utils.JWTClaims); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(utils.JWTClaims)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJWTProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewJWTProvider creates a new instance of JWTProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJWTProvider(t mockConstructorTestingTNewJWTProvider) *JWTProvider {
	mock := &JWTProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
