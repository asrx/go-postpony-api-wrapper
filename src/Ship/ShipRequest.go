package Ship

import (
	"encoding/xml"
	"fmt"
	"github.com/asrx/go-postpony-api-wrapper/src"
	"github.com/asrx/go-postpony-api-wrapper/src/ComplexType"
	"github.com/asrx/go-postpony-api-wrapper/src/Config"
	"github.com/beevik/etree"
)

type ShipRequest struct {
	UserCredential  *ComplexType.UserCredential `xml:"UserCredential"`
	RequstInfo		*RequstInfo `xml:"RequstInfo"`
	ShipType		string	`xml:"ShipType"`
	AuthorizedKey	string	`xml:"AuthorizedKey"`
}

type RequstInfo struct {
	Shipper		*ComplexType.Address	`xml:"Shipper"`
	Recipient	*ComplexType.Address	`xml:"Recipient"`
	Package		*ShipPackage			`xml:"Package"`
	PackageItems []*ComplexType.Package	`xml:"PackageItems"`
	LbSize		string			`xml:"LbSize"`
}

type ShipPackage struct {
	ShipDate	string `xml:"ShipDate"`
	InvoiceNumber	string `xml:"InvoiceNumber"`
	FTRCode	string `xml:"FTRCode"`
	CustomerReference string `xml:"CustomerReference"`
	ContentsType string `xml:"ContentsType"`
	ElectronicExportType string `xml:"ElectronicExportType"`
}

func (ship ShipRequest)ToNode(doc *etree.Document)(string, error)  {
	if doc == nil {
		doc = etree.NewDocument()
		doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	}
	node := doc.CreateElement("ShipRequest")
	node.CreateComment("This is Ship Element")

	// User
	ship.UserCredential.ToNode(node)

	// RequstInfo
	// Shiper
	requestInfo := node.CreateElement("RequstInfo")
	ship.RequstInfo.Shipper.NodeName = "Shipper"
	ship.RequstInfo.Shipper.ToNode(requestInfo)
	// Recipient
	ship.RequstInfo.Recipient.NodeName = "Recipient"
	ship.RequstInfo.Recipient.ToNode(requestInfo)
	// Package
	packageNode := requestInfo.CreateElement("Package")
	shipData := packageNode.CreateElement("ShipDate")
	shipData.CreateText(ship.RequstInfo.Package.ShipDate)
	//LbSize
	lbSizeNode := requestInfo.CreateElement("LbSize")
	lbSizeNode.CreateText(ship.RequstInfo.LbSize)

	invoiceNumberNode := packageNode.CreateElement("InvoiceNumber")
	invoiceNumberNode.CreateText(ship.RequstInfo.Package.InvoiceNumber)
	fTRCodeNode := packageNode.CreateElement("FTRCode")
	fTRCodeNode.CreateText(ship.RequstInfo.Package.FTRCode)
	customerReferenceNode := packageNode.CreateElement("CustomerReference")
	customerReferenceNode.CreateText(ship.RequstInfo.Package.CustomerReference)
	contentsTypeNode := packageNode.CreateElement("ContentsType")
	contentsTypeNode.CreateText(ship.RequstInfo.Package.ContentsType)
	electronicExportTypeNode := packageNode.CreateElement("ElectronicExportType")
	electronicExportTypeNode.CreateText(ship.RequstInfo.Package.ElectronicExportType)

	// PackageItems
	packageItemsNode := requestInfo.CreateElement("PackageItems")
	for _,pkg := range ship.RequstInfo.PackageItems {
		pkg.NodeName = "PackageItemInfo"
		pkg.ToNode(packageItemsNode)
	}


	shipType := node.CreateElement("ShipType")
	shipType.CreateText(ship.ShipType)

	authorizedKey := node.CreateElement("AuthorizedKey")
	authorizedKey.CreateText(ship.AuthorizedKey)

	return doc.WriteToString()
}

func (ship ShipRequest) ProcessShip()(shipReply *ShipResponse, err error) {
	requestXml, _ := ship.ToNode(nil)
	data,err := src.PostRequest(Config.API_SHIP, requestXml)
	if err != nil {
		fmt.Println("Post Request Ship Error:", err)
		return
	}
	//data := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ShipResponse><LableData><base64Binary>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cCBOViBcKEFHUEwtdmVyc2lvblwpKS9DcmVhdGlvbkRhdGUoRDoyMDIwMTIwODExMjI1MiswOCcwMCcpL01vZERhdGUoRDoyMDIwMTIwODExMjI1MiswOCcwMCcpPj4KZW5kb2JqCnhyZWYKMCAxMAowMDAwMDAwMDAwIDY1NTM1IGYgCjAwMDAwMDAwMTUgMDAwMDAgbiAKMDAwMDAwODA5NyAwMDAwMCBuIAowMDAwMDE2OTQzIDAwMDAwIG4gCjAwMDAwMDgyMTkgMDAwMDAgbiAKMDAwMDAwODMzNiAwMDAwMCBuIAowMDAwMDE2NzA0IDAwMDAwIG4gCjAwMDAwMTY4MjYgMDAwMDAgbiAKMDAwMDAxNzAwMCAwMDAwMCBuIAowMDAwMDE3MDQ1IDAwMDAwIG4gCnRyYWlsZXIKPDwvU2l6ZSAxMC9Sb290IDggMCBSL0luZm8gOSAwIFIvSUQgWzxhNDczYTgxMDlmODVmYjA2MmM3ZGVjNGZhNTE4NmI4Yz48YTQ3M2E4MTA5Zjg1ZmIwNjJjN2RlYzRmYTUxODZiOGM+XT4+CiVpVGV4dC01LjUuMTAKc3RhcnR4cmVmCjE3MjAzCiUlRU9GCg==</base64Binary></LableData><TrackNoList><PackageItemInfo><PackageId>0</PackageId><Length>10.000</Length><Width>11.000</Width><Height>12.000</Height><Weight>20.000</Weight><Insurance>0.000</Insurance><IsOurInsurance>false</IsOurInsurance><UspsMailpiece></UspsMailpiece><TrackNo>781047198665</TrackNo></PackageItemInfo><PackageItemInfo><PackageId>0</PackageId><Length>20.000</Length><Width>21.000</Width><Height>22.000</Height><Weight>25.000</Weight><Insurance>0.000</Insurance><IsOurInsurance>false</IsOurInsurance><UspsMailpiece></UspsMailpiece><TrackNo>781047199856</TrackNo></PackageItemInfo></TrackNoList><Msg></Msg><Sucess>true</Sucess><MainTrackingNum>781047198665</MainTrackingNum><LabelId>17864161</LabelId><Url>http://file.postpony.com/LabelCache/2020/12/08/a1c64395-c2b8-46ba-9225-b979b50aa4fb.pdf</Url><TotalFreight>36.85</TotalFreight><Code></Code><PresortNo></PresortNo><ResidentialAddress>false</ResidentialAddress></ShipResponse>`
	shipReply = new(ShipResponse)
	if err = xml.Unmarshal(data, shipReply); err != nil {
		fmt.Println("XML Unmarshal Error:",err)
		return
	}

	return
}