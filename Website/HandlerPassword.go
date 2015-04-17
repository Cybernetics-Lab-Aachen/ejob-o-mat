package Website

import (
	"github.com/SommerEngineering/Ocean/Templates"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
	"strings"
)

func HandlerPassword(response http.ResponseWriter, request *http.Request) {

	lang := Tools.GetRequestLanguage(request)[0]
	data := PagePassword{}
	data.Basis.Version = VERSION
	data.Basis.Lang = lang.Language

	if strings.Contains(lang.Language, `de`) {
		data.TextPassword = `Passwort`
		data.Basis.Name = NAME_DE
	} else {
		data.TextPassword = `Password`
		data.Basis.Name = NAME_EN
	}

	Tools.SendChosenLanguage(response, lang)
	Templates.ProcessHTML(`password`, response, data)
}
