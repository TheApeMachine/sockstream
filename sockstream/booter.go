package sockstream

import "gocv.io/x/gocv"

// Booter is a wrapper around starting all the services for SockStream.
type Booter struct {
	hub     *Hub
	router  *Router
	capture *Capture
}

// NewBooter returns a reference to a new instance of Booter.
func NewBooter() *Booter {
	capCh := make(chan gocv.Mat)
	hub := NewHub(capCh)

	return &Booter{
		hub:     hub,
		router:  NewRouter(hub),
		capture: NewCapture(hub),
	}
}

// Run the boot process.
func (booter *Booter) Run() {
	go booter.hub.Run()
	go booter.capture.Run()
	booter.router.Run()
}
