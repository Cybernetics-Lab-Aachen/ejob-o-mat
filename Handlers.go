package main

import (
	"github.com/SommerEngineering/Ocean/Handlers"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/ejob-o-mat/Website"
)

//registerHandlers registers all request handlers for ejob-o-mat.
func registerHandlers() {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers.`)
	defer Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameSTARTUP, `Register now all app handlers done.`)

	// Public Handlers:
	Handlers.AddPublicHandler(`/`, Website.HandlerRedirect)
	Handlers.AddPublicHandler(`/start`, Website.HandlerStart)
	Handlers.AddPublicHandler(`/question1`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question2`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question3`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question4`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question5`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question6`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question7`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question8`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question9`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question10`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question11`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question12`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question13`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question14`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question15`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question16`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question17`, Website.HandlerQuestion)
	Handlers.AddPublicHandler(`/question18`, Website.HandlerQuestion)
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
