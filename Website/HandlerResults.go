package Website

import (
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/Algorithm"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/XML"
	"net/http"
	"strings"
	"time"
)

func HandlerResults(response http.ResponseWriter, request *http.Request) {

	session := request.FormValue(`session`)
	lang := Tools.GetRequestLanguage(request)[0]
	answers := DB.LoadAnswers(session)
	assessedGroups := Algorithm.ExecuteAnswers(answers)
	groups := XML.GetData()
	resultSet := DB.Recommendation{}

	resultSet.CreateTimeUTC = time.Now().UTC()
	resultSet.ProductGroups = assessedGroups
	resultSet.Session = session

	data := PageResults{}
	data.Basis.Name = NAME
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = session
	data.Groups = groups.ProductsCollection.Products // Problem => Hier steckt nicht mehr drinnen, wie viel %! Brauche also noch ein Array hier drinnen!

	if strings.Contains(lang.Language, `de`) {

		data.LangPos = 0
		data.TextHeader = `Bitte wählen Sie eine Gruppe um die Details anzuzeigen`
		data.TextMatch = `Übereinstimmung mit Ihren Antworten`
		data.TextGroup = `Gruppe`
		data.TextExamples = `Beispiele für`
	} else {

		data.LangPos = 1
		data.TextHeader = `Please choose a group to show the details`
		data.TextMatch = `match with your answers`
		data.TextGroup = `Group`
		data.TextExamples = `Examples for`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`results`, response, data)
}
