package joe

import (
	"errors"
	"fmt"
	"net/http"
)

// Orders returns all orders
func Orders(apiKey string) ([]byte, error) {
	// initialize request
	req, err := initRequest(http.MethodGet, BackendURL, nil, apiKey)
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

	return responseBody, nil
}
