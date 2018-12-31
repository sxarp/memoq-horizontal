package main

func (s *server) routes() {
	s.router.HandleFunc("/test/index", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/simple", s.simpleCreate()).Methods("POST")
	s.router.HandleFunc("/simple/{id:[0-9]+}", s.simpleShow()).Methods("GET")
	s.router.HandleFunc("/simple/{id:[0-9]+}", s.simpleDestroy()).Methods("DELETE")
}
