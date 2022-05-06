package server

import "net/http"

func Routes() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/countries", (func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(w, r)
		case http.MethodPost:
			addCountries(w, r)
		case http.MethodDelete:
			deleteCountry(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
}
