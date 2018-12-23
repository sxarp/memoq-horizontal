package main

func (s *server) routes() {
	s.router.HandleFunc("/test/index", s.handleIndex())
}
