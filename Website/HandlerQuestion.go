package Website

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/ejob-o-mat/DB"
	"github.com/SommerEngineering/ejob-o-mat/XML"
)

//HandlerQuestion displays the current question.
func HandlerQuestion(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)

	// Check if session is valid
	if session == `` {
		defer http.Redirect(response, request, "/", 302)
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelSECURITY, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameUSER, `A request without session.`)
		return
	}
	if len(session) != 36 {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSTATE, `Session's length was not valid!`, session)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	//Get question number
	noQuestion, _ := strconv.Atoi(request.URL.Path[9:]) //Extract 14 from e.g. "/question14"
	if noQuestion < 0 || noQuestion >= len(XML.Questions) {
		//This should never happen, as this function only handles valid questions, but one never knows ...
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Prepare data for html template
	data := PageQuestion{}
	data.Basis.Session = session
	data.Basis.Version = VERSION
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")

	//Read question number
	data.NoQuestion = noQuestion
	data.NoQuestions = len(XML.Questions)

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

	survey, loadsurveyError := DB.LoadAnswers(session)
	// Check if session exists, otherwise redirect to start page
	if loadsurveyError {
		http.Redirect(response, request, `/start`, 302)
		return
	}

	//Prepare localized strings
	groups := XML.GetData()
	questionGroup := XML.Questions[survey.Questions[noQuestion]]
	strings := groups.QuestionStrings

	data.TextBackButton = strings.TextBackButton[langID].Text
	data.TextImportant = strings.TextImportant[langID].Text
	data.TextQuestion = strings.TextQuestion[langID].Text
	data.QuestionInfoHeader = strings.QuestionInfoHeader[langID].Text
	data.QuestionInfoClose = strings.QuestionInfoClose[langID].Text

	data.TextQuestionTopic = questionGroup.Topics[langID].Text
	data.TextQuestionBody = questionGroup.QuestionBodies[langID].Text

	// Add image
	if len(questionGroup.QuestionImages) != 0 {
		data.QuestionImageFilename = questionGroup.QuestionImages[0].Filename
		data.QuestionImageSource = questionGroup.QuestionImages[0].Source
	} else {
		data.QuestionImageFilename = "/staticFiles/images/home.gif"
	}

	//Show additional information if available
	if questionGroup.Hints != nil {
		data.QuestionInfoText = questionGroup.Hints[langID].Text
		data.ButtonInfoStatus = BUTTON_SHOW
	} else { //Otherwise disable info button
		data.QuestionInfoText = ``
		data.ButtonInfoStatus = BUTTON_HIDDEN
	}

	//Option Buttons
	data.Buttons = make([]PageQuestionButton, len(questionGroup.Buttons))
	for i, button := range questionGroup.Buttons {
		data.Buttons[i].Data = button.Data
		data.Buttons[i].Text = button.Texts[langID].Text
	}

	// Progress states
	data.ProgressStates = make([]string, data.NoQuestions)
	for pos := 0; pos < data.NoQuestions; pos++ {
		if _, ok := survey.Answers[survey.Questions[pos]]; ok {
			data.ProgressStates[pos] = ` progressitemdone`
		} else {
			data.ProgressStates[pos] = ``
		}
	}

	data.ShowImportant = questionGroup.Weighted

	//Store questionaire start time (time of first question)
	if noQuestion == 0 {
		survey.StartTimeQ1 = time.Now().UTC()
		DB.UpdateAnswers(survey)
	}

	// Finally, execute the template
	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`question`, response, data)
}

// PageQuestion contains data for the question template.
type PageQuestion struct {
	Basis                 Basis
	NoQuestion            int
	NoQuestions           int
	TextQuestion          string
	QuestionInfoText      string
	QuestionInfoHeader    string
	QuestionInfoClose     string
	TextQuestionTopic     string
	TextQuestionBody      string
	QuestionImageFilename string
	QuestionImageSource   string
	TextBackButton        string
	TextImportant         string
	ButtonInfoStatus      string
	Buttons               []PageQuestionButton
	ProgressStates        []string
	ShowImportant         bool
}

// PageQuestionButton contains data for a single answer button.
type PageQuestionButton struct {
	Text string
	Data string
}
