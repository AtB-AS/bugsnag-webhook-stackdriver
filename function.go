package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"

	"github.com/atb-as/bugsnag-webhook-stackdriver/bugsnag"
)

var (
	logClient *logging.Client
)

// Response is the response structure returned by the webhook
type Response struct {
	Status  string `json:"status"`
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
	projectID, ok := os.LookupEnv("GCP_PROJECT")
	if !ok {
		log.Fatal("GCP_PROJECT not found in environment")
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
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r := Response{
			Status:  "failed to read request body",
			Success: false,
		}
		jsonResponse(w, r.String(), http.StatusInternalServerError)
		return
	}

	ev := bugsnag.Event{}
	err = json.Unmarshal(body, &ev)
	if err != nil {
		r := Response{
			Status:  "failed to parse request body",
			Success: false,
		}
		jsonResponse(w, r.String(), http.StatusBadRequest)
		return
	}

	// Do something with Event here
	logToStackDriver(&ev)

	r := Response{
		Status:  "OK",
		Success: true,
	}
	jsonResponse(w, r.String(), http.StatusOK)

}
