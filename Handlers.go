package main

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/Re4EEE/Website"
	"net/http"
)

func registerHandlers() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers done.`)

	Handlers.AddPublicHandler(`/`, Website.HandlerPassword)
	Handlers.AddPublicHandler(`/start`, Website.HandlerStart)
	Handlers.AddPublicHandler(`/question1`, Website.HandlerQuestion1)
	Handlers.AddPublicHandler(`/question2`, Website.HandlerQuestion2)
	Handlers.AddPublicHandler(`/question3`, Website.HandlerQuestion3)
	Handlers.AddPublicHandler(`/question4`, Website.HandlerQuestion4)
	Handlers.AddPublicHandler(`/question5`, Website.HandlerQuestion5)
	Handlers.AddPublicHandler(`/question6`, Website.HandlerQuestion6)
	Handlers.AddPublicHandler(`/question7`, Website.HandlerQuestion7)
	Handlers.AddPublicHandler(`/question8`, Website.HandlerQuestion8)
	Handlers.AddPublicHandler(`/question9`, Website.HandlerQuestion9)
	Handlers.AddPublicHandler(`/question10`, Website.HandlerQuestion10)
	Handlers.AddPublicHandler(`/question11`, Website.HandlerQuestion11)
	Handlers.AddPublicHandler(`/question12`, Website.HandlerQuestion12)
	Handlers.AddPublicHandler(`/question13`, Website.HandlerQuestion13)
	Handlers.AddPublicHandler(`/question14`, Website.HandlerQuestion14)
	Handlers.AddPublicHandler(`/question15`, Website.HandlerQuestion15)
	Handlers.AddPublicHandler(`/question16`, Website.HandlerQuestion16)
	Handlers.AddPublicHandler(`/question17`, Website.HandlerQuestion17)
	Handlers.AddPublicHandler(`/question18`, Website.HandlerQuestion18)
	Handlers.AddPublicHandler(`/question19`, Website.HandlerQuestion19)
	Handlers.AddPublicHandler(`/question20`, Website.HandlerQuestion20)
	Handlers.AddPublicHandler(`/question21`, Website.HandlerQuestion21)
	Handlers.AddPublicHandler(`/question22`, Website.HandlerQuestion22)
	Handlers.AddPublicHandler(`/question23`, Website.HandlerQuestion23)
	Handlers.AddPublicHandler(`/question24`, Website.HandlerQuestion24)
	Handlers.AddPublicHandler(`/question25`, Website.HandlerQuestion25)
	Handlers.AddPublicHandler(`/question26`, Website.HandlerQuestion26)
	Handlers.AddPublicHandler(`/question27`, Website.HandlerQuestion27)
	Handlers.AddPublicHandler(`/answer`, Website.HandlerAnswer)
}

func testHandler(response http.ResponseWriter, request *http.Request) {
	langs := Tools.GetRequestLanguage(request)
	text := fmt.Sprintf("Count langs=%d     ", len(langs))
	text += fmt.Sprintf("Best lang=%s with %f     ", langs[0].Language, langs[0].Factor)
	text += fmt.Sprintf("Original header=%s    ", request.Header.Get(`Accept-Language`))

	Tools.SendChosenLanguage(response, langs[0])
	response.Write([]byte(text))
}
