package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Re4EEE/Algorithm"
	"github.com/SommerEngineering/Re4EEE/DB"
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func HandlerAnswer(response http.ResponseWriter, request *http.Request) {
	noText := request.FormValue(`no`)
	session := request.FormValue(`session`)
	data := request.FormValue(`a`)
	important := request.FormValue(`important`)
	lang := request.FormValue(`lang`)
	answers := DB.LoadAnswers(session)
	no := 0
	weight := 1

	if len(lang) > 6 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(session) != 36 {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSTATE, `Session's length was not valid!`, session)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(noText) > 2 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(data) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if important == `important` {
		weight = 2
	}

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

		field = element.FieldByName(fmt.Sprintf("A%dWeight", no))
		field.Set(reflect.ValueOf(byte(weight)))

		DB.UpdateAnswers(answers)
	}

	if no+1 > TOTAL_QUESTIONS {
		http.Redirect(response, request, fmt.Sprintf("/results?lang=%s&session=%s&amount=6", lang, session), 302)
	} else {
		http.Redirect(response, request, fmt.Sprintf("/question%d?lang=%s&session=%s", (no+1), lang, session), 302)
	}
}

func HandlerVRAnswers(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)
	creationTimeUTC := request.FormValue(`creationTimeUTC`)

	a1 := request.FormValue(`a1`)
	a2 := request.FormValue(`a2`)
	a3 := request.FormValue(`a3`)
	a4 := request.FormValue(`a4`)
	a5 := request.FormValue(`a5`)
	a6 := request.FormValue(`a6`)
	a7 := request.FormValue(`a7`)
	a8 := request.FormValue(`a8`)
	a9 := request.FormValue(`a9`)
	a10 := request.FormValue(`a10`)
	a11 := request.FormValue(`a11`)
	a12 := request.FormValue(`a12`)
	a13 := request.FormValue(`a13`)
	a14 := request.FormValue(`a14`)
	a15 := request.FormValue(`a15`)
	a16 := request.FormValue(`a16`)
	a17 := request.FormValue(`a17`)
	a18 := request.FormValue(`a18`)

	a1T := request.FormValue(`a1T`)
	a2T := request.FormValue(`a2T`)
	a3T := request.FormValue(`a3T`)
	a4T := request.FormValue(`a4T`)
	a5T := request.FormValue(`a5T`)
	a6T := request.FormValue(`a6T`)
	a7T := request.FormValue(`a7T`)
	a8T := request.FormValue(`a8T`)
	a9T := request.FormValue(`a9T`)
	a10T := request.FormValue(`a10T`)
	a11T := request.FormValue(`a11T`)
	a12T := request.FormValue(`a12T`)
	a13T := request.FormValue(`a13T`)
	a14T := request.FormValue(`a14T`)
	a15T := request.FormValue(`a15T`)
	a16T := request.FormValue(`a16T`)
	a17T := request.FormValue(`a17T`)
	a18T := request.FormValue(`a18T`)

	if len(session) != 36 {
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityCritical, LM.ImpactCritical, LM.MessageNameSTATE, `Session's length was not valid!`, session)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(creationTimeUTC) != 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a1) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a2) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a3) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a4) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a5) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a6) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a7) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a8) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a9) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a10) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a11) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a12) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a13) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a14) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a15) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a16) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a17) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a18) > 16 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a1T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a2T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a3T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a4T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a5T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a6T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a7T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a8T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a9T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a10T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a11T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a12T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a13T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a14T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a15T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a16T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a17T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if len(a18T) > 45 {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	answer := Scheme.Answers{}
	answer.CreateTimeUTC, _ = time.Parse(time.RFC1123, creationTimeUTC)
	answer.StartTimeQ1 = answer.CreateTimeUTC
	answer.Session = session
	answer.A1Data = a1
	answer.A2Data = a2
	answer.A3Data = a3
	answer.A4Data = a4
	answer.A5Data = a5
	answer.A6Data = a6
	answer.A7Data = a7
	answer.A8Data = a8
	answer.A9Data = a9
	answer.A10Data = a10
	answer.A11Data = a11
	answer.A12Data = a12
	answer.A13Data = a13
	answer.A14Data = a14
	answer.A15Data = a15
	answer.A16Data = a16
	answer.A17Data = a17
	answer.A18Data = a18
	answer.A1TimeUTC, _ = time.Parse(time.RFC1123, a1T)
	answer.A2TimeUTC, _ = time.Parse(time.RFC1123, a2T)
	answer.A3TimeUTC, _ = time.Parse(time.RFC1123, a3T)
	answer.A4TimeUTC, _ = time.Parse(time.RFC1123, a4T)
	answer.A5TimeUTC, _ = time.Parse(time.RFC1123, a5T)
	answer.A6TimeUTC, _ = time.Parse(time.RFC1123, a6T)
	answer.A7TimeUTC, _ = time.Parse(time.RFC1123, a7T)
	answer.A8TimeUTC, _ = time.Parse(time.RFC1123, a8T)
	answer.A9TimeUTC, _ = time.Parse(time.RFC1123, a9T)
	answer.A10TimeUTC, _ = time.Parse(time.RFC1123, a10T)
	answer.A11TimeUTC, _ = time.Parse(time.RFC1123, a11T)
	answer.A12TimeUTC, _ = time.Parse(time.RFC1123, a12T)
	answer.A13TimeUTC, _ = time.Parse(time.RFC1123, a13T)
	answer.A14TimeUTC, _ = time.Parse(time.RFC1123, a14T)
	answer.A15TimeUTC, _ = time.Parse(time.RFC1123, a15T)
	answer.A16TimeUTC, _ = time.Parse(time.RFC1123, a16T)
	answer.A17TimeUTC, _ = time.Parse(time.RFC1123, a17T)
	answer.A18TimeUTC, _ = time.Parse(time.RFC1123, a18T)

	DB.StoreNewAnswers(answer)

	resultSet := Scheme.Recommendation{}
	assessedGroups := Algorithm.ExecuteAnswers(answer)
	resultSet.ProductGroups = assessedGroups
	resultSet.CreateTimeUTC = time.Now().UTC()
	resultSet.Session = session
	DB.StoreRecommendation(resultSet)

	for _, entry := range resultSet.ProductGroups {
		fmt.Fprintf(response, "%s#%d\n", entry.Name, entry.Percent)
	}
}
