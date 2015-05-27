package Website

import (
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
)

type Basis struct {
	Name    string
	Version string
	Lang    string
	Logo    string
	Session string
}

type PagePassword struct {
	Basis        Basis
	TextPassword string
}

type PageStart struct {
	Basis           Basis
	TextWelcome     string
	TextStartButton string
	TextProject     string
	TextExecuted    string
	TextPromoted    string
}

type PageFeedback struct {
	Basis            Basis
	TextFeedback     string
	TextYourFeedback string
	TextYourRating   string
	TextRatingLeft   string
	TextRatingRight  string
	TextSubmit       string
	SourceLocation   string
}

type PageImpressum struct {
	Basis              Basis
	TextPrefix4English string
}

type PageResults struct {
	Basis            Basis
	LangPos          int
	AmountCurrent    int
	AmountToggle     int
	Groups           []XML.ProductGroup
	Recommendation   Scheme.Recommendation
	TextHeader1      string
	TextHeader2      string
	TextHeader3      string
	TextHeaderPrefix string
	TextMatch        string
	TextGroup        string
	TextExamples     string
	TextRestart      string
	TextAllGroups    string
	TextOptionen     string
	TextResults      string
}

type PageQuestion struct {
	Basis              Basis
	NoQuestion         string
	PreNoQuestion      string
	NoQuestions        string
	Progress           int
	TextQuestion       string
	QuestionInfoText   string
	QuestionInfoHeader string
	QuestionInfoClose  string
	TextQuestionTopic  string
	TextQuestionBody   string
	TextButton1        string
	TextButton2        string
	TextButton3        string
	TextButton4        string
	TextButton5        string
	TextBackButton     string
	TextImportant      string
	Button1Status      string
	Button2Status      string
	Button3Status      string
	Button4Status      string
	Button5Status      string
	ButtonBackStatus   string
	ButtonInfoStatus   string
	Button1Data        string
	Button2Data        string
	Button3Data        string
	Button4Data        string
	Button5Data        string
}
