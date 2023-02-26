package convert

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"time"

	"github.com/mczaplinski/joe-client/joeconvert/pkg/data"
	"github.com/mczaplinski/joe-client/utils"
)

// Convert converts the json data to OpenTrans2.1 format end-to-end
func Convert(content []byte) ([]byte, error) {
	// is valid json?
	if !json.Valid(content) {
		return nil, fmt.Errorf("invalid json")
	}

	// parse and validate input json
	orderData, err := data.Parse(content)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error parsing input json"))
	}
	err = data.Validate(orderData)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error validating input json"))
	}

	fmt.Println(orderData)

	// get involved parties
	var (
		supplier = data.GetSupplierByID(fmt.Sprintf("%d", orderData.LieferantNr))
		buyer    = data.BuyerParty
		delivery = data.DeliveryParty
	)

	// get items
	items, totalItems, totalAmount := convertItemsData(orderData, supplier, buyer)

	// convert orderData to OpenTrans2.1 format
	// TODO
	order := &utils.Order{
		Header: utils.OrderHeader{
			OrderInfo: utils.OrderInfo{
				OrderID:   orderData.BestellNummer,
				OrderDate: orderData.ODate.Format("2006-01-02T15:04:05"), // Assumption: variation of time.RFC3339 (as in the example of the JOE API)
				Language:  "ger",                                         // Assumption
				Parties: utils.Parties{
					Parties: []utils.Party{
						supplier,
						buyer,
						delivery,
					},
				},
				OrderPartiesReference: utils.OrderPartiesReference{
					BuyerIDRef:    data.BuyerParty.PartyID,
					SupplierIDRef: supplier.PartyID,
				},
				Currency: "EUR", // Assumption
			},
			ControlInfo: utils.ControlInfo{
				StopAutomaticProcessing: fmt.Sprintf("%d", orderData.Testbestellung),
				GeneratorInfo:           "data bridged using joeconvert",
				GenerationDate:          time.Now(),
			},
		},
		Items: utils.OrderItemList{
			OrderItem: items,
		},
		Summary: utils.OrderSummary{
			TotalItemsNum: fmt.Sprintf("%d", totalItems),
			TotalAmount:   fmt.Sprintf("%f", totalAmount),
		},
	}

	// back to bytes
	result, err := xml.MarshalIndent(order, "", "  ")
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error marshaling output xml"))
	}

	return []byte(xml.Header + string(result)), nil
}

func convertItemsData(input data.Data, supplier, buyer utils.Party) ([]utils.OrderItem, int, float64) {
	var (
		items       []utils.OrderItem = make([]utils.OrderItem, len(input.Artikel))
		totalItems  int               = 0
		totalAmount float64           = 0.0
	)

	for i, item := range input.Artikel {
		priceLineAmount := float64(item.BestellMenge) * item.Preis
		items = append(items, utils.OrderItem{
			LineItemID: fmt.Sprintf("%d", i+1),
			ProductID: utils.ProductID{
				SupplierPID:      item.ArtikelNummer, // Assumption, same as BuyerPID
				BuyerPID:         item.ArtikelNummer, // Assumption, same as SupplierPID
				DescriptionShort: item.ArtikelBeschreibung,
			},
			Quantity:  fmt.Sprintf("%d", item.BestellMenge),
			OrderUnit: "C62", // Assumption: UN/ECE 20 format: C62 = piece
			ProductPriceFix: utils.ProductPriceFix{
				PriceAmount: fmt.Sprintf("%f", item.Preis),
			},
			PriceLineAmount: fmt.Sprintf("%f", priceLineAmount),
		})
		totalAmount += priceLineAmount
	}

	totalItems = len(input.Artikel)

	return items, totalItems, totalAmount
}
