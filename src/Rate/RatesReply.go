package Rate

type RateResponse struct {
	Success string 			`xml:"Sucess"`

	Fedex 	*RateResult 	`xml:"Fedex"`

	Usps 	*RateResult		`xml:"Usps"`

	Ups		*RateResult		`xml:"Ups"`
}

type RateResult struct {
	Data 	RateResultDetail	`xml:"Data"`
	Msg		string 				`xml:"Msg"`
	Success string 				`xml:"Success"`

}

type RateResultDetail struct {
	Result []*Detail `xml:"RateResultDetail"`
}

type Detail struct {
	ShipType			string		`xml:"ShipType"`
	ShippingBaseRate	string		`xml:"ShippingBaseRate"`
	Price				string		`xml:"Price"`
}