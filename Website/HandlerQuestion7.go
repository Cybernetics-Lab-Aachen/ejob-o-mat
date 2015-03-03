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

func HandlerQuestion7(response http.ResponseWriter, request *http.Request) {
	noQuestion := 7
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

	data.Button1Data = `small`
	data.Button2Data = `mid`
	data.Button3Data = `huge`
	data.Button4Data = `masses`
	data.Button5Data = ``

	data.NoQuestion = fmt.Sprintf(`%d`, noQuestion)
	data.NoQuestions = totalQuestions
	data.Progress = fmt.Sprintf("%d", (int((float32(noQuestion) / float32(TOTAL_QUESTIONS)) * 100.0)))

	if strings.Contains(lang.Language, `de`) {
		data.TextButton1 = `Kleingruppen`
		data.TextButton2 = `Mittel (Dutzend(e))`
		data.TextButton3 = `Groß (Hundert(e))`
		data.TextButton4 = `Massen (Tausend(e)+)`
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Anzahl der gleichzeitgen Zugriffe`
		data.TextQuestionBody = `Wie viele Studierende greifen gleichzeitig auf das E-Learning-System zu?
		Bedenken Sie dabei, dass z.B. bei einer Gruppenarbeit nur die Anzahl der Gruppen zählt.`
	} else {
		data.TextButton1 = `Small`
		data.TextButton2 = `Mid (dozen(s))`
		data.TextButton3 = `Huge (hundred(s))`
		data.TextButton4 = `Masses (thousand(s)+)`
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Count of Concurrent Accesses`
		data.TextQuestionBody = `How many students using the e-learning system concurrently? Please also
		consider in case of e.g. group work, that instead the amount of groups is necessary.`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
