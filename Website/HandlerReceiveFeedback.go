package Website

import (
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"net/http"
	"strconv"
	"strings"
)

func HandlerReceiveFeedback(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)
	text := request.FormValue(`text`)
	sourceLocation := request.FormValue(`sourceLocation`)
	ratingRAW := request.FormValue(`rating`)
	rating := byte(255)

	if len(text) > 2048 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(session) != 36 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(sourceLocation) == 0 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(sourceLocation) > 1024 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if !strings.HasPrefix(sourceLocation, `http://`) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if ratingRAW != `` && len(ratingRAW) != 1 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

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
