package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/atb-as/bugsnag-webhook-stackdriver/bugsnag"
)

var (
	projectID string
)

func init() {
	env, ok := os.LookupEnv("GCP_PROJECT")
	projectID = env

	if !ok {
		log.Fatal("GCP_PROJECT not found in environment")
	}
}

// BugsnagWebhook is a HTTP webhook for receiving errors from bugsnag
func BugsnagWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	trigger := &bugsnag.Trigger{}
	err = json.Unmarshal(body, trigger)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Do something with trigger here

	fmt.Fprintf(w, "OK")
}
