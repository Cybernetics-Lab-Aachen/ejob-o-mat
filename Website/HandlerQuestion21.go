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

func HandlerQuestion21(response http.ResponseWriter, request *http.Request) {
	noQuestion := 21
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
	data.Button3Status = BUTTON_SHOW
	data.Button4Status = BUTTON_SHOW
	data.Button5Status = BUTTON_HIDDEN

	data.Button1Data = `static`
	data.Button2Data = `change`
	data.Button3Data = `interact`
	data.Button4Data = `*`
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Statisch`
		data.TextButton2 = `Veränderbar`
		data.TextButton3 = `Interagierbar`
		data.TextButton4 = `Egal`
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Art der Inhalte`
		data.TextQuestionBody = `Welcher Art sind Ihre Inhalte für das E-Learning-Tool? Statische Inhalte lassen
		sich nur konsumieren (z.B. Präsentationen, Videos, Bücher), während veränderbare Inhalte von den Studierenden
		bearbeitet werden können (z.B. Wikis). Interagierbare Inhalte erlauben die Interaktion zwischen
		Studierenden und E-Learning-Tool (z.B. eine Simulation).`
	} else {
		data.TextButton1 = `Static`
		data.TextButton2 = `Changeable`
		data.TextButton3 = `Interactable`
		data.TextButton4 = `Does not matter`
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Kind of Content`
		data.TextQuestionBody = `Which kind of content do you want to provide for the e-learning solution? Static
		content is just consumable (e.g. presentations, videos, books), wherever changeable content can be edit
		by the students (e.g. Wikis). Interactable content allows the students to interact with the e-learning
		solution (e.g. a simulation).`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
