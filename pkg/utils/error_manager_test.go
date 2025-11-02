package utils

import (
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/mocks/externallibs"
	"testing"
)

func Test_ProcessError(t *testing.T) {
	err := fmt.Errorf("this is the error message")

	var emptyMap map[string]interface{}
	mockLog := new(externallibs.MockLog)
	mockLog.On("Error", err, emptyMap)
	loggingError = mockLog.Error

	mockOs := new(externallibs.MockOS)
	mockOs.On("Exit", 2)
	osExit = mockOs.Exit

	ProcessError(err)

	mockLog.AssertExpectations(t)
	mockOs.AssertExpectations(t)
}
