package ComplexType

import (
	"github.com/beevik/etree"
)

type UserCredential struct {
	Key string `xml:"Key"`
	Pwd string `xml:"Pwd"`
}

func (uc *UserCredential) ToNode(node *etree.Element) *etree.Element {

	uNode := node.CreateElement("UserCredential")
	uNode.CreateComment("This is UserCredential Element")

	key := uNode.CreateElement("Key")
	key.CreateText(uc.Key)

	pwd := uNode.CreateElement("Pwd")
	pwd.CreateText(uc.Pwd)

	return node
}