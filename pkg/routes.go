package goskeleton

func (s *server) routes() {
	s.router.Path("/hello").Methods("GET").HandlerFunc(s.handleHello())
}
