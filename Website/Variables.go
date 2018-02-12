package Website

import (
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Tools"
)

const (
	NAME_DE        string = `elearn-o-mat.net`
	NAME_EN        string = `elearning-finder.net`
	NAME_DE_PLAIN  string = `elearn-o-mat`
	NAME_EN_PLAIN  string = `elearning-finder`
	LOGO_DE        string = `logoDE`
	LOGO_UK        string = `logoUK`
	VERSION        string = `1.1`
	BUTTON_HIDDEN  string = ` buttonhidden`
	BUTTON_SHOW    string = ``
	PROGRESS_DONE  string = ` progressitemdone`
	PROGRESS_NDONE string = ``
	LANG_DE        int    = 0
	LANG_EN        int    = 1
)

var (
	betaPassword           = `866a2058-11bb-47c0-a0a2-4d6e179417ab`
	senderName   LM.Sender = `Website`
	langBook               = Tools.Language{Factor: 1.0, Language: `de`}
)
