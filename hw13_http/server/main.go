package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var ip, port string
	pflag.StringVarP(&ip, "ip", "", "0.0.0.0", "IP address")
	pflag.StringVarP(&port, "port", "", "8080", "Port")

	pflag.Parse()

	fmt.Println("Run server on", ip, ":", port)

	mux := http.NewServeMux()
	mux.HandleFunc(`/get`, getRequest)
	mux.HandleFunc(`/post`, postRequest)

	server := &http.Server{
		Addr:         ip + ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

func getRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Method not allowed"))

		return
	}

	body := "New request:\r\n"
	body += fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	for k, v := range req.URL.Query() {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}

	res.Write([]byte(body))

	fmt.Println(body)
}

func postRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Method not allowed"))

		return
	}

	body := "New request:\r\n"
	body += fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	var data map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil && !errors.Is(err, io.EOF) {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error decoding Body"))

		return
	}
	for k, v := range data {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))

	fmt.Println(body)
}
