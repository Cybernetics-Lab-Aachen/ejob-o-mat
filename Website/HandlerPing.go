package Website

import (
	"fmt"
	"net/http"
)

//HandlerQuestion displays "pong".
func HandlerPing(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "pong")
}
