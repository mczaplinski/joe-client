package convert

import (
	"encoding/xml"
	"testing"

	_ "embed"

	"github.com/mczaplinski/joe-client/utils"
)

//go:embed test/convert_input.json
var exampleOrderJSON []byte

//go:embed test/convert_output.xml
var exampleOrderXML []byte

func TestConvert(t *testing.T) {
	// convert content
	result, err := Convert(exampleOrderJSON)
	if err != nil {
		t.Fatal(err)
	}

	// parse results
	resultData := &utils.Order{}
	err = xml.Unmarshal(result, resultData)
	if err != nil {
		t.Fatal(err)
	}

	// parse reference
	exampleData := &utils.Order{}
	xml.Unmarshal(exampleOrderXML, exampleData)

	// compare results
	if resultData.Header.OrderInfo.OrderID != exampleData.Header.OrderInfo.OrderID {
		t.Fatalf("invalid OrderID:\n%s\nWant:\n%s", resultData.Header.OrderInfo.OrderID, exampleData.Header.OrderInfo.OrderID)
	}

	// ... and so on
}
