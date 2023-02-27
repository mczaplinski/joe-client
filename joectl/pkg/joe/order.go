package joe

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Order places a new order
func Order(body []byte, apiKey string) ([]byte, error) {
	// initialize request
	req, err := initRequest(http.MethodPost, BackendURL, body, apiKey)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to initialize request"))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to send request to JOE"))
	}

	// read response
	status, responseBody, err := readResponse(resp)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to read response from JOE"))
	}

	// handle response
	fmt.Println(status)
	// fmt.Println(string(responseBody)) // empty

	// handle status
	if status != http.StatusCreated {
		return nil, fmt.Errorf("failed to create order: %d", status)
	}

	// get location header
	location := resp.Header.Get("Location")
	fmt.Printf("Location: %s , use 'joectl get %s' to get the order\n", location, strings.Split(location, "/")[5])

	return responseBody, nil
}
