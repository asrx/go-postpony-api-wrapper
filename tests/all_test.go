package tests

import (
	"fmt"
	"github.com/asrx/go-postpony-api-wrapper/src/ComplexType"
	"github.com/asrx/go-postpony-api-wrapper/src/Rate"
	"github.com/beevik/etree"
	"log"
	"testing"
)

const (
	_Key  = ""
	_Pwd = ""
	_Token = ""

)

var uc *ComplexType.UserCredential
var shipper *ComplexType.Address
var recipient *ComplexType.Address

var pkgs []*ComplexType.Package

func init() {
	uc = &ComplexType.UserCredential{
		Key: _Key,
		Pwd: _Pwd,
	}

	shipper = &ComplexType.Address{
		PersonName:          "Donovan",
		CompanyName:         "ANL",
		PhoneNumber:         "6262258083",
		StreetLines:         []string{"16018 Adelante st Suite D"},
		City:                "Irwindale",
		StateOrProvinceCode: "CA",
		PostalCode:          "91702",
		CountryCode:         "US",
		//CountryName:		 "United States of America",
	}

	recipient = &ComplexType.Address{
		PersonName:          "Alex",
		CompanyName:         "ANL",
		PhoneNumber:         "8000000000",
		StreetLines:         []string{"401 Independence Rd"},
		City:                "FLORENCE",
		StateOrProvinceCode: "NJ",
		PostalCode:          "08518",
		CountryCode:         "US",
		//CountryName:		 "United States of America",
	}

	pkgs = []*ComplexType.Package{}
	pkgs = append(pkgs, &ComplexType.Package{
		Length:         10,
		Width:          11,
		Height:         12,
		Weight:			20,
	})

	pkgs = append(pkgs, &ComplexType.Package{
		Length:         20,
		Width:          21,
		Height:         22,
		Weight:			25,
	})

	pkgs = append(pkgs, &ComplexType.Package{
		Length:         15,
		Width:          16,
		Height:         17,
		Weight:			10,
	})
}

func Test_user(t *testing.T)  {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	node := &doc.Element
	uc.ToNode(node)
	xml, err := doc.WriteToString()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(xml)
}

func Test_rate(t *testing.T)  {
	rateShipment := &Rate.ShippingRatesRequest{
		UserCredential:     uc,
		PackageInfos:       pkgs,
		OriginalAddress:    shipper,
		DestinationAddress: recipient,
	}

	rateReply,err := rateShipment.GetRate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rateReply.Fedex.Data.Result[0].Price)
	for _,p := range rateReply.Fedex.Data.Result {
		fmt.Printf("ShipType:%s\t Price:%s\n",p.ShipType, p.Price)
	}
	//src.GetRate()
}