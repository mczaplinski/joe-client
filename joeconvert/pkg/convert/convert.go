package convert

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/mczaplinski/joe-client/joeconvert/pkg/models"
	"github.com/mczaplinski/joe-client/utils"
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

	// convert data to OpenTrans2.1 format
	// TODO
	order := &utils.Order{
		// TODO mapping
		Header: utils.Header{
			OrderInfo: utils.OrderInfo{
				ID:       "",
				Date:     "",
				Language: "",
				Parties: utils.Parties{
					Parties: []utils.Party{
						{
							ID:   "",
							Role: "",
							Address: utils.Address{
								Name:         "",
								Name2:        "",
								Street:       "",
								Zip:          "",
								City:         "",
								Country:      "",
								CountryCoded: "",
								Email:        "",
							},
						},
						// TODO
					},
				},
				Reference: utils.Reference{
					BuyerID:    "",
					SupplierID: "",
				},
				Currency: "",
			},
			ControlInfo: utils.ControlInfo{
				StopAutomaticProcessing: "",
			},
		},
	}

	// back to bytes
	result, err := xml.Marshal(order)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error marshaling output xml"))
	}

	return result, nil // TODO
}
