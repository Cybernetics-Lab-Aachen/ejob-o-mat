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

func HandlerQuestion6(response http.ResponseWriter, request *http.Request) {
	noQuestion := 6
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
	data.Button4Status = BUTTON_SHOW
	data.Button5Status = BUTTON_HIDDEN
	data.ButtonBackStatus = BUTTON_SHOW
	data.ButtonInfoStatus = BUTTON_HIDDEN

	data.Button1Data = `1`
	data.Button2Data = `2`
	data.Button3Data = `3`
	data.Button4Data = `4`
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.PreNoQuestion = fmt.Sprintf(`%d`, noQuestion-1)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.TextButton1 = `1 bis 12`
		data.TextButton2 = `13 bis 48`
		data.TextButton3 = `49 bis 500`
		data.TextButton4 = `mehr als 500`
		data.TextButton5 = ``
		data.TextBackButton = `Vorherige Frage`
		data.TextImportant = `Diese Aussage ist mir besonderst wichtig`
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Anzahl der Zugriffe`
		data.TextQuestionBody = `Wie viele Studierende besuchen Ihren Kurs bzw. Ihre Veranstaltung?`
		data.QuestionInfoHeader = `Zusätzliche Hinweise`
		data.QuestionInfoClose = `Schließen`
		data.QuestionInfoText = ``
	} else {
		data.Basis.Name = NAME_EN
		data.TextButton1 = `1 til 12`
		data.TextButton2 = `13 til 48`
		data.TextButton3 = `49 til 500`
		data.TextButton4 = `more than 500`
		data.TextButton5 = ``
		data.TextBackButton = `Previous question`
		data.TextImportant = `This statement is important for me`
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Amount of Accesses`
		data.TextQuestionBody = `How many students are in your course?`
		data.QuestionInfoHeader = `Additional Information`
		data.QuestionInfoClose = `Close`
		data.QuestionInfoText = ``
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
