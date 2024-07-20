package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", helloHTTP) // must same entrypoint in config
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == os.Getenv("MASTER_NAME") { // setup env config https://cloud.google.com/functions/docs/configuring/env-var#google-cloud-console-ui
		fmt.Fprintf(w, "Hello, %s! You are the master!", html.EscapeString(d.Name))
		return
	}

	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
