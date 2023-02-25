package utils

import "encoding/xml"

// Data model for OpenTrans 2.1 order

type Order struct {
	XMLName xml.Name `xml:"ORDER"`
	Header  Header   `xml:"ORDER_HEADER"`
	Items   Items    `xml:"ORDER_ITEM_LIST"`
	Summary Summary  `xml:"ORDER_SUMMARY"`
}

type Header struct {
	XMLName     xml.Name    `xml:"ORDER_HEADER"`
	Info        OrderInfo   `xml:"ORDER_INFO"`
	ControlInfo ControlInfo `xml:"CONTROL_INFO"`
}

type OrderInfo struct {
	XMLName   xml.Name  `xml:"ORDER_INFO"`
	ID        string    `xml:"ORDER_ID"`
	Date      string    `xml:"ORDER_DATE"`
	Language  string    `xml:"LANGUAGE"`
	Parties   Parties   `xml:"PARTIES"`
	Reference Reference `xml:"ORDER_PARTIES_REFERENCE"`
	Currency  string    `xml:"CURRENCY"`
}

type ControlInfo struct {
	XMLName                 xml.Name `xml:"CONTROL_INFO"`
	StopAutomaticProcessing string   `xml:"STOP_AUTOMATIC_PROCESSING"`
}

type Parties struct {
	XMLName xml.Name `xml:"PARTIES"`
	Parties []Item   `xml:"PARTY"`
}

type Party struct {
	XMLName xml.Name `xml:"PARTY"`
	ID      string   `xml:"PARTY_ID"`
	Role    string   `xml:"PARTY_ROLE"`
	Address Address  `xml:"ADDRESS"`
}

type Address struct {
	XMLName      xml.Name `xml:"ADDRESS"`
	Name         string   `xml:"NAME"`
	Name2        string   `xml:"NAME2"`
	Street       string   `xml:"STREET"`
	Zip          string   `xml:"ZIP"`
	City         string   `xml:"CITY"`
	Country      string   `xml:"COUNTRY"`
	CountryCoded string   `xml:"COUNTRY_CODED"`
	Email        string   `xml:"EMAIL"`
}

type Reference struct {
	XMLName    xml.Name `xml:"ORDER_PARTIES_REFERENCE"`
	BuyerID    string   `xml:"BUYER_IDREF"`
	SupplierID string   `xml:"SUPPLIER_IDREF"`
}

type Items struct {
	XMLName xml.Name `xml:"ORDER_ITEM_LIST"`
	Items   []Item   `xml:"ORDER_ITEM"`
}

type Item struct {
	XMLName    xml.Name `xml:"ORDER_ITEM"`
	ID         string   `xml:"LINE_ITEM_ID"`
	Product    Product  `xml:"PRODUCT_ID"`
	Quantity   string   `xml:"QUANTITY"`
	Unit       string   `xml:"ORDER_UNIT"`
	Price      Price    `xml:"PRODUCT_PRICE_FIX"`
	LineAmount string   `xml:"PRICE_LINE_AMOUNT"`
}

type Product struct {
	XMLName     xml.Name `xml:"PRODUCT_ID"`
	SupplierID  string   `xml:"SUPPLIER_PID"`
	BuyerID     string   `xml:"BUYER_PID"`
	Description string   `xml:"DESCRIPTION_SHORT"`
}

type Price struct {
	XMLName xml.Name `xml:"PRODUCT_PRICE_FIX"`
	Amount  string   `xml:"PRICE_AMOUNT"`
}

type Summary struct {
	XMLName     xml.Name `xml:"ORDER_SUMMARY"`
	TotalItems  string   `xml:"TOTAL_ITEM_NUM"`
	TotalAmount string   `xml:"TOTAL_AMOUNT"`
}
