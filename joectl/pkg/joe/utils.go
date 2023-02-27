package joe

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func initRequest(method, url string, xmlBody []byte, apiKey string) (*http.Request, error) {
	var req *http.Request
	var err error
	if xmlBody != nil {
		req, err = http.NewRequest(
			method,
			url,
			bytes.NewBuffer(xmlBody))
		req.Header.Set("Content-Type", "application/xml")
	} else {
		req, err = http.NewRequest(
			method,
			url,
			nil)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/xml")

	// appending to existing query args
	q := req.URL.Query()
	q.Add("apikey", apiKey)

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func readResponse(resp *http.Response) (int, []byte, error) {
	defer resp.Body.Close()
	status := resp.StatusCode
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, errors.Join(err, fmt.Errorf("failed to read response from JOE"))
	}

	return status, responseBody, nil
}
