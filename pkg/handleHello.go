package goskeleton

import (
	"fmt"
	"net/http"
)

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		data := fmt.Sprintf("Hello %s.", q.Get("name"))
		s.respond(w, r, 200, data)
	}
}
