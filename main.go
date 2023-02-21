package main

import (
	"goTask/router"
)

func main() {

	// http.HandleFunc("/", sayhi)

	// err := http.ListenAndServe(":9090", nil)
	// if err != nil {
	// 	log.Fatal("ListenAnndAerve:", err)
	// }
	server := router.SetRouter()
	server.Run(":9090")
}
