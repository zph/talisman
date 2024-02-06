// Code generated by MockGen. DO NOT EDIT.
// Source: checksumcalculator/checksumcalculator.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChecksumCalculator is a mock of ChecksumCalculator interface.
type MockChecksumCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockChecksumCalculatorMockRecorder
}

// MockChecksumCalculatorMockRecorder is the mock recorder for MockChecksumCalculator.
type MockChecksumCalculatorMockRecorder struct {
	mock *MockChecksumCalculator
}

// NewMockChecksumCalculator creates a new mock instance.
func NewMockChecksumCalculator(ctrl *gomock.Controller) *MockChecksumCalculator {
	mock := &MockChecksumCalculator{ctrl: ctrl}
	mock.recorder = &MockChecksumCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecksumCalculator) EXPECT() *MockChecksumCalculatorMockRecorder {
	return m.recorder
}

// CalculateCollectiveChecksumForPattern mocks base method.
func (m *MockChecksumCalculator) CalculateCollectiveChecksumForPattern(fileNamePattern string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateCollectiveChecksumForPattern", fileNamePattern)
	ret0, _ := ret[0].(string)
	return ret0
}

// CalculateCollectiveChecksumForPattern indicates an expected call of CalculateCollectiveChecksumForPattern.
func (mr *MockChecksumCalculatorMockRecorder) CalculateCollectiveChecksumForPattern(fileNamePattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateCollectiveChecksumForPattern", reflect.TypeOf((*MockChecksumCalculator)(nil).CalculateCollectiveChecksumForPattern), fileNamePattern)
}

// SuggestTalismanRC mocks base method.
func (m *MockChecksumCalculator) SuggestTalismanRC(fileNamePatterns []string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestTalismanRC", fileNamePatterns)
	ret0, _ := ret[0].(string)
	return ret0
}

// SuggestTalismanRC indicates an expected call of SuggestTalismanRC.
func (mr *MockChecksumCalculatorMockRecorder) SuggestTalismanRC(fileNamePatterns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestTalismanRC", reflect.TypeOf((*MockChecksumCalculator)(nil).SuggestTalismanRC), fileNamePatterns)
}