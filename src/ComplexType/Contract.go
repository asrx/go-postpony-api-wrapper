package ComplexType

import "github.com/beevik/etree"

type XmlStruct interface {
	ToNode(doc *etree.Element) *etree.Element
}

