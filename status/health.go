package status


import (

	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("content-type", "text/plain")

	status := "Ok"

	w.Write([]byte(status))

}
