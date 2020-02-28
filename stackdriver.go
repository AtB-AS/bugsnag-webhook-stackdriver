package function

import (
	"cloud.google.com/go/logging"
)

func logToStackDriver(msg []byte) {
	logName := "bugsnag-errors"
	logger := logger.Logger(logName).StandardLogger(logging.Info)

	logger.Println(msg)
}
