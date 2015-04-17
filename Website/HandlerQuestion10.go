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
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = readSession

	data.Button1Status = BUTTON_SHOW
	data.Button2Status = BUTTON_SHOW
	data.Button3Status = BUTTON_SHOW
	data.Button4Status = BUTTON_HIDDEN
	data.Button5Status = BUTTON_HIDDEN
	data.ButtonBackStatus = BUTTON_SHOW

	data.Button1Data = `1`
	data.Button2Data = `0`
	data.Button3Data = `*`
	data.Button4Data = ``
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.PreNoQuestion = fmt.Sprintf(`%d`, noQuestion-1)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.TextButton1 = `Ja`
		data.TextButton2 = `Nein`
		data.TextButton3 = `Enthaltung`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextBackButton = `Vorherige Frage`
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Cloud`
		data.TextQuestionBody = `Darf das E-Learning-Format nach den Statuten Ihrer Hochschule als Cloud-Dienst angeboten werden?`
	} else {
		data.Basis.Name = NAME_EN
		data.TextButton1 = `Yes`
		data.TextButton2 = `No`
		data.TextButton3 = `Skip question`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextBackButton = `Previous question`
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Cloud`
		data.TextQuestionBody = `Do the bylaws of your university allow an e-learning format provided by a cloud service?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
