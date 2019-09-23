package main

import "github.com/theapemachine/sockstream/sockstream"

func main() {
	booter := sockstream.NewBooter()
	booter.Run()
}
