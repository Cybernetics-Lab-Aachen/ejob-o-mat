package XML

import (
	"encoding/xml"

	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/StaticFiles"
)

//init parses the Data.xml file and stores its representation in memory.
func init() {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Init the XML component.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Done init the XML component.`)

	xmlData := StaticFiles.FindAndReadFile(`Data.xml`)
	if errParse := xml.Unmarshal(xmlData, &data); errParse != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNamePARSE, `Was not able to parse the XML data.`, errParse.Error())
		return
	}
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameINIT, `The XML data file was parsed without issues.`)

	Questions = make(map[string]QuestionGroup)
	for i, question := range data.QuestionsCollection.Questions {
		// Make buttons searchable by data
		buttonsByData := make(map[string]QuestionButton)
		for _, button := range question.Buttons {
			buttonsByData[button.Data] = button
		}
		data.QuestionsCollection.Questions[i].ButtonsByData = buttonsByData

		Questions[question.InternalName] = question
	}
}
