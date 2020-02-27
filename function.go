package function

import (
	"fmt"
	"net/http"
)

// HelloHTTP is an HTTP Cloud function
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
