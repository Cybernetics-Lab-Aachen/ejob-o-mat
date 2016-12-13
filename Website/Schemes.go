package Website

import (
	"github.com/SommerEngineering/Re4EEE/DB/Scheme"
	"github.com/SommerEngineering/Re4EEE/XML"
)

type Basis struct {
	Name                  string
	Version               string
	Lang                  string
	Logo                  string
	Session               string
	SiteVerificationToken string
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
	Basis          Basis
	LangPos        int
	AmountCurrent  int
	AmountToggle   int
	Groups         []XML.ProductGroup
	Questions      []XML.QuestionGroup
	Recommendation Scheme.Recommendation
	TextAllGroups  string
	Strings        XML.ResultStrings
	Answers        Scheme.Answers
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
	TextBackButton     string
	TextImportant      string
	ButtonBackStatus   string
	ButtonInfoStatus   string
	Buttons            []PageQuestionButton
}

type PageQuestionButton struct {
	Text string
	Data string
}
