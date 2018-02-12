package Website

import (
	"net/http"

	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

//HandlerRedirect redirects to start page while keeping the session, if present.
func HandlerRedirect(response http.ResponseWriter, request *http.Request) {
	readSession := request.FormValue(`session`)
	if readSession != `` && len(readSession) != 36 {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSTATE, `Session's length was not valid!`, readSession)
		readSession = ``
	}

	if readSession != `` {
		http.Redirect(response, request, `/start?session=`+readSession, 302)
	} else {
		http.Redirect(response, request, `/start`, 302)
	}
}
