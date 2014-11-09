package Website

import (
	"fmt"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

const (
	NAME            string = `Re4EEE`
	VERSION         string = `1.0b1`
	BUTTON_HIDDEN   string = ` buttonhidden`
	BUTTON_SHOW     string = ``
	TOTAL_QUESTIONS int    = 28
)

var (
	betaPassword   string    = `866a2058-11bb-47c0-a0a2-4d6e179417ab`
	senderName     LM.Sender = `Website`
	totalQuestions string    = fmt.Sprintf("%d", TOTAL_QUESTIONS)
)
