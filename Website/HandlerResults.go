package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/Algorithm"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandlerResults(response http.ResponseWriter, request *http.Request) {

	session := request.FormValue(`session`)
	amountText := request.FormValue(`amount`)
	lang := Tools.GetRequestLanguage(request)[0]
	answers := DB.LoadAnswers(session)
	groups := XML.GetData()
	amountValue := -1
	resultSet := Scheme.Recommendation{}

	if !DB.CheckRecommendation(session) {

		assessedGroups := Algorithm.ExecuteAnswers(answers)

		resultSet.ProductGroups = assessedGroups
		resultSet.CreateTimeUTC = time.Now().UTC()
		resultSet.Session = session
		DB.StoreRecommendation(resultSet)

	} else {
		resultSet = DB.LoadRecommendation(session)
	}

	if value, errConv := strconv.Atoi(amountText); errConv != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameREQUEST, `Cannot read the amount value!`, amountText, errConv.Error())
	} else {
		amountValue = value
	}

	if amountValue >= 1 {
		resultSet.ProductGroups = resultSet.ProductGroups[0:amountValue]
	}

	data := PageResults{}
	data.Basis.Name = NAME
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = session
	data.Groups = groups.ProductsCollection.Products
	data.Recommendation = resultSet
	data.AmountCurrent = amountValue

	if strings.Contains(lang.Language, `de`) {

		if amountValue > 0 {
			data.TextAllGroups = `Alle Gruppen anzeigen`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Top 8 anzeigen`
			data.AmountToggle = 8
		}

		data.LangPos = 0
		data.TextMatch = `Übereinstimmung mit Ihren Antworten`
		data.TextGroup = `Gruppe`
		data.TextExamples = `Beispiele für`
		data.TextOptionen = `Optionen`
		data.TextHeaderPrefix = `Ihre E-Learning-Empfehlung`
		data.TextRestart = `Den Fragebogen erneut starten`
		data.TextHeader = `Unten sehen Sie eine Empfehlung für verschiedene Gruppen von E-Learning-Systemen, basierend
		auf Ihren Antworten. Bitte wählen Sie eine Gruppe um unten die Details inkl. Beispiele zu sehen. Sie sehen
		zunächst nur die Top 8 aller E-Learning-Gruppen: Mit der nachfolgenden Schaltfläche (siehe Optionen) können Sie
		auch alle Gruppenergebnisse einsehen.`

	} else {

		if amountValue > 0 {
			data.TextAllGroups = `Show all groups`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Show top 8`
			data.AmountToggle = 8
		}

		data.LangPos = 1
		data.TextMatch = `match with your answers`
		data.TextGroup = `Group`
		data.TextOptionen = `Options`
		data.TextExamples = `Examples for`
		data.TextHeaderPrefix = `Your E-Learning Recommendation`
		data.TextRestart = `Restart the questionnaire`
		data.TextHeader = `Please find below your e-learning recommendation, based on your answers. Please choose a
		group to show the details with the examples. Initially, just the top 8 groups are visible. With the options
		below, you can access also the results of all groups.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`results`, response, data)
}
