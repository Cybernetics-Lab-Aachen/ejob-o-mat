package main

import (
	"strconv"

	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/ejob-o-mat/Website"
	"github.com/SommerEngineering/ejob-o-mat/XML"
)

//registerHandlers registers all request handlers for ejob-o-mat.
func registerHandlers() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers done.`)

	// Public Handlers:
	Handlers.AddPublicHandler(`/`, Website.HandlerRedirect)
	Handlers.AddPublicHandler(`/start`, Website.HandlerStart)
	for i := 0; i < len(XML.Questions); i++ { // Register handlers for all questions
		Handlers.AddPublicHandler("/question"+strconv.Itoa(i), Website.HandlerQuestion)
	}
	Handlers.AddPublicHandler(`/answer`, Website.HandlerAnswer)
	Handlers.AddPublicHandler(`/results`, Website.HandlerResults)
	Handlers.AddPublicHandler(`/impressum`, Website.HandlerImpressum)
	Handlers.AddPublicHandler(`/feedback`, Website.HandlerFeedback)
	Handlers.AddPublicHandler(`/writeFeedback`, Website.HandlerReceiveFeedback)
	Handlers.AddPublicHandler(`/ping`, Website.HandlerPing)

	// Admin Handlers:
	Handlers.AddAdminHandler(`/book`, Website.HandlerBook)
	Handlers.AddAdminHandler(`/Version`, Website.HandlerVersion)
}
