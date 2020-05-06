package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {

	// Get the Cloud Run $PORT variable from the environment
	listenPort := os.Getenv("PORT")

	// Define the request handler for the default / route
	http.HandleFunc("/", httpDefaultHandler)
	http.ListenAndServe(":"+listenPort, nil)

	//

}

type logEntry struct {
	ResponseCode int    `json:"responseCode"`
	RequestURL   string `json:"requestURL"`
	RequestIP    string `json:"requestIP"`
}

type httpResponse struct {
	HTTPCode         int    `json:"httpCode"`
	CloudRunService  string `json:"cloudRunService"`
	CloudRunRevision string `json:"cloudRunRevision"`
}

func jsonLogRequest(responseCode int, requestURL string, requestIP string) {

	// Define a new JSON Encoder
	var jsonEncoder = json.NewEncoder(os.Stdout)

	// Craft a log entry
	logEntry := logEntry{ResponseCode: responseCode, RequestURL: requestURL, RequestIP: requestIP}

	// Return the Response
	if err := jsonEncoder.Encode(&logEntry); err != nil {
		panic(err)
	}

}

func httpDefaultHandler(w http.ResponseWriter, r *http.Request) {

	// Set the Response Type to Application JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the Version of the Service
	cloudRunService := os.Getenv("K_SERVICE")

	// Get the Version of the Revision
	cloudRunRevision := os.Getenv("K_REVISION")

	// Define a new JSON Encoder
	var jsonEncoder = json.NewEncoder(w)

	// Define a response
	response := httpResponse{HTTPCode: http.StatusOK, CloudRunService: cloudRunService, CloudRunRevision: cloudRunRevision}

	// Log the Response
	jsonLogRequest(response.HTTPCode, r.URL.Path, r.RemoteAddr)

	// Return the Response
	w.WriteHeader(response.HTTPCode)
	if err := jsonEncoder.Encode(&response); err != nil {
		panic(err)
	}
	return

}
