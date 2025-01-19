package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var protocol, method, domain, resource, postParams string
	pflag.StringVarP(&protocol, "protocol", "p", "", "Request method")
	pflag.StringVarP(&method, "method", "m", "", "Request method")
	pflag.StringVarP(&domain, "domain", "d", "", "Request domain")
	pflag.StringVarP(&resource, "resource", "r", "", "Request resource without first slash")
	pflag.StringVarP(&postParams, "postParams", "", "", "Request post parameters")

	pflag.Parse()

	if protocol == "" {
		fmt.Println("No Protocol specified")

		return
	}

	if method == "" {
		fmt.Println("No Method specified")

		return
	}

	if domain == "" {
		fmt.Println("No Domain specified")

		return
	}

	switch strings.ToUpper(method) {
	case http.MethodGet:
		err := getRequest(protocol, domain, resource)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case http.MethodPost:
		err := postRequest(protocol, domain, resource, postParams)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	default:
		fmt.Printf("Unsupported method: %s\n", method)
	}
}

func getRequest(protocol, domain, resource string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s://%s/%s", protocol, domain, resource),
		nil,
	)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Printf("Response body:\n%s\n", string(body))

	return nil
}

func postRequest(protocol, domain, resource, postParams string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s://%s/%s", protocol, domain, resource),
		strings.NewReader(postParams),
	)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Printf("Response body:\n%s\n", string(body))

	return nil
}
