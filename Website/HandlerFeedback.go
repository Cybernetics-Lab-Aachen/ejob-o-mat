package Website

import (
	"net/http"
	"strings"

	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
)

//HandlerFeedback displays the feedback form
func HandlerFeedback(response http.ResponseWriter, request *http.Request) {
	lang := Tools.GetRequestLanguage(request)[0]

	// Prepare data for html template
	data := PageFeedback{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = request.FormValue(`session`) // Not validation session here, just passign it on for storing
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")
	data.SourceLocation = request.Referer()

	//Prepare localized strings
	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.TextRatingLeft = `sehr gut`
		data.TextRatingRight = `unbrauchbar`
		data.TextYourFeedback = `Ihr Feedback`
		data.TextYourRating = `Ihre Bewertung`
		data.TextSubmit = `Absenden`
		data.TextFeedbackParagraph1 = `Vielen Dank, dass Sie den ejob-o-mat benutzen.`
		data.TextFeedbackParagraph2 = `Hier haben Sie die Möglichkeit, Ihr Feedback ans uns zu richten. Wir freuen uns über Ihre Meinung und bedanken uns im Voraus, dass Sie sich die Zeit dafür nehmen.`
	} else {
		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.TextRatingLeft = `very good`
		data.TextRatingRight = `poor`
		data.TextYourFeedback = `Your Feedback`
		data.TextYourRating = `Your Rating`
		data.TextSubmit = `Submit`
		data.TextFeedbackParagraph1 = `Here you have the possibility to provide feedback for us. `
		data.TextFeedbackParagraph2 = `We are very interested on your opinion about the ` + NAME_EN_PLAIN + `. Thank you very much for your support and using of the ` + NAME_EN_PLAIN + `.`
	}

	// Execute the template
	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`feedback`, response, data)
}

// PageFeedback contains data for the feedback template.
type PageFeedback struct {
	Basis                  Basis
	TextFeedbackParagraph1 string
	TextFeedbackParagraph2 string
	TextYourFeedback       string
	TextYourRating         string
	TextRatingLeft         string
	TextRatingRight        string
	TextSubmit             string
	SourceLocation         string
}
