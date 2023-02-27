package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mczaplinski/joe-client/joeconvert/pkg/convert"
	"github.com/mczaplinski/joe-client/joectl/pkg/joe"
)

func main() {
	apiKey := os.Getenv("JOE_API_KEY")

	// create a new order
	orderID := "S1401" // TODO: generate a random order ID to avoid conflicts
	data := fmt.Sprintf(`{
		"ODate": "2023-02-17T18:25:43.511Z",
		"BestellNummer": "%s",
		"Testbestellung": 1,
		"LieferantNr": 100001,
		"Artikel": [
			{
			"ArtikelNummer": "1301014",
			"Artikel Name ": "Das Elektronische Hanuta",
			"Artikel Beschreibung": "Auch ein Hanuta schmeckt besser mit verbauten Halbleitern",
			"BestellMenge": 2,
			"Preis": 10.00
			}
		]
	}`, orderID)

	fmt.Printf("Let's check our input JSON\n")
	fmt.Println(data)

	order, err := convert.Convert([]byte(data))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nLet's check our order YAML before sending it\n")
	fmt.Println(string(order))

	// send order to JOE
	_, err = joe.Order(order, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// get order from JOE
	response, err := joe.Get(orderID, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nLet's check our order inside JOE:\n")
	fmt.Print(string(response))
}
