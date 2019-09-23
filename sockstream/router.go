package sockstream

import (
	"fmt"
	"log"
	"net/http"
)

// Router is a wrapper that routes endpoints to functions.
type Router struct {
	hub *Hub
}

// NewRouter returns a reference to a new instance of Router.
func NewRouter(hub *Hub) *Router {
	return &Router{
		hub: hub,
	}
}

// Run the Router.
func (router *Router) Run() {
	fmt.Println("starting server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(router.hub, w, r)
	})

	log.Fatal(
		http.ListenAndServe(":6969", nil),
	)
}
