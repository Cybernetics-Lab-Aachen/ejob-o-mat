package Website

import (
	"fmt"
	"net/http"

	"github.com/SommerEngineering/Ocean/Log"
	LM "github.com/SommerEngineering/Ocean/Log/Meta"
)

//HandlerVersion displays the current website version.
func HandlerVersion(response http.ResponseWriter, request *http.Request) {
	Log.LogShort(senderName, LM.CategoryAPP, LM.LevelINFO, LM.MessageNameREQUEST, `Someone has requested the version.`, request.RemoteAddr)
	fmt.Fprintf(response, "%s\n", VERSION)
}
