package status


import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request)  {
	w.Header().Set("content-type", "text/plain")

	status := "OK"

	_, _ = fmt.Fprintf(w, "Health: %v", status)

}
