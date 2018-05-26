package s3rver

import (
	"net/http"

	"github.com/Cloraxlan/master-s-sock"
)

type Server struct {
	handle func()
	port   int
	path   string
	status bool
	hub    *sock.Hub
}

func NewServer(handleFunction func(), port int, path string) *Server {
	server := &Server{handle: handleFunction, port: port, path: path}
	go startServer(server)
	return server
}

func startServer(server *Server) {
	server.hub = sock.NewHub()
	http.HandleFunc(server.path, func(w http.ResponseWriter, r *http.Request) {
		sock.ServeWs(server.hub, w, r)
	})
	go server.hub.Run()
	go server.handle()
}
