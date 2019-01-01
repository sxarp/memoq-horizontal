package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	repo   Repo
	logger *log.Logger
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
		repo:   NewDStore("prj", 5000),
		logger: log.New(os.Stdout, "logger: ", log.Llongfile|log.Lmicroseconds),
	}

	srv.routes()

	srv.listen(8080)
}
