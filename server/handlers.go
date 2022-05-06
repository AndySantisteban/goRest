package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Restful API")
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)

}

func addCountries(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	countries = append(countries, country)
	fmt.Fprintf(w, "Country added successfully")
}

func deleteCountry(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	for i, c := range countries {
		if c.ID == country.ID {
			countries = append(countries[:i], countries[i+1:]...)
			fmt.Fprintf(w, "Country deleted successfully")
			return
		}
	}

	fmt.Fprintf(w, "Country not found")

}
