package server

import (
	"net/http"
)

var countries = []*Country{
	{ID: "1", Name: "United States", Code: "US", Language: "English"},
	{ID: "2", Name: "Spain", Code: "ES", Language: "Spanish"},
	{ID: "3", Name: "France", Code: "FR", Language: "French"},
	{ID: "4", Name: "Germany", Code: "DE", Language: "German"},
	{ID: "5", Name: "Italy", Code: "IT", Language: "Italian"},
}

func Server(addr string) *http.Server {
	Routes()
	return &http.Server{
		Addr: addr,
	}
}
