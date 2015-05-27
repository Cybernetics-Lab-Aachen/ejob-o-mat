package Website

import (
	"fmt"
	"net/http"
)

func HandlerPing(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "pong")
}
