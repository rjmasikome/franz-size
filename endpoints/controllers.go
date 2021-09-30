package endpoints

import "net/http"

func ok() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseJSON(w, "ok")
	}
}
