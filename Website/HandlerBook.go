package Website

import (
	"fmt"
	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
	"github.com/SommerEngineering/Ocean/MimeTypes"
	"github.com/SommerEngineering/Ocean/StaticFiles"
	"github.com/SommerEngineering/Ocean/Tools"
	"net/http"
)

//HandlerBook sends Data.xml file.
func HandlerBook(response http.ResponseWriter, request *http.Request) {

	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameREQUEST, `Someone has requested the whole data source.`, request.RemoteAddr)

	Tools.SendChosenLanguage(response, langBook)
	MimeTypes.Write2HTTP(response, MimeTypes.TypeXML)
	fmt.Fprintln(response, string(StaticFiles.FindAndReadFile(`Data.xml`)))
}
