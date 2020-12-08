package Ship

import "github.com/asrx/go-postpony-api-wrapper/src/ComplexType"

type ShipResponse struct {
	LableData *LableData `xml:"LableData"`
	TrackNoList *TrackNoList `xml:"TrackNoList"`
	Msg			string `xml:"Msg"`
	Sucess		string `xml:"Sucess"`
	MainTrackingNum string `xml:"MainTrackingNum"`
	LabelId string `xml:"LabelId"`

	Url string `xml:"Url"`
	TotalFreight string `xml:"TotalFreight"`
	Code string `xml:"Code"`
	PresortNo string `xml:"PresortNo"`
	ResidentialAddress bool `xml:"ResidentialAddress"`
}

type LableData struct {
	Base64Binary string `xml:"base64Binary"`
}

type TrackNoList struct {
	PackageItemInfo []*ComplexType.Package `xml:"PackageItemInfo"`
}