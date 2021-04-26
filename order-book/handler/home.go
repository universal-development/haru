package handler

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Order book home handler")
}
