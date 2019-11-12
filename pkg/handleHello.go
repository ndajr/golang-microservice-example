package app

import (
	"fmt"
	"net/http"
)

func greetings(name string) string {
	return fmt.Sprintf("Hello %s.", name)
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		name := q.Get("name")
		response := greetings(name)
		s.respond(w, r, 200, response)
	}
}
