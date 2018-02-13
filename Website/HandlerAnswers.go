package Website

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SommerEngineering/ejob-o-mat/DB"
	"github.com/SommerEngineering/ejob-o-mat/DB/Scheme"
	"github.com/SommerEngineering/ejob-o-mat/XML"
)

//HandlerAnswer stores the given answer and redirects to the next question or results page.
func HandlerAnswer(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)
	data := request.FormValue(`a`)
	lang := request.FormValue(`lang`)
	survey, loadSurveyError := DB.LoadAnswers(session)

	// Check if session does not exists or is outdated
	if loadSurveyError || survey.SchemeVersion < Scheme.CurrentVersion {
		http.Redirect(response, request, fmt.Sprintf("/start?lang=%s", lang), 302)
		return
	}

	// Read question number
	no, errNo := strconv.Atoi(request.FormValue(`no`))
	if errNo != nil || no < 0 || no >= len(XML.Questions) { // Invalid question number
		http.Redirect(response, request, fmt.Sprintf("/start?lang=%s", lang), 302)
		return
	}

	//Check if question is marked as important
	weight := 1
	if request.FormValue(`important`) == `important` {
		weight = 2
	}

	//Validate input data
	dataValid := false
	for _, button := range XML.Questions[survey.Questions[no]].Buttons { // Given data should match one of the response options
		if data == button.Data {
			dataValid = true
		}
	}
	if !dataValid { // Data was not valid, take question again
		http.Redirect(response, request, fmt.Sprintf("/question%d?lang=%s&session=%s", no, lang, session), 302)
		return
	}

	// Store the new answer:
	survey.Answers[survey.Questions[no]] = Scheme.Answer{
		TimeUTC: time.Now().UTC(),
		Data:    data,
		Weight:  byte(weight)}
	DB.UpdateAnswers(survey)

	if no+1 >= len(XML.Questions) {
		// Last questions already answered
		http.Redirect(response, request, fmt.Sprintf("/results?lang=%s&session=%s&amount=6", lang, session), 302)
	} else {
		http.Redirect(response, request, fmt.Sprintf("/question%d?lang=%s&session=%s", no+1, lang, session), 302)
	}
}
