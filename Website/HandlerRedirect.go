package Website

import (
	"net/http"
)

func HandlerRedirect(response http.ResponseWriter, request *http.Request) {
	readSession := request.FormValue(`session`)
	if readSession != `` {
		http.Redirect(response, request, `/start?session=`+readSession, 302)
	} else {
		http.Redirect(response, request, `/start`, 302)
	}
}
