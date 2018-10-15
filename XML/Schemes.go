package XML

import (
	"encoding/xml"
)

type ExampleProduct struct {
	XMLName             xml.Name      `xml:"ExampleProduct"`
	Website             string        `xml:"website,attr"`
	ProductNames        []Name        `xml:"Name"`
	ProductDescriptions []Description `xml:"Description"`
}

type Description struct {
	XMLName  xml.Name `xml:"Description"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:",chardata"`
}

type Name struct {
	XMLName  xml.Name `xml:"Name"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:"text,attr"`
}

type ProductGroup struct {
	XMLName           xml.Name         `xml:"ProductGroup"`
	InternalName      string           `xml:"internalName,attr"`
	GroupNames        []Name           `xml:"Name"`
	GroupDescriptions []Description    `xml:"Description"`
	ExampleProducts   []ExampleProduct `xml:"ExampleProduct"`
}

type ProductsGroup struct {
	XMLName  xml.Name       `xml:"ProductsGroup"`
	Products []ProductGroup `xml:"ProductGroup"`
}

type Data struct {
	XMLName             xml.Name        `xml:"xml"`
	ProductsCollection  ProductsGroup   ``
	QuestionsCollection QuestionsGroup  `xml:"QuestionsGroup"`
	QuestionStrings     QuestionStrings `xml:"QuestionStrings"`
	ResultStrings       ResultStrings   `xml:"ResultStrings"`
}

type QuestionsGroup struct {
	XMLName   xml.Name        `xml:"QuestionsGroup"`
	Questions []QuestionGroup `xml:"QuestionGroup"`
}

type QuestionGroup struct {
	XMLName        xml.Name                  `xml:"QuestionGroup"`
	InternalName   string                    `xml:"internalName,attr"`
	Topics         []Topic                   `xml:"Topic"`
	QuestionBodies []QuestionBody            `xml:"QuestionBody"`
	Hints          []QuestionHint            `xml:"Hint"`
	Buttons        []QuestionButton          `xml:"Button"`
	Weighted       bool                      `xml:"weighted,attr"`
	ButtonsByData  map[string]QuestionButton `xml:"-"`
}

type Topic struct {
	XMLName  xml.Name `xml:"Topic"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:"text,attr"`
}

type QuestionBody struct {
	XMLName  xml.Name `xml:"QuestionBody"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:",chardata"`
}

type QuestionHint struct {
	XMLName  xml.Name `xml:"Hint"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:",chardata"`
}

type QuestionButton struct {
	XMLName xml.Name     `xml:"Button"`
	Data    string       `xml:"data,attr"`
	Texts   []TextButton `xml:"TextButton"`
}

type TextButton struct {
	XMLName  xml.Name `xml:"TextButton"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:"text,attr"`
}

type QuestionStrings struct {
	XMLName            xml.Name `xml:"QuestionStrings"`
	TextBackButton     []String `xml:"TextBackButton>String"`
	TextImportant      []String `xml:"TextImportant>String"`
	TextQuestion       []String `xml:"TextQuestion>String"`
	QuestionInfoHeader []String `xml:"QuestionInfoHeader>String"`
	QuestionInfoClose  []String `xml:"QuestionInfoClose>String"`
}

type ResultStrings struct {
	XMLName          xml.Name `xml:"ResultStrings"`
	GoodInfluence    []String `xml:"GoodInfluence>String"`
	BadInfluence     []String `xml:"BadInfluence>String"`
	TextHeaderAll    []String `xml:"TextHeaderAll>String"`
	TextHeaderNone   []String `xml:"TextHeaderNone>String"`
	TextHeaderSome   []String `xml:"TextHeaderSome>String"`
	TextHeaderPrefix []String `xml:"TextHeaderPrefix>String"`
	TextMatch        []String `xml:"TextMatch>String"`
	TextGroup        []String `xml:"TextGroup>String"`
	TextExamples     []String `xml:"TextExamples>String"`
	TextRestart      []String `xml:"TextRestart>String"`
	TextOptionen     []String `xml:"TextOptionen>String"`
	TextResults      []String `xml:"TextResults>String"`
}

type String struct {
	XMLName  xml.Name `xml:"String"`
	Language string   `xml:"language,attr"`
	Text     string   `xml:",chardata"`
}
