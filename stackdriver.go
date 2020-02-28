package function

import (
	"cloud.google.com/go/logging"
	"github.com/atb-as/bugsnag-webhook-stackdriver/bugsnag"
)

func logToStackDriver(event *bugsnag.Event) {
	logName := "bugsnag-errors"
	logger := logClient.Logger(logName)

	logger.Log(logging.Entry{
		Severity: logging.Alert,
		Labels: map[string]string{
			"platform": event.Error.Device.OSName,
			"app":      event.Error.App.ID,
		},
		Payload: event,
	})
}
