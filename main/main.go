package main

import (
	"library/tools"
	"net/http"
)

func main() {
	serverConfig := tools.ServerConfig()
	server := http.Server{
		Addr: serverConfig.Host + ":" + serverConfig.Port,
	}

	server.ListenAndServe()
}
