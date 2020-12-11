package main

import (
	"library/config"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: config.IP + ":" + config.PORT,
	}

	server.ListenAndServe()
}
