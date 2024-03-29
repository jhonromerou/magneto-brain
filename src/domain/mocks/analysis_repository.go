// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jhonromerou/magneto-brain/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// AnalysisRepository is an autogenerated mock type for the AnalysisRepository type
type AnalysisRepository struct {
	mock.Mock
}

// GetAnalysisResult provides a mock function with given fields: dnaType, dnaSequence
func (_m *AnalysisRepository) GetAnalysisResult(dnaType string, dnaSequence string) (domain.AnalysisModel, error) {
	ret := _m.Called(dnaType, dnaSequence)

	var r0 domain.AnalysisModel
	if rf, ok := ret.Get(0).(func(string, string) domain.AnalysisModel); ok {
		r0 = rf(dnaType, dnaSequence)
	} else {
		r0 = ret.Get(0).(domain.AnalysisModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dnaType, dnaSequence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ites
func (_m *AnalysisRepository) Register(ites domain.AnalysisModel) error {
	ret := _m.Called(ites)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.AnalysisModel) error); ok {
		r0 = rf(ites)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
