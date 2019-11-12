package server

func (s *Server) routes() {
	s.router.Path("/version").Methods("GET").HandlerFunc(s.handleVersion())
	s.router.Path("/hello").Methods("GET").HandlerFunc(s.handleHello())
}
