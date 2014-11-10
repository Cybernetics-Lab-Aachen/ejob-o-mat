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

func HandlerQuestion11(response http.ResponseWriter, request *http.Request) {
	noQuestion := 11
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
	data.Button5Status = BUTTON_SHOW

	data.Button1Data = `none`
	data.Button2Data = `up10`
	data.Button3Data = `up25`
	data.Button4Data = `up100`
	data.Button5Data = `up1000`

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Keines`
		data.TextButton2 = `Bis 10 €`
		data.TextButton3 = `Bis 25 €`
		data.TextButton4 = `Bis 100 €`
		data.TextButton5 = `Bis 1000 €`
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Budget`
		data.TextQuestionBody = `Welches Budget haben Sie für die einmalige Anschaffung pro
		Studierenden?`
	} else {
		data.TextButton1 = `None`
		data.TextButton2 = `Up to USD 10`
		data.TextButton3 = `Up to USD 25`
		data.TextButton4 = `Up to USD 100`
		data.TextButton5 = `Up to USD 1000`
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Budget`
		data.TextQuestionBody = `What is your budget for the one-time purchase per student?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
