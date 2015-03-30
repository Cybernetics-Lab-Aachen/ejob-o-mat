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

func HandlerQuestion23(response http.ResponseWriter, request *http.Request) {
	noQuestion := 23
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
		data.TextButton1 = `Ja`
		data.TextButton2 = `Nein`
		data.TextButton3 = `Unbekannt`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextBackButton = `Vorherige Frage`
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Erarbeitende Lehrverfahren`
		data.TextQuestionBody = `Planen Sie erarbeitende Lehrverfahren einzusetzen wie z.B. problembasiertes Lernen,
		angeleitetes Praktikum/Übungen, einen Fachdialog zwischen Studierenden und Lehrenden, etc.?`
	} else {
		data.TextButton1 = `Yes`
		data.TextButton2 = `No`
		data.TextButton3 = `Unknown`
		data.TextButton4 = ``
		data.TextButton5 = ``
		data.TextBackButton = `Previous question`
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Developing Teaching Methods`
		data.TextQuestionBody = `Do you want to use developing teaching methods like e.g. problem-based learning,
		guided traineesheps or practices, a professional dialogue between students and teacher/professor, etc.?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
