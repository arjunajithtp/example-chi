package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

// Data holds all data from request
type Data struct {
	Param  *Param  `json:"param,omitempty"`
	Header *Header `json:"header,omitempty"`
	Body   *Body   `json:"body,omitempty"`
}

// Param holds url params
type Param struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Header holds request Header data
type Header struct {
	Token string `json:"token,omitempty"`
}

// Body holds request body data
type Body struct {
	Message string `json:"message,omitempty"`
}

// GetID is our example for handler with GET method
func GetID(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Param: &Param{
			ID:   chi.URLParam(r, "id"),
			Name: r.URL.Query()["name"][0],
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("error while marshaling: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// GetAll is our example for handler with POST method
func GetAll(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Param: &Param{
			ID:   chi.URLParam(r, "id"),
			Name: r.URL.Query()["name"][0],
		},
		Header: &Header{
			Token: r.Header.Get("token"),
		},
		Body: &Body{
			Message: r.FormValue("message"),
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("error while marshaling: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
