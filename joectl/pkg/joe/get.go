package joe

import (
	"errors"
	"fmt"
	"net/http"
)

// Get returns the order with the given orderID
func Get(orderID, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/order", BackendURL, orderID)

	// initialize request
	req, err := initRequest(http.MethodGet, url, nil, apiKey)
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
	// fmt.Println(string(responseBody))

	return responseBody, nil
}
