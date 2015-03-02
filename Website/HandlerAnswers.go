package Website

import (
	"fmt"
	"github.com/SommerEngineering/Re4EEE/DB"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func HandlerAnswer(response http.ResponseWriter, request *http.Request) {

	noText := request.FormValue(`no`)
	session := request.FormValue(`session`)
	data := request.FormValue(`a`)
	lang := request.FormValue(`lang`)
	answers := DB.LoadAnswers(session)
	no := 0

	if value, errValue := strconv.Atoi(noText); errValue != nil {
		no = 0
	} else {
		no = value

		// Store the new answer:
		re := reflect.ValueOf(&answers)
		element := re.Elem()
		field := element.FieldByName(fmt.Sprintf("A%dTimeUTC", no))
		field.Set(reflect.ValueOf(time.Now().UTC()))
		field = element.FieldByName(fmt.Sprintf("A%dData", no))
		field.SetString(data)
		DB.UpdateAnswers(answers)
	}

	if no+1 > 24 {
		http.Redirect(response, request, fmt.Sprintf("/results?lang=%s&session=%s&amount=8", lang, session), 307)
	} else {
		http.Redirect(response, request, fmt.Sprintf("/question%d?lang=%s&session=%s", (no+1), lang, session), 307)
	}
}
