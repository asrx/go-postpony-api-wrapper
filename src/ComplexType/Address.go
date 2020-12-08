package ComplexType

import (
	"github.com/beevik/etree"
	"strconv"
)

type Address struct {
	NodeName string `xml:"-"`

	PersonName string `xml:"PersonName,omitempty"`
	CompanyName string `xml:"CompanyName,omitempty"`
	PhoneNumber string `xml:"PhoneNumber,omitempty"`
	StreetLines []string `xml:"StreetLines"`
	City string `xml:"City"`
	StateOrProvinceCode string `xml:"StateOrProvinceCode"`
	PostalCode string `xml:"PostalCode"`
	Zip4 string `xml:"Zip4,omitempty"`
	CountryCode string `xml:"CountryCode"`
	CountryName string `xml:"CountryName"`
	IsResidentialAddress bool `xml:"IsResidentialAddress"`
}

func (addr *Address)ToNode(doc *etree.Element) *etree.Element {
	node := doc.CreateElement(addr.NodeName)

	if addr.PersonName != "" {
		personName := node.CreateElement("PersonName")
		personName.CreateText(addr.PersonName)
	}

	if addr.CompanyName != "" {
		personName := node.CreateElement("CompanyName")
		personName.CreateText(addr.CompanyName)
	}

	if addr.PhoneNumber != "" {
		personName := node.CreateElement("PhoneNumber")
		personName.CreateText(addr.PhoneNumber)
	}

	if addr.Zip4 != "" {
		personName := node.CreateElement("Zip4")
		personName.CreateText(addr.Zip4)
	}

	streetLines := node.CreateElement("StreetLines")
	for _, line := range addr.StreetLines{
		lineElem := streetLines.CreateElement("string")
		lineElem.CreateText(line)
	}

	city := node.CreateElement("City")
	city.CreateText(addr.City)

	stateOrProvinceCode := node.CreateElement("StateOrProvinceCode")
	stateOrProvinceCode.CreateText(addr.StateOrProvinceCode)

	postalCode := node.CreateElement("PostalCode")
	postalCode.CreateText(addr.PostalCode)

	countryCode := node.CreateElement("CountryCode")
	countryCode.CreateText(addr.CountryCode)

	countryName := node.CreateElement("CountryName")
	countryName.CreateText(addr.CountryName)

	IsResidentialAddress := node.CreateElement("IsResidentialAddress")
	IsResidentialAddress.CreateText(strconv.FormatBool(addr.IsResidentialAddress))

	return doc
}