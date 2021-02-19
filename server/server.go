package server

import (
	// "encoding/json"
	// "net/http"
	hnd "github.com/rishi-org-stack/pz/handler"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", hnd.Shout).Methods("GET","POST")
	r.HandleFunc("/all", hnd.Listall).Methods("GET","POST")
	return r
}
