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
		data.TextButton2 = `Mittel (Dutzende)`
		data.TextButton3 = `Groß (Tausende)`
		data.TextButton4 = `Massen (Zehntausende+)`
		data.TextButton5 = ``
		data.TextQuestion = `Frage`
		data.TextQuestionTopic = `Anzahl Studierende`
		data.TextQuestionBody = `Welche der vier Größenangaben passt am besten zu Ihrem Szenario?
		Beachten Sie die angedachte Gruppengröße: Sollen alle Ihre Studierende zusammen das Tool
		nutzen oder werden die Studierenden in z.B. 4er-Gruppen eingeteilt?`
	} else {
		data.TextButton1 = `Small`
		data.TextButton2 = `Mid (dozens)`
		data.TextButton3 = `Huge (thousands)`
		data.TextButton4 = `Masses (ten-thousends+)`
		data.TextButton5 = ``
		data.TextQuestion = `Question`
		data.TextQuestionTopic = `Count of Students`
		data.TextQuestionBody = `Which of the four sizes is the best match for your scenario?
		Consider the conceived group size: Should all your students use the e-learning tool
		at the same time or do you divide the students into e.g. groups of four students?`
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}
