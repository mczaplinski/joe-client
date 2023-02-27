package utils

import (
	"encoding/xml"
	"time"
)

// Data model for OpenTrans 2.1 order
// Note: Some fields are not used in this example and are therefore dropped!

// Order is the root element of the order
type Order struct {
	XMLName xml.Name      `xml:"ORDER"`
	Header  OrderHeader   `xml:"ORDER_HEADER"`
	Items   OrderItemList `xml:"ORDER_ITEM_LIST"`
	Summary OrderSummary  `xml:"ORDER_SUMMARY"`
}

// OrderHeader contains the header information of the order
type OrderHeader struct {
	XMLName     xml.Name    `xml:"ORDER_HEADER"`
	OrderInfo   OrderInfo   `xml:"ORDER_INFO"`
	ControlInfo ControlInfo `xml:"CONTROL_INFO"`
}

// OrderInfo contains the order information
type OrderInfo struct {
	XMLName               xml.Name              `xml:"ORDER_INFO"`
	OrderID               string                `xml:"ORDER_ID"`
	OrderDate             string                `xml:"ORDER_DATE"`
	Language              string                `xml:"LANGUAGE"`
	Parties               Parties               `xml:"PARTIES"`
	OrderPartiesReference OrderPartiesReference `xml:"ORDER_PARTIES_REFERENCE"`
	Currency              string                `xml:"CURRENCY"`
}

// ControlInfo contains the order parties reference
type ControlInfo struct {
	XMLName                 xml.Name  `xml:"CONTROL_INFO"`
	StopAutomaticProcessing string    `xml:"STOP_AUTOMATIC_PROCESSING"` // 1 = test order, empty or else: normal order
	GeneratorInfo           string    `xml:"GENERATOR_INFO"`
	GenerationDate          time.Time `xml:"GENERATION_DATE"`
}

// Parties contains the parties of the order
type Parties struct {
	XMLName xml.Name `xml:"PARTIES"`
	Parties []Party  `xml:"PARTY"`
}

// Party contains the party information
type Party struct {
	XMLName   xml.Name `xml:"PARTY"`
	PartyID   string   `xml:"PARTY_ID"`
	PartyRole string   `xml:"PARTY_ROLE"`
	Address   Address  `xml:"ADDRESS"`
}

// Address contains the address information
type Address struct {
	XMLName      xml.Name `xml:"ADDRESS"`
	Name         string   `xml:"NAME"`
	Name2        string   `xml:"NAME2,omitempty"`
	Street       string   `xml:"STREET,omitempty"`
	Zip          string   `xml:"ZIP,omitempty"`
	City         string   `xml:"CITY,omitempty"`
	Country      string   `xml:"COUNTRY,omitempty"`
	CountryCoded string   `xml:"COUNTRY_CODED,omitempty"`
	Email        string   `xml:"EMAIL,omitempty"`
}

// OrderPartiesReference contains the order parties reference
type OrderPartiesReference struct {
	XMLName       xml.Name `xml:"ORDER_PARTIES_REFERENCE"`
	BuyerIDRef    string   `xml:"BUYER_IDREF"`
	SupplierIDRef string   `xml:"SUPPLIER_IDREF"`
}

// OrderItemList contains the order item list
type OrderItemList struct {
	XMLName   xml.Name    `xml:"ORDER_ITEM_LIST"`
	OrderItem []OrderItem `xml:"ORDER_ITEM"`
}

// OrderItem contains the order item information
type OrderItem struct {
	XMLName         xml.Name        `xml:"ORDER_ITEM"`
	LineItemID      string          `xml:"LINE_ITEM_ID"`
	ProductID       ProductID       `xml:"PRODUCT_ID"`
	Quantity        string          `xml:"QUANTITY"`
	OrderUnit       string          `xml:"ORDER_UNIT"`
	ProductPriceFix ProductPriceFix `xml:"PRODUCT_PRICE_FIX"`
	PriceLineAmount string          `xml:"PRICE_LINE_AMOUNT"`
}

// ProductID contains the product ID information
type ProductID struct {
	XMLName          xml.Name `xml:"PRODUCT_ID"`
	SupplierPID      string   `xml:"SUPPLIER_PID"`
	BuyerPID         string   `xml:"BUYER_PID"`
	DescriptionShort string   `xml:"DESCRIPTION_SHORT"`
}

// ProductPriceFix contains the product price fix information
type ProductPriceFix struct {
	XMLName     xml.Name `xml:"PRODUCT_PRICE_FIX"`
	PriceAmount string   `xml:"PRICE_AMOUNT"`
}

// OrderSummary contains the order summary information
type OrderSummary struct {
	XMLName       xml.Name `xml:"ORDER_SUMMARY"`
	TotalItemsNum string   `xml:"TOTAL_ITEM_NUM"`
	TotalAmount   string   `xml:"TOTAL_AMOUNT"`
}
