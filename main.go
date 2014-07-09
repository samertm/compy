package main

import (
	"fmt"
	"github.com/samertm/compy/server"
)

func main() {
	ip := "localhost:4444"
	fmt.Println("Listening on", ip)
	server.ListenAndServe(ip)
}
