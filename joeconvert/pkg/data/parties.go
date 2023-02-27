package data

import "github.com/mczaplinski/joe-client/utils"

// TODO: get buyer and delivery parties from a database or somewhere else
var (
	BuyerParty = utils.Party{
		PartyID:   "7023320",
		PartyRole: "buyer",
		Address: utils.Address{
			Name:         "Jacob Elektronik GmbH",
			Name2:        "",
			Street:       "An der Rossweid 5",
			Zip:          "76229",
			City:         "Karlsruhe",
			Country:      "Deutschland",
			CountryCoded: "DE",
			Email:        "info@jacob.de",
		},
	}
	DeliveryParty = utils.Party{
		PartyID:   "7023320",
		PartyRole: "delivery",
		Address: utils.Address{
			Name:         "Jacob Elektronik GmbH",
			Name2:        "z. H. Maximilan Glaeser",
			Street:       "Greschbachstraße 2",
			Zip:          "76229",
			City:         "Karlsruhe",
			Country:      "Deutschland",
			CountryCoded: "DE",
			Email:        "info@jacob.de",
		},
	}
)

// GetSupplierByID returns a supplier party by its ID
func GetSupplierByID(supplierID string) utils.Party {
	// TODO: get supplier party from a database or somewhere else
	if supplierID == "100001" {
		return utils.Party{
			PartyID:   "100001",
			PartyRole: "supplier",
			Address: utils.Address{
				Name:         "Jacob Elektronik GmbH",
				Name2:        "",
				Street:       "Greschbachstraße 2",
				Zip:          "76229",
				City:         "Karlsruhe",
				Country:      "Deutschland",
				CountryCoded: "DE",
				Email:        "info@jacob.de",
			},
		}
	}

	// Fallback for unknown supplier
	return utils.Party{
		PartyID:   supplierID,
		PartyRole: "supplier",
		Address: utils.Address{
			Name: "Unknown supplier",
		},
	}
}
