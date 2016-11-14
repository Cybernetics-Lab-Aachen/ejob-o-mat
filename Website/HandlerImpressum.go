package Website

import (
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

//HandlerImpressum displays the impressum.
func HandlerImpressum(response http.ResponseWriter, request *http.Request) {
	lang := Tools.GetRequestLanguage(request)[0]
	data := PageImpressum{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language
	data.Basis.Session = request.FormValue(`session`) // No session validation neccessary, as session is not used here except for passing it on.

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
