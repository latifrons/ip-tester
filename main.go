package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// 1. Remote IP
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		remoteIP = r.RemoteAddr
	}
	response["remote_ip"] = remoteIP

	// 2. Client IP
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = remoteIP
	}
	response["client_ip"] = clientIP

	// 3. All headers
	headers := make(map[string][]string)
	for name, values := range r.Header {
		headers[name] = values
	}
	response["headers"] = headers

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
