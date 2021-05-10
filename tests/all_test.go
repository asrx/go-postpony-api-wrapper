package tests

import (
	"fmt"
	"github.com/asrx/go-postpony-api-wrapper/src/Cancel"
	"github.com/asrx/go-postpony-api-wrapper/src/ComplexType"
	"github.com/asrx/go-postpony-api-wrapper/src/Rate"
	"github.com/asrx/go-postpony-api-wrapper/src/Ship"
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
		CompanyName:         "AMERICAN NEW LOGISTICS",
		PhoneNumber:         "6262258083",
		StreetLines:         []string{"2440 S. Milliken Avenue"},
		City:                "Ontario",
		StateOrProvinceCode: "CA",
		PostalCode:          "91761",
		CountryCode:         "US",
		//CountryName:		 "United States of America",
		IsResidentialAddress: false,
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
		IsResidentialAddress: false,
	}

	pkgs = []*ComplexType.Package{}
	pkgs = append(pkgs, &ComplexType.Package{
		Length:         10,
		Width:          11,
		Height:         12,
		Weight:			20,
	})

	//pkgs = append(pkgs, &ComplexType.Package{
	//	Length:         20,
	//	Width:          21,
	//	Height:         22,
	//	Weight:			25,
	//})

	//pkgs = append(pkgs, &ComplexType.Package{
	//	Length:         15,
	//	Width:          16,
	//	Height:         17,
	//	Weight:			10,
	//})
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

	rateReply,err := rateShipment.ProcessRate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rateReply.Fedex.Data.Result[0].Price)
	for _,p := range rateReply.Fedex.Data.Result {
		fmt.Printf("ShipType:%s\t Price:%s\n",p.ShipType, p.Price)
	}
	//src.GetRate()
}

func Test_ship(t *testing.T) {
	shipShipment := &Ship.ShipRequest{
		UserCredential: uc,
		RequstInfo:     &Ship.RequstInfo{
			Shipper:      shipper,
			Recipient:    recipient,
			Package:      &Ship.ShipPackage{
				FTRCode:              "30.37 (a)",
				CustomerReference:    "xxxx11122233yyy",
				ContentsType:         "Gift",
				ElectronicExportType: "NoEEISED",
			},
			PackageItems: pkgs,
			LbSize:       "S4X6",
			// 签名类型
			Signature: Ship.SignatureType_FedEx_Direct,
		},
		ShipType:       "FedExGround",// FedExHomeDelivery
		AuthorizedKey:  _Token,
	}

	shipReply, err := shipShipment.ProcessShip()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", shipReply)
	fmt.Println("TrackNoList")
	for i, p := range shipReply.TrackNoList.PackageItemInfo{
		fmt.Println("Pkgs:",i+1)
		fmt.Printf("\tPackageId: %v\n",p.PackageId)
		fmt.Printf("\tLength: %v\n",p.Length)
		fmt.Printf("\tWidth: %v\n",p.Width)
		fmt.Printf("\tHeight: %v\n",p.Height)
		fmt.Printf("\tWeight: %v\n",p.Weight)
		fmt.Printf("\tInsurance: %v\n",p.Insurance)
		fmt.Printf("\tIsOurInsurance: %v\n",p.IsOurInsurance)
		fmt.Printf("\tUspsMailpiece: %v\n",p.UspsMailpiece)
		fmt.Printf("\tTrackNo: %v\n",p.TrackNo)
		fmt.Println("-----------------------")
	}
	fmt.Printf("Msg: %s\n",shipReply.Msg)
	fmt.Printf("Success: %s\n",shipReply.Sucess)
	fmt.Printf("MainTrackingNum: %s\n",shipReply.MainTrackingNum)
	fmt.Printf("LabelId: %s\n",shipReply.LabelId)
	fmt.Printf("Url: %s\n",shipReply.Url)
	fmt.Printf("totalFreight: %s\n",shipReply.TotalFreight)
	fmt.Printf("Code: %s\n",shipReply.Code)
	fmt.Printf("PresortNo: %s\n",shipReply.PresortNo)
	fmt.Printf("ResidentialAddress: %v\n",shipReply.ResidentialAddress)
}

func Test_cancel(t *testing.T) {
	cancelRequest := &Cancel.CancelShipRequst{
		UserCredential: uc,
		LabelId:        "18665423",
	}

	cancelReply,err := cancelRequest.ProcessCancel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cancel Result: %v\n", cancelReply.Success)
}