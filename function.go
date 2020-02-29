package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"

	"github.com/atb-as/bugsnag-webhook-stackdriver/bugsnag"
)

var (
	projectID string
	logName   string
	logClient *logging.Client
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

	ctx := context.Background()

	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	logClient = client
}

// BugsnagWebhook is a HTTP webhook for receiving errors from bugsnag
func BugsnagWebhook(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
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

	r := Response{
		Success: true,
	}
	jsonResponse(w, r.String(), http.StatusOK)
}
