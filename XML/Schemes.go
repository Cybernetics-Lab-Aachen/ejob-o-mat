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
	XMLName                               xml.Name `xml:"SharedProperties"`
	Exam                                  string   `xml:"Exam"`
	Assistant                             string   `xml:"Assistant"`
	UserComments                          string   `xml:"UserComments"`
	AnonymousUsers                        string   `xml:"AnonymousUsers"`
	LiveCollaboration                     string   `xml:"LiveCollaboration"`
	CommunityCollaboration                string   `xml:"CommunityCollaboration"`
	AppropriateCountStudents              string   `xml:"AppropriateCountStudents"`
	Downloads                             string   `xml:"Downloads"`
	Possibility2DeclareLearningObjectives string   `xml:"Possibility2DeclareLearningObjectives"`
	OperationType                         string   `xml:"OperationType"`
	Costs                                 string   `xml:"Costs"`
	AlsoFreeProducts                      string   `xml:"AlsoFreeProducts"`
	CloudBased                            string   `xml:"CloudBased"`
	Intranet                              string   `xml:"Intranet"`
	StandaloneSoftware                    string   `xml:"StandaloneSoftware"`
	SCROMSupport                          string   `xml:"SCROMSupport"`
	DeploymentNecessary                   string   `xml:"DeploymentNecessary"`
	VideoContent                          string   `xml:"VideoContent"`
	StudentRoles                          string   `xml:"StudentRoles"`
	TrackedProgress                       string   `xml:"TrackedProgress"`
	DisplayEquations                      string   `xml:"DisplayEquations"`
	WriteEquations                        string   `xml:"WriteEquations"`
	ContentType                           string   `xml:"ContentType"`
	HomeUse                               string   `xml:"HomeUse"`
	TeachingTypePresentation              string   `xml:"TeachingTypePresentation"`
	TeachingTypeDevelopment               string   `xml:"TeachingTypeDevelopment"`
	TeachingTypeExplorative               string   `xml:"TeachingTypeExplorative"`
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
