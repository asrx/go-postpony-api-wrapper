package Rate

import (
	xml2 "encoding/xml"
	"fmt"
	"github.com/asrx/go-postpony-api-wrapper/src"
	"github.com/asrx/go-postpony-api-wrapper/src/ComplexType"
	"github.com/asrx/go-postpony-api-wrapper/src/Config"
	"github.com/beevik/etree"
	"strconv"
)

type ShippingRatesRequest struct {
	UserCredential *ComplexType.UserCredential `xml:"UserCredential"`
	PackageInfos []*ComplexType.Package `xml:"PackageInfos"`
	OriginalAddress *ComplexType.Address `xml:"OriginalAddress"`
	DestinationAddress *ComplexType.Address `xml:"DestinationAddress"`
	CustomValue float64 `xml:"CustomValue,omitempty"`
	UpsRate bool `xml:"UpsRate,omitempty"`
	CanGetBaseRate bool `xml:"CanGetBaseRate,omitempty"`
}

func (rate ShippingRatesRequest)ToNode(doc *etree.Document) (string, error) {
	if doc == nil {
		doc = etree.NewDocument()
		doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	}

	node := doc.CreateElement("ShippingRatesRequest")
	node.CreateComment("This is Rate Element")

	// User
	rate.UserCredential.ToNode(node)

	// Packages
	pkginfo := node.CreateElement("PackageInfos")
	for _,pkg := range rate.PackageInfos{
		pkg.NodeName = "PackageItemInfo"
		pkg.ToNode(pkginfo)
	}

	// Shipper
	rate.OriginalAddress.NodeName = "OriginalAddress"
	rate.OriginalAddress.ToNode(node)

	// Recipient
	rate.DestinationAddress.NodeName = "DestinationAddress"
	rate.DestinationAddress.ToNode(node)

	if rate.CustomValue != 0 {
		customValue := node.CreateElement("CustomValue")
		customValue.CreateText(src.Float642String(rate.CustomValue))
	}

	if rate.UpsRate {
		upsRate := node.CreateElement("UpsRate")
		upsRate.CreateText(strconv.FormatBool(rate.UpsRate))
	}

	if rate.CanGetBaseRate {
		canGetBaseRate := node.CreateElement("CanGetBaseRate")
		canGetBaseRate.CreateText(strconv.FormatBool(rate.CanGetBaseRate))
	}

	return doc.WriteToString()
}

func (rate ShippingRatesRequest)GetRate() (rateReply *RateResponse, err error) {
	requestXml, _ := rate.ToNode(nil)
	data, err := src.PostRequest(Config.API_RATE, requestXml)
	if err != nil {
		fmt.Println(err)
		return
	}
	//data := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><RateResponse><Sucess>true</Sucess><Fedex><Data><RateResultDetail><ShipType>FedExPriorityOvernight</ShipType><ShippingBaseRate>844.27</ShippingBaseRate><Price>274.76</Price></RateResultDetail><RateResultDetail><ShipType>FedExStandardOvernight</ShipType><ShippingBaseRate>818.55</ShippingBaseRate><Price>266.26</Price></RateResultDetail><RateResultDetail><ShipType>FedEx2DayAM</ShipType><ShippingBaseRate>700.64</ShippingBaseRate><Price>266.06</Price></RateResultDetail><RateResultDetail><ShipType>FedEx2Day</ShipType><ShippingBaseRate>615.88</ShippingBaseRate><Price>233.95</Price></RateResultDetail><RateResultDetail><ShipType>FedExExpressSaver</ShipType><ShippingBaseRate>471.28</ShippingBaseRate><Price>178.99</Price></RateResultDetail><RateResultDetail><ShipType>FedExGround</ShipType><ShippingBaseRate>111.14</ShippingBaseRate><Price>50.12</Price></RateResultDetail></Data><Msg></Msg><Sucess>true</Sucess></Fedex><Usps><Msg>Usps single package only!</Msg><Sucess>false</Sucess></Usps><Ups><Data><RateResultDetail><ShipType>Ups2NdDayAir</ShipType><ShippingBaseRate>767.85</ShippingBaseRate><Price>253.61</Price></RateResultDetail><RateResultDetail><ShipType>Ups3DaySelect</ShipType><ShippingBaseRate>536.47</ShippingBaseRate><Price>160.31</Price></RateResultDetail><RateResultDetail><ShipType>UpsGround</ShipType><ShippingBaseRate>152.98</ShippingBaseRate><Price>56.30</Price></RateResultDetail><RateResultDetail><ShipType>UpsNextDayAir</ShipType><ShippingBaseRate>1102.07</ShippingBaseRate><Price>266.72</Price></RateResultDetail><RateResultDetail><ShipType>UpsNextDayAirSaver</ShipType><ShippingBaseRate>1075.80</ShippingBaseRate><Price>261.14</Price></RateResultDetail></Data><Sucess>true</Sucess></Ups><Dhl><Msg>未配置站点信息</Msg><Sucess>false</Sucess></Dhl><Pitneybowes><Msg>未配置站点信息</Msg><Sucess>false</Sucess></Pitneybowes><Presort><Msg>打印失败，此线路已暂停服务</Msg><Sucess>false</Sucess></Presort><UspsMerchant><Msg>未配置站点信息</Msg><Sucess>false</Sucess></UspsMerchant></RateResponse>`)
	rateReply = new(RateResponse)
	if err = xml2.Unmarshal(data, rateReply); err != nil {
		fmt.Println("XML Unmarshal Error:",err)
		return
	}
	return
}