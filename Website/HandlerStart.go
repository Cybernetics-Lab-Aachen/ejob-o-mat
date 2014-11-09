package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

func HandlerStart(response http.ResponseWriter, request *http.Request) {
	pwd := request.FormValue(`PWD`)
	readSession := request.FormValue(`session`)

	if readSession == `` {
		if pwd != betaPassword {
			defer http.Redirect(response, request, `/`, 307)
			Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A wrong password was used.`, pwd)
			return
		}
	}

	lang := Tools.GetRequestLanguage(request)[0]
	data := PageStart{}
	data.Basis.Name = NAME
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language

	if readSession != `` {
		data.Basis.Session = readSession
	} else {
		data.Basis.Session = Tools.RandomGUID()
	}

	if strings.Contains(lang.Language, `de`) {
		data.TextStartButton = `Fragebogen beginnen`
		data.TextWelcome = `Herzlich Willkommen bei ` + NAME + `. ` + NAME + ` ist ein E-Learning-Empfehlungssystem
		mit dem Fokus auf die MINT-Wissenschaften. Ziel ist es, dass Sie einen Überblick über geeignete Tools erhalten,
		die Sie in Ihrer Veranstaltung einsetzen können. Wenn Sie mit der Schaltfläche unten den Fragebogen starten,
		werden Ihnen einige Frage gestellt. ` + NAME + ` wird dann, basierend auf Ihren Antworten, eine geeignete
		Empfehlung aussprechen.` + NAME + ` wird realisiert in dem Projekt ELLI, welches vom
		Bundesministereium für xyz gefördert wird.`
	} else {
		data.TextStartButton = `Start Questionnaire`
		data.TextWelcome = `Welcome to` + NAME + `. ` + NAME + ` is an e-learning recommendation system with the focus.
		to STEAM scieces. The aim of ` + NAME + ` is to provide a overview about suitable tools. Therefore, it should
		enable you to use these tools at your lectures, etc. If you start the questionnaire by pressing the button below,
		you will receive a few questions. Afterwards, ` + NAME + ` is then able to recommend e-learning solutions to you,
		based on your answers.` + NAME + ` is a product from the project ELLI. This project is supported from
		the federal agency of xyz, Germany.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`start`, response, data)
}
