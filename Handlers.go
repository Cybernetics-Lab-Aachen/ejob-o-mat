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

// /question?n=1
// /answer?n=1

func registerHandlers() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers done.`)

	Handlers.AddPublicHandler(`/`, Website.HandlerPassword)
	Handlers.AddPublicHandler(`/start`, Website.HandlerStart)
}

func testHandler(response http.ResponseWriter, request *http.Request) {
	langs := Tools.GetRequestLanguage(request)
	text := fmt.Sprintf("Count langs=%d     ", len(langs))
	text += fmt.Sprintf("Best lang=%s with %f     ", langs[0].Language, langs[0].Factor)
	text += fmt.Sprintf("Original header=%s    ", request.Header.Get(`Accept-Language`))

	Tools.SendChosenLanguage(response, langs[0])
	response.Write([]byte(text))
}
