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

	data.Button1Data = `small`
	data.Button2Data = `mid`
	data.Button3Data = `huge`
	data.Button4Data = `masses`
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.PreNoQuestion = fmt.Sprintf(`%d`, noQuestion-1)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.TextButton1 = `Kleingruppen`
		data.TextButton2 = `Mittel (Dutzend(e))`
		data.TextButton3 = `Groß (Hundert(e))`
		data.TextButton4 = `Massen (Tausend(e)+)`
		data.TextButton5 = ``
		data.TextBackButton = `Vorherige Frage`
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Anzahl der Zugriffe`
		data.TextQuestionBody = `Wie viele Studierende besuchen Ihren Kurs bzw. Ihre Veranstaltung?`
	} else {
		data.Basis.Name = NAME_EN
		data.TextButton1 = `Small`
		data.TextButton2 = `Mid (dozen(s))`
		data.TextButton3 = `Huge (hundred(s))`
		data.TextButton4 = `Masses (thousand(s)+)`
		data.TextButton5 = ``
		data.TextBackButton = `Previous question`
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Count of Accesses`
		data.TextQuestionBody = `How many students are in your course?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
