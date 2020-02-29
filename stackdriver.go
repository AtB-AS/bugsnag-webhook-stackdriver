package function

import (
	"cloud.google.com/go/logging"
	"cloud.google.com/go/pubsub"
	"github.com/atb-as/bugsnag-webhook-stackdriver/bugsnag"
)

func logToStackDriver(event *bugsnag.Event) {
	logClient.Logger(logName).Log(logging.Entry{
		Severity: logging.Alert,
		Labels: map[string]string{
			"platform": event.Error.Device.OSName,
			"app":      event.Error.App.ID,
		},
		Payload: event,
	})
}

func publish(event *bugsnag.Event) {
	topic.Publish(ctx, &pubsub.Message{
		Attributes: map[string]string{
			"platform": event.Error.Device.OSName,
			"app":      event.Error.App.ID,
		},
	})
}
