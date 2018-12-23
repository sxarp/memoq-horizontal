package main

func (s *server) routes() {
	s.router.HandleFunc("/test/index", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/simple/create", s.simpleCreate()).Methods("POST")
}
