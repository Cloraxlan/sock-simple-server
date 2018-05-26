package s3rver

import (
	"net/http"

	"github.com/Cloraxlan/master-s-sock"
)

type handle func()
type Server struct {
	port   int
	path   string
	status bool
	hub    *sock.Hub
}

func NewServer(port int, path string) *Server {
	server := &Server{port: port, path: path}

	startServer(server)
	return server
}

func startServer(server *Server) {
	server.hub = sock.NewHub()
	http.HandleFunc(server.path, func(w http.ResponseWriter, r *http.Request) {
		sock.ServeWs(server.hub, w, r)
	})
	go server.hub.Run()

}
