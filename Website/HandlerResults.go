package Website

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"github.com/SommerEngineering/ejob-o-mat/Algorithm"
	"github.com/SommerEngineering/ejob-o-mat/DB"
	"github.com/SommerEngineering/ejob-o-mat/DB/Scheme"
	"github.com/SommerEngineering/ejob-o-mat/XML"
)

// HandlerResults displays the results of the questionnaire
func HandlerResults(response http.ResponseWriter, request *http.Request) {
	session := request.FormValue(`session`)
	amountText := request.FormValue(`amount`)
	lang := Tools.GetRequestLanguage(request)[0]
	answers, loadAnswersError := DB.LoadAnswers(session)
	groups := XML.GetData()

	// Check if session exists, otherwise redirect to start page
	if loadAnswersError {
		http.Redirect(response, request, `/start`, 302)
		return
	}

	amountValue, errConv := strconv.Atoi(amountText)
	if errConv != nil { // Validate input
		Log.LogFull(senderName, LM.CategoryAPP, LM.LevelERROR, LM.SeverityMiddle, LM.ImpactNone, LM.MessageNameREQUEST, `Cannot read the amount value!`, amountText, errConv.Error())
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// Load/calculate recommendations
	resultSet, loaRecommendationError := DB.LoadRecommendation(session)
	if loaRecommendationError { // Calculate recommendation
		resultSet.ProductGroups = Algorithm.ExecuteAnswers(answers.Answers)
		resultSet.CreateTimeUTC = time.Now().UTC()
		resultSet.Session = session
		resultSet.SchemeVersion = Scheme.CurrentVersion
		DB.StoreRecommendation(resultSet)
	} else {
		// Check for old session. Can't work with outdated data.
		if resultSet.SchemeVersion < Scheme.CurrentVersion {
			http.Redirect(response, request, `/start`, 302)
			return
		}
	}

	// Reduce number of shown product groups, if requested
	if amountValue > len(resultSet.ProductGroups) {
		amountValue = len(resultSet.ProductGroups)
	}
	if amountValue >= 1 {
		resultSet.ProductGroups = resultSet.ProductGroups[0:amountValue]
	}

	// Prepare localized strings
	data := PageResults{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = session
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")
	data.Groups = make(map[string]XML.ProductGroup)
	for _, productGroup := range groups.ProductsCollection.Products {
		data.Groups[productGroup.InternalName] = productGroup
	}
	data.Questions = groups.QuestionsCollection.Questions
	data.Recommendation = resultSet
	data.AmountCurrent = amountValue
	data.Strings = groups.ResultStrings
	data.Answers = answers

	if strings.Contains(lang.Language, `de`) {

		if amountValue > 0 {
			data.TextAllGroups = `Alle Gruppen anzeigen`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Top 6 anzeigen`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.LangPos = LANG_DE
	} else {

		if amountValue > 0 {
			data.TextAllGroups = `Show all groups`
			data.AmountToggle = -1
		} else {
			data.TextAllGroups = `Show top 6`
			data.AmountToggle = 6
		}

		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.LangPos = LANG_EN
	}

	// Prepare localized header string (using reseults)
	// Choose header string using average percentage on recommendations
	// Calc average percentage
	sumPercent := 0
	for _, productGroup := range resultSet.ProductGroups {
		sumPercent += productGroup.Percent
	}
	averagePercent := sumPercent / len(resultSet.ProductGroups)
	// Set header string
	if strings.Contains(lang.Language, `de`) {
		if averagePercent == 0 {
			data.HeaderText = `Auch wenn Sie Ihre Stärken und Potentiale aktuell nicht eindeutig identifizieren können, finden Sie sicher passende 
			Empfehlungen in unseren Job-Shadowing-Angeboten (link).`
		} else if averagePercent == 100 {
			data.HeaderText = `Ihre Stärken und Potentiale sind sehr vielfältig ausgeprägt. Schauen Sie sich gerne die Empfehlungen aller Typen, 
			die sich im Drop-Down-Menü aufrufen lassen, an und wählen Sie ein für Sie passendes Job-Shadowing-Angebot aus.`
		} else {
			data.HeaderText = `Ihre Stärken und Potentiale liegen entsprechend der prozentualen Aufteilung in den nachfolgend aufgeführten Empfehlungen. 
			Schauen Sie sich aber auch gerne die anderen Angebote an.`
		}
	} else {
		// Create english header text here, if needed
	}

	// Finally, execute the template
	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`results`, response, data)
}

// PageResults contains data for the results template.
type PageResults struct {
	Basis          Basis
	LangPos        int
	AmountCurrent  int
	AmountToggle   int
	Groups         map[string]XML.ProductGroup
	Questions      []XML.QuestionGroup
	Recommendation Scheme.Recommendation
	TextAllGroups  string
	Strings        XML.ResultStrings
	Answers        Scheme.Survey
	HeaderText     string
}

// GetProgressState returns the css class representing the progress.
func (data PageResults) GetProgressState(influence int8) string {
	if influence > 0 {
		return ` progressitemdone`
	} else if influence < 0 {
		return ` progressitemundone`
	} else {
		return ``
	}
}

// GetGroupName returns the localized name of a product group by its internal name.
func (data PageResults) GetGroupName(internalName string) string {
	return data.Groups[internalName].GroupNames[data.LangPos].Text
}

// GetGroupId returns an html identifer fo a product group
func (data PageResults) GetGroupId(internalName string) string {
	// Simply remove whitespaces
	return strings.Replace(internalName, " ", "", -1)
}

// Lang returns the localized string using the language id.
func (data PageResults) Lang(strings []XML.String) string {
	return strings[data.LangPos].Text
}

// FormatAnswer returns localized string for a given answer by it's internal representation.
func (data PageResults) FormatAnswer(question XML.QuestionGroup) string {
	answer := data.Answers.Answers[question.InternalName].Data
	return question.ButtonsByData[answer].Texts[data.LangPos].Text
}

// GetMaxPercentGroup returns the name of the group with maximal percentage
func (data PageResults) GetMaxPercentGroupId() string {
	maxGroup := data.Recommendation.ProductGroups[0]
	for _, group := range data.Recommendation.ProductGroups {
		if maxGroup.Percent < group.Percent {
			maxGroup = group
		}
	}
	return data.GetGroupId(maxGroup.Name)
}
