package convert

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mczaplinski/joe-client/joeconvert/pkg/models"
)

func Convert(content []byte) ([]byte, error) {
	// is valid json?
	if !json.Valid(content) {
		return nil, fmt.Errorf("invalid json")
	}

	// parse and validate input json
	data, err := models.Parse(content)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error parsing input json"))
	}
	err = models.Validate(data)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error validating input json"))
	}

	fmt.Println(data)

	// convert data to joectl format
	// TODO

	// back to bytes
	// TODO

	return content, nil // TODO
}
