package Website

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandlerAnswer(response http.ResponseWriter, request *http.Request) {
	noText := request.FormValue(`no`)
	session := request.FormValue(`session`)
	//data := request.FormValue(`a`)
	lang := request.FormValue(`lang`)

	no := 0
	if value, errValue := strconv.Atoi(noText); errValue != nil {
		no = 0
	} else {
		no = value
	}

	if no+1 > 28 {
		http.Redirect(response, request, `/result`, 307)
	} else {
		http.Redirect(response, request, fmt.Sprintf("/question%d?lang=%s&session=%s", (no+1), lang, session), 307)
	}
}
