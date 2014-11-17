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

func HandlerQuestion16(response http.ResponseWriter, request *http.Request) {
	noQuestion := 16
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
	data.Button4Status = BUTTON_HIDDEN
	data.Button5Status = BUTTON_HIDDEN

	data.Button1Data = `1`
	data.Button2Data = `0`
	data.Button3Data = `*`
	data.Button4Data = ``
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Ja`
		data.TextButton2 = `Nein`
		data.TextButton3 = `Egal`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Video-Inhalte`
		data.TextQuestionBody = `Soll die Möglichkeit bestehen, Video-Inhalte über das E-Learning-Tool anzubieten?`
	} else {
		data.TextButton1 = `Yes`
		data.TextButton2 = `No`
		data.TextButton3 = `Does not matter`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Video Content`
		data.TextQuestionBody = `Should it possible to provide video content with the e-learning solution?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
