package Website

import (
	"fmt"
	"net/http"
)

// HandlerPing displays "pong".
func HandlerPing(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "pong")
}
