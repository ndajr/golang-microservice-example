package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fasterness/cors"
	"github.com/gorilla/mux"
)

// New : server creation
func New() http.Handler {
	router := mux.NewRouter()
	s := &Server{
		router: router,
	}
	s.routes()
	return cors.New(s)
}

// Server : initializing all server dependencies
type Server struct {
	router *mux.Router
}

func (s Server) handle(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s Server) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("encode response: %s", err)
	}
}

func (s Server) responderr(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.WriteHeader(status)
	var data struct {
		Error string `json:"error"`
	}
	if err != nil {
		data.Error = err.Error()
	} else {
		data.Error = "Something went wrong"
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("encode response: %s", err)
	}
}
