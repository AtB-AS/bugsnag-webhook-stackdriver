package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/pubsub"

	"github.com/atb-as/bugsnag-webhook-stackdriver/bugsnag"
)

var (
	ctx       context.Context
	projectID string
	logName   string
	logClient *logging.Client
	topic     *pubsub.Topic
)

// Response is the response structure returned by the webhook
type Response struct {
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

func (r Response) String() string {
	j, _ := json.Marshal(r)
	return string(j)
}

func jsonResponse(w http.ResponseWriter, j string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, j)
}

func init() {
	projectID = os.Getenv("GCP_PROJECT")
	if projectID == "" {
		log.Fatal("GCP_PROJECT not found in environment")
	}

	logName = os.Getenv("LOG_NAME")
	if logName == "" {
		log.Fatal("LOG_NAME not found in environment")
	}

	topicName := os.Getenv("PUBSUB_TOPIC")
	if topicName == "" {
		topicName = logName
	}

	ctx = context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create log client: %v", err)
	}
	logClient = client

	pubSubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create pubsub client: %v", err)
	}

	topic, err = pubSubClient.CreateTopic(ctx, topicName)
	if err != nil {
		topic = pubSubClient.Topic(topicName)
	}
}

// BugsnagWebhook is a HTTP webhook for receiving errors from bugsnag
func BugsnagWebhook(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	defer req.Body.Close()

	ev := bugsnag.Event{}

	err := dec.Decode(&ev)
	if err != nil {
		r := Response{
			Error:   fmt.Errorf("decode: %w", err).Error(),
			Success: false,
		}
		jsonResponse(w, r.String(), http.StatusBadRequest)
		return
	}

	// TODO: request validation?
	logToStackDriver(&ev)
	publish(&ev)

	r := Response{
		Success: true,
	}
	jsonResponse(w, r.String(), http.StatusOK)
}
