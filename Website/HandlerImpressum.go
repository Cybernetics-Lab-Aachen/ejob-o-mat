package Website

import (
	"net/http"
	"strings"

	"github.com/SommerEngineering/Ocean/ConfigurationDB"
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
)

//HandlerImpressum displays the impressum.
func HandlerImpressum(response http.ResponseWriter, request *http.Request) {
	lang := Tools.GetRequestLanguage(request)[0]
	data := PageImpressum{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = request.FormValue(`session`) // No session validation neccessary, as session is not used here except for passing it on.
	data.Basis.SiteVerificationToken = ConfigurationDB.Read("SiteVerificationToken")

	if strings.Contains(lang.Language, `de`) {
		data.Basis.Name = NAME_DE
		data.Basis.Logo = LOGO_DE
		data.TextPrefix4English = `parhidden`
	} else {
		data.Basis.Name = NAME_EN
		data.Basis.Logo = LOGO_UK
		data.TextPrefix4English = ``
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`impressum`, response, data)
}

// PageImpressum contains data for the impressum template.
type PageImpressum struct {
	Basis              Basis
	TextPrefix4English string
}
