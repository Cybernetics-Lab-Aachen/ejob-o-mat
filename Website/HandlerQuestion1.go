package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/DB"
	"net/http"
	"strings"
	"time"
)

func HandlerQuestion1(response http.ResponseWriter, request *http.Request) {
	noQuestion := 1
	readSession := request.FormValue(`session`)
	if readSession == `` {
		defer http.Redirect(response, request, "/", 307)
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A request without session.`)
		return
	}

	answers := DB.LoadAnswers(readSession)
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
		data.TextButton1 = `Ja, es ist erforderlich`
		data.TextButton2 = `Nein, ich möchte diese Funktion nicht`
		data.TextButton3 = `Nein, es ist nicht erforderlich`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Klausuren`
		data.TextQuestionBody = `Ist es erforderlich, dass das E-Learning-Tool dazu verwendet werden kann, um Klausuren durchzuführen?`
	} else {
		data.TextButton1 = `Yes`
		data.TextButton2 = `No`
		data.TextButton3 = `No`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Exams`
		data.TextQuestionBody = `Is it necessary, that the e-learning solution is able to perform proper exams?`
	}

	answers.StartTimeQ1 = time.Now().UTC()
	DB.UpdateAnswers(answers)

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
