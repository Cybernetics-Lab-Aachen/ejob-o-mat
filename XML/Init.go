package XML

import (
	"encoding/xml"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/StaticFiles"
)

func init() {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Init the XML component.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Done init the XML component.`)

	xmlData := StaticFiles.FindAndReadFile(`Data.xml`)
	if errParse := xml.Unmarshal(xmlData, &data); errParse != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNamePARSE, `Was not able to parse the XML data.`, errParse.Error())
	} else {
		Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameINIT, `The XML data file was parsed without issues.`)
	}
}
