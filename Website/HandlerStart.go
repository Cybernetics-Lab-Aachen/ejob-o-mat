package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"net/http"
	"strings"
	"time"
)

func HandlerStart(response http.ResponseWriter, request *http.Request) {
	pwd := request.FormValue(`PWD`)
	readSession := request.FormValue(`session`)

	if readSession == `` || DB.LoadAnswers(readSession).Session == `` {
		if pwd != betaPassword {
			defer http.Redirect(response, request, `/`, 307)
			Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A wrong password was used.`, pwd)
			return
		}
	}

	lang := Tools.GetRequestLanguage(request)[0]
	data := PageStart{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language

	if readSession != `` {
		data.Basis.Session = readSession
	} else {
		data.Basis.Session = Tools.RandomGUID()
		answers := Scheme.Answers{}
		answers.Session = data.Basis.Session
		answers.CreateTimeUTC = time.Now().UTC()
		DB.StoreNewAnswers(answers)
	}

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.TextStartButton = `Fragebogen beginnen`
		data.TextProject = `Projekt`
		data.TextExecuted = `Durchgeführt von`
		data.TextPromoted = `Gefördert von`
		data.TextWelcome = `Herzlich Willkommen bei ` + NAME_DE + `. ` + NAME_DE + ` ist ein E-Learning-Empfehlungssystem
		mit dem Fokus auf die MINT-Wissenschaften. Ziel ist es, dass Sie einen Überblick über geeignete Tools erhalten,
		die Sie in Ihrer Veranstaltung einsetzen können. Wenn Sie mit der Schaltfläche unten den Fragebogen starten,
		werden Ihnen einige Frage gestellt. ` + NAME_DE + ` wird dann, basierend auf Ihren Antworten, eine geeignete
		Empfehlung aussprechen. ` + NAME_DE + ` wird realisiert in dem Projekt ELLI, welches vom
		Bundesministerium für Bildung und Forschung gefördert wird.`
	} else {
		data.TextStartButton = `Start Questionnaire`
		data.Basis.Name = NAME_EN
		data.TextProject = `Project`
		data.TextExecuted = `Executed by`
		data.TextPromoted = `Promoted by`
		data.TextWelcome = `Welcome to ` + NAME_EN + `. ` + NAME_EN + ` is an e-learning recommendation system with the focus.
		to STEAM scieces. The aim of ` + NAME_EN + ` is to provide a overview about suitable tools. Therefore, it should
		enable you to use these tools at your lectures, etc. If you start the questionnaire by pressing the button below,
		you will receive a few questions. Afterwards, ` + NAME_EN + ` is then able to recommend e-learning solutions to you,
		based on your answers. ` + NAME_EN + ` is a product from the project ELLI. This project is supported from
		the Federal Ministry of Education and Research, Germany.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`start`, response, data)
}
