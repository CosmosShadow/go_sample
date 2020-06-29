package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct{
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestSomething(t *testing.T) {
	testObj := new(MyMockedObject)

	// setup
	testObj.On("DoSomething", 123).Return(true, nil)
	// testObj.On("DoSomething", mock.Anything).Return(true, nil)

	value, err := testObj.DoSomething(123)
	assert.True(t, value)
	require.NoError(t, err)

	// assert that the expectations were met
	// testObj.AssertExpectations(t)
}









