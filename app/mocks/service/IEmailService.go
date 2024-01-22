// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	decimal "github.com/shopspring/decimal"
	mock "github.com/stretchr/testify/mock"

	service "github.com/FrancoLiberali/stori_challenge/app/service"
)

// IEmailService is an autogenerated mock type for the IEmailService type
type IEmailService struct {
	mock.Mock
}

// Send provides a mock function with given fields: destinationEmail, totalBalance, transactionsPerMonth, avgDebit, avgCredit
func (_m *IEmailService) Send(destinationEmail string, totalBalance decimal.Decimal, transactionsPerMonth []service.TransactionsPerMonth, avgDebit decimal.Decimal, avgCredit decimal.Decimal) error {
	ret := _m.Called(destinationEmail, totalBalance, transactionsPerMonth, avgDebit, avgCredit)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, decimal.Decimal, []service.TransactionsPerMonth, decimal.Decimal, decimal.Decimal) error); ok {
		r0 = rf(destinationEmail, totalBalance, transactionsPerMonth, avgDebit, avgCredit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIEmailService creates a new instance of IEmailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIEmailService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IEmailService {
	mock := &IEmailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
