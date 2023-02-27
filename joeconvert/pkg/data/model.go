package data

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

// validator instance
var validate *validator.Validate

// Data is the example JSON data type
type Data struct {
	ODate          time.Time `json:"ODate" validate:"required"`
	BestellNummer  string    `json:"BestellNummer" validate:"required"`
	Testbestellung int       `json:"Testbestellung"`
	LieferantNr    int       `json:"LieferantNr" validate:"required"`
	Artikel        []Artikel `json:"Artikel" validate:"required,dive,required"`
}

// Artikel is one of the items in the data
type Artikel struct {
	ArtikelNummer       string  `json:"ArtikelNummer" validate:"required"`
	ArtikelName         string  `json:"Artikel Name " validate:"required"`
	ArtikelBeschreibung string  `json:"Artikel Beschreibung"`
	BestellMenge        int     `json:"BestellMenge" validate:"required"`
	Preis               float64 `json:"Preis" validate:"required"`
}

// Parse JSON data type
func Parse(data []byte) (Data, error) {
	var d Data
	err := json.Unmarshal(data, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Validate JSON data type
func Validate(data Data) error {
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
