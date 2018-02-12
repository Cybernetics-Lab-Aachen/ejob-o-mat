package Website

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SommerEngineering/ejob-o-mat/DB"
	"github.com/SommerEngineering/ejob-o-mat/DB/Scheme"
)

//HandlerAnswer stores the given answer and redirects to the next question or results page.
func HandlerAnswer(response http.ResponseWriter, request *http.Request) {
	noText := request.FormValue(`no`)
	session := request.FormValue(`session`)
	data := request.FormValue(`a`)
	important := request.FormValue(`important`)
	lang := request.FormValue(`lang`)
	answers, loadAnswersError := DB.LoadAnswers(session)
	no := 0
	weight := 1

	//Validate input
	if len(lang) > 6 || len(noText) > 2 || len(data) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Check if session exists, otherwise redirect to start page
	if loadAnswersError {
		http.Redirect(response, request, `/start`, 302)
		return
	}

	//Check for old session. Can't work with outdated data.
	if answers.SchemeVersion < Scheme.CurrentVersion {
		http.Redirect(response, request, `/start`, 302)
		return
	}

	//Check if question is marked as important
	if important == `important` {
		weight = 2
	}

	if value, errValue := strconv.Atoi(noText); errValue != nil {
		no = 0
	} else {
		no = value

		// Store the new answer:
		answers.Answers["Question"+strconv.Itoa(int(no))] = Scheme.Answer{TimeUTC: time.Now().UTC(), Data: data, Weight: byte(weight)}

		DB.UpdateAnswers(answers)
	}

	if no+1 > TOTAL_QUESTIONS {
		// Last questions already answered
		http.Redirect(response, request, fmt.Sprintf("/results?lang=%s&session=%s&amount=6", lang, session), 302)
	} else {
		http.Redirect(response, request, fmt.Sprintf("/question%d?lang=%s&session=%s", (no+1), lang, session), 302)
	}
}
