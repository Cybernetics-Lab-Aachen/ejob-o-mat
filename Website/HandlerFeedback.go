package Website

import (
	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
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
		data.TextFeedback = `Hier haben Sie die Möglichkeit Ihr Feedback ans uns zu richten. Wir freuen uns über Ihre Meinung und bedanken uns im Voraus, dass Sie sich die Zeit dafür nehmen. Vielen Dank, dass Sie den ` + NAME_DE_PLAIN + ` benutzen.`
	} else {
		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.TextRatingLeft = `very good`
		data.TextRatingRight = `poor`
		data.TextYourFeedback = `Your Feedback`
		data.TextYourRating = `Your Rating`
		data.TextSubmit = `Submit`
		data.TextFeedback = `Here you have the possibility to provide feedback for us. We are very interested on your opinion about the ` + NAME_EN_PLAIN + `. Thank you very much for your support and using of the ` + NAME_EN_PLAIN + `.`
	}

	// Execute the template
	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`feedback`, response, data)
}
