package XML

import (
	"encoding/xml"
)

type ExampleProduct struct {
	XMLName            xml.Name     `xml:"ExampleProduct"`
	InternalName       string       `xml:"internalName,attr"`
	Website            string       `xml:"website,attr"`
	ProductName        Names        `xml:"Names"`
	ProductDescription Descriptions `xml:"Descriptions"`
}

type ExampleGroup struct {
	XMLName          xml.Name         `xml:"ExampleGroup"`
	GroupName        Names            `xml:"Names"`
	GroupDescription Descriptions     `xml:"Descriptions"`
	GroupProducts    []ExampleProduct `xml:"ExampleProduct"`
}

type SharedProperties struct {
	XMLName                  xml.Name `xml:"SharedProperties"`
	Assessments              string   `xml:"Assessments"`
	Tutoring                 string   `xml:"Tutoring"`
	UserComments             string   `xml:"UserComments"`
	SynchronousInteraction   string   `xml:"SynchronousInteraction"`
	AsynchronousInteraction  string   `xml:"AsynchronousInteraction"`
	AmountAccesses           string   `xml:"AmountAccesses"`
	Downloads                string   `xml:"Downloads"`
	ShowLearningObjectives   string   `xml:"ShowLearningObjectives"`
	Purpose                  string   `xml:"Purpose"`
	CloudBased               string   `xml:"CloudBased"`
	Intranet                 string   `xml:"Intranet"`
	VideoContent             string   `xml:"VideoContent"`
	StudentRoles             string   `xml:"StudentRoles"`
	DisplayEquations         string   `xml:"DisplayEquations"`
	WriteEquations           string   `xml:"WriteEquations"`
	TeachingTypePresentation string   `xml:"TeachingTypePresentation"`
	TeachingTypeDevelopment  string   `xml:"TeachingTypeDevelopment"`
	TeachingTypeExplorative  string   `xml:"TeachingTypeExplorative"`
}

type Description struct {
	XMLName  xml.Name `xml:"Description"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:"Text"`
}

type Descriptions struct {
	XMLName      xml.Name      `xml:"Descriptions"`
	Descriptions []Description `xml:"Description"`
}

type Name struct {
	XMLName  xml.Name `xml:"Name"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:"text,attr"`
}

type Names struct {
	XMLName xml.Name `xml:"Names"`
	Names   []Name   `xml:"Name"`
}

type ProductGroup struct {
	XMLName          xml.Name         `xml:"ProductGroup"`
	InternalName     string           `xml:"internalName,attr"`
	GroupName        Names            `xml:"Names"`
	GroupDescription Descriptions     `xml:"Descriptions"`
	SharedProperties SharedProperties `xml:"SharedProperties"`
	ExampleGroup     []ExampleGroup   `xml:"ExampleGroup"`
}

type ProductsGroup struct {
	XMLName  xml.Name       `xml:"ProductsGroup"`
	Products []ProductGroup `xml:"ProductGroup"`
}

type Data struct {
	XMLName            xml.Name      `xml:"xml"`
	ProductsCollection ProductsGroup ``
}
