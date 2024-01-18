package main

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter

	a(&t, expected, actual, msgAndArgs...)

	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertError is a helper function to allow the step function to call
// assertion functions where you want to compare an error value to a
// predefined state like nil, empty or true/false.
func assertError(a errorAssertion, actual error, msgAndArgs ...interface{}) error {
	var t asserter

	a(&t, actual, msgAndArgs...)

	return t.err
}

type errorAssertion func(t assert.TestingT, actual error, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...) //nolint:goerr113 // necessary to implement assert.TestingT
}
