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
		answers.SchemeVersion = Scheme.CURRENT_VERSION
		answers.Session = data.Basis.Session
		answers.CreateTimeUTC = time.Now().UTC()
		DB.StoreNewAnswers(answers)
	}

	if strings.Contains(lang.Language, `de`) {
		var nameDEPlain = strings.Replace(NAME_DE, `.net`, ``, 1)
		data.Basis.Name = NAME_DE
		data.TextStartButton = `Fragebogen beginnen`
		data.TextProject = `Projekt`
		data.TextExecuted = `Durchgeführt von`
		data.TextPromoted = `Gefördert von`
		data.TextWelcome = `Herzlich willkommen beim ` + nameDEPlain + `. Der ` + nameDEPlain + ` ist ein E-Learning-Empfehlungssystem
		mit dem Fokus auf die MINT-Wissenschaften. Ziel ist es, dass Sie einen Überblick über geeignete Formate erhalten,
		die Sie in Ihrer Veranstaltung einsetzen können. Wenn Sie mit der Schaltfläche unten den Fragebogen starten,
		werden Ihnen einige Fragen gestellt. Basierend auf Ihren Antworten wird dann eine geeignete Empfehlung
		ausgegeben. Der ` + nameDEPlain + ` wird realisiert in dem Projekt ELLI (Exzellentes Lehren und Lernen in den
		Ingenieurswissenschaften), welches vom Bundesministerium für Bildung und Forschung gefördert wird.`
	} else {
		var nameENPlain = strings.Replace(NAME_EN, `.net`, ``, 1)
		data.TextStartButton = `Start Questionnaire`
		data.Basis.Name = NAME_EN
		data.TextProject = `Project`
		data.TextExecuted = `Executed by`
		data.TextPromoted = `Promoted by`
		data.TextWelcome = `Welcome to ` + NAME_EN + `. The ` + nameENPlain + ` is an e-learning recommendation system
		with focus on STEM sciences. Our aim is to provide you with an overview about suitable tools which you can use
		e.g. in your courses, lectures, etc. As soon as you start the questionnaire by pressing the button below, you
		will receive a few questions. Afterwards, the ` + nameENPlain + ` will be able to recommend e-learning formats for
		you, based on your answers. The ` + nameENPlain + ` is part of the project ELLI. This project is supported by
		the Federal Ministry of Education and Research, Germany.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`start`, response, data)
}
