package Website

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/XML"
)

//HandlerQuestion displays the current question.
func HandlerQuestion(response http.ResponseWriter, request *http.Request) {
	readSession := request.FormValue(`session`)

	// Check if session is valid
	if readSession == `` {
		defer http.Redirect(response, request, "/", 302)
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A request without session.`)
		return
	}
	if len(readSession) != 36 {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSTATE, `Session's length was not valid!`, readSession)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	//Get question number
	noQuestion, _ := strconv.Atoi(request.URL.Path[9:]) //Extract 14 from e.g. "/question14"
	if noQuestion < 1 || noQuestion > TOTAL_QUESTIONS {
		//This should never happen, as this function only handles valid questions, but one never knows ...
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Prepare data for html template
	data := PageQuestion{}
	data.Basis.Session = readSession
	data.Basis.Version = VERSION
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")

	//Read question number
	data.Progress = noQuestion
	data.NoQuestion = fmt.Sprintf(`%d`, data.Progress)
	data.PreNoQuestion = fmt.Sprintf(`%d`, data.Progress-1)
	data.NoQuestions = totalQuestions

	//Detect language
	lang := Tools.GetRequestLanguage(request)[0]
	data.Basis.Lang = lang.Language
	var langID int
	if strings.Contains(lang.Language, `de`) {
		langID = LANG_DE
		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
	} else {
		langID = LANG_EN
		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
	}

	//Prepare localized strings
	groups := XML.GetData()
	questionGroup := groups.QuestionsCollection.Questions[data.Progress-1]
	strings := groups.QuestionStrings

	data.TextBackButton = strings.TextBackButton[langID].Text
	data.TextImportant = strings.TextImportant[langID].Text
	data.TextQuestion = strings.TextQuestion[langID].Text
	data.QuestionInfoHeader = strings.QuestionInfoHeader[langID].Text
	data.QuestionInfoClose = strings.QuestionInfoClose[langID].Text

	data.TextQuestionTopic = questionGroup.Topics[langID].Text
	data.TextQuestionBody = questionGroup.QuestionBodies[langID].Text

	//Show additional information if available
	if questionGroup.Hints != nil {
		data.QuestionInfoText = questionGroup.Hints[langID].Text
		data.ButtonInfoStatus = BUTTON_SHOW
	} else { //Otherwise disable info button
		data.QuestionInfoText = ``
		data.ButtonInfoStatus = BUTTON_HIDDEN
	}

	//Hide back button on first question
	if noQuestion == 1 {
		data.ButtonBackStatus = BUTTON_HIDDEN
	} else {
		data.ButtonBackStatus = BUTTON_SHOW
	}

	//Option Buttons
	data.Buttons = make([]PageQuestionButton, len(questionGroup.Buttons))
	for i, button := range questionGroup.Buttons {
		data.Buttons[i].Data = button.Data
		data.Buttons[i].Text = button.Texts[langID].Text
	}

	//Store questionaire start time (time of first question)
	if noQuestion == 1 {
		answers, _ := DB.LoadAnswers(readSession)
		answers.StartTimeQ1 = time.Now().UTC()
		DB.UpdateAnswers(answers)
	}

	// Finally, execute the template
	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}

//GetProgressState returns the css class representing the progress.
func (pg PageQuestion) GetProgressState(pos int) (ret string) {
	if pg.Progress >= pos {
		ret = ` progressitemdone`
	}
	return
}
