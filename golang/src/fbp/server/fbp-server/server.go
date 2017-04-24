package fbpserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//https://flowbased.github.io/fbp-protocol/

/*
var ws = new WebSocket("ws://localhost:8080/websocket");
ws.addEventListener("message", function(e) {console.log(e);});
ws.send(`
{
  "sub-protocol": "runtime",
  "topic": "getruntime",
  "payload": {
    "secret": "SECRET"
  }
}`)
*/
type Server interface {
	Start()
}

type server struct {
	router *mux.Router
}

func New() Server {
	return &server{router: newRouter()}
}

func (svr *server) Start() {
	log.Fatal(http.ListenAndServe(":8080", svr.router))
}
