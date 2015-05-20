package Website

import (
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

func HandlerImpressum(response http.ResponseWriter, request *http.Request) {
	readSession := request.FormValue(`session`)
	lang := Tools.GetRequestLanguage(request)[0]
	data := PageImpressum{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language

	if readSession != `` {
		data.Basis.Session = readSession
	} else {
		data.Basis.Session = Tools.RandomGUID()
	}

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
