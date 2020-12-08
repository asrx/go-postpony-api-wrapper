package ComplexType

import (
	"github.com/asrx/go-postpony-api-wrapper/src"
	"github.com/beevik/etree"
	"strconv"
)

type Package struct {
	NodeName string `xml:"PackageItemInfo"`

	PackageId int `xml:"PackageId"`
	UspsMailpiece string `xml:"UspsMailpiece"`
	Length float64 `xml:"Length"`
	Width float64 `xml:"Width"`
	Height float64 `xml:"Height"`
	Weight float64 `xml:"Weight"`
	WeightOz float64 `xml:"WeightOz"`
	Insurance float64 `xml:"Insurance"`
	IsOurInsurance bool `xml:"IsOurInsurance"`
}

func (pkg *Package) ToNode(n *etree.Element) *etree.Element {

	node := n.CreateElement(pkg.NodeName)

	if pkg.PackageId != 0 {
		packageId := node.CreateElement("PackageId")
		packageId.CreateText(strconv.Itoa(pkg.PackageId))
	}

	if pkg.UspsMailpiece != "" {
		uspsMailpiece := node.CreateElement("UspsMailpiece")
		uspsMailpiece.CreateText(pkg.UspsMailpiece)
	}

	length := node.CreateElement("Length")
	length.CreateText(src.Float642String(pkg.Length))

	width := node.CreateElement("Width")
	width.CreateText(src.Float642String(pkg.Width))

	height := node.CreateElement("Height")
	height.CreateText(src.Float642String(pkg.Height))

	weight := node.CreateElement("Weight")
	weight.CreateText(src.Float642String(pkg.Weight))

	if pkg.WeightOz != 0 {
		weightOz := node.CreateElement("WeightOz")
		weightOz.CreateText(src.Float642String(pkg.WeightOz))
	}

	insurance := node.CreateElement("Insurance")
	insurance.CreateText(src.Float642String(pkg.Insurance))

	if pkg.IsOurInsurance {
		isOurInsurance := node.CreateElement("IsOurInsurance")
		isOurInsurance.CreateText(strconv.FormatBool(pkg.IsOurInsurance))
	}

	return n
}