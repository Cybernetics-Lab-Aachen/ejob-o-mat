package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

func HandlerQuestion10(response http.ResponseWriter, request *http.Request) {
	noQuestion := 10
	readSession := request.FormValue(`session`)
	if readSession == `` {
		defer http.Redirect(response, request, "/", 307)
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A request without session.`)
		return
	}

	lang := Tools.GetRequestLanguage(request)[0]
	data := PageQuestion{}
	data.Basis.Name = NAME
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = readSession

	data.Button1Status = BUTTON_SHOW
	data.Button2Status = BUTTON_SHOW
	data.Button3Status = BUTTON_HIDDEN
	data.Button4Status = BUTTON_HIDDEN
	data.Button5Status = BUTTON_HIDDEN

	data.Button1Data = `support4lecture`
	data.Button2Data = `replace`
	data.Button3Data = ``
	data.Button4Data = ``
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Vorlesung/Übung etc. unterstützen`
		data.TextButton2 = `Präsenzveranstaltung ersetzen`
		data.TextButton3 = ``
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Einsatzzweck`
		data.TextQuestionBody = `Suchen Sie ein E-Learning-System zur Unterstützung Ihrer Vorlesung,
		Übung etc. oder eines um die Präsenzveranstaltung zu ersetzen?`
	} else {
		data.TextButton1 = `Support Lecture`
		data.TextButton2 = `Replace on-site attendance`
		data.TextButton3 = ``
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Purpose`
		data.TextQuestionBody = `Do you want an e-learning solution to support your lecture,
		exercise, etc. or a solution to replace the on-site attendance?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
