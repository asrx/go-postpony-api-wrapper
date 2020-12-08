package Cancel

import (
	"encoding/xml"
	"fmt"
	"github.com/asrx/go-postpony-api-wrapper/src"
	"github.com/asrx/go-postpony-api-wrapper/src/ComplexType"
	"github.com/asrx/go-postpony-api-wrapper/src/Config"
	"github.com/beevik/etree"
)

type CancelShipRequst struct {
	UserCredential *ComplexType.UserCredential `xml:"UserCredential"`
	LabelId		string `xml:"LabelId"`
}

func (c *CancelShipRequst) ToNode(doc *etree.Document)(string, error)  {
	if doc == nil {
		doc = etree.NewDocument()
		doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	}
	node := doc.CreateElement("CancelShipRequst")
	node.CreateComment("This is Cancel Element")

	// User
	c.UserCredential.ToNode(node)

	labelIdNode := node.CreateElement("LabelId")
	labelIdNode.CreateText(c.LabelId)

	return doc.WriteToString()
}

func (c *CancelShipRequst) ProcessCancel() (cancelReply *CancelShipResponse, err error)  {
	requestXml, _ := c.ToNode(nil)
	data, err := src.PostRequest(Config.API_CANCEL, requestXml)
	if err != nil {
		fmt.Println("Post Request Ship Error:",err)
		return
	}

	//data := []byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><CancelShipResponse><Success>true</Success></CancelShipResponse>`)
	cancelReply = new(CancelShipResponse)
	if err = xml.Unmarshal(data, cancelReply); err != nil {
		fmt.Println("XML Unmarshal Error:",err)
		return
	}
	return
}
