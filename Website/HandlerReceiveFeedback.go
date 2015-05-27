package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"net/http"
	"strconv"
)

func HandlerReceiveFeedback(response http.ResponseWriter, request *http.Request) {

	session := request.FormValue(`session`)
	text := request.FormValue(`text`)
	sourceLocation := request.FormValue(`sourceLocation`)
	ratingRAW := request.FormValue(`rating`)
	rating := byte(255)

	if value, err := strconv.Atoi(ratingRAW); err != nil {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityLow, LM.ImpactNone, LM.MessageNamePARSE, `Was not able to parse the rating from the feedback.`)
	} else {
		rating = byte(value)
	}

	feedback := Scheme.Feedback{}
	feedback.Rating = rating
	feedback.Session = session
	feedback.Text = text
	feedback.SourceLocation = sourceLocation

	DB.StoreFeedback(feedback)
	http.Redirect(response, request, sourceLocation, 302)
}
