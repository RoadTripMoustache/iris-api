package utils

import (
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"os"
)

var (
	loggingError = logging.Error
	osExit       = os.Exit
)

// ProcessError - Process the error given in parameter, log it and shutdown the app.
func ProcessError(err error) {
	loggingError(err, nil)
	osExit(2)
}
