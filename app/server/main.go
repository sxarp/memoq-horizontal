package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func (s *server) listen(port int) {
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: s.router,
	}

	fmt.Println("server started at port: ", port)

	server.ListenAndServe()
}

func main() {

	srv := server{
		router: mux.NewRouter(),
	}

	srv.routes()

	srv.listen(8080)
}
