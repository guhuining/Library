package main

import (
	"library/config"
	"net/http"
)

func main() {
	c := &config.Config{}
	serverConfig := c.GetServer()
	server := http.Server{
		Addr: serverConfig.Host + ":" + serverConfig.Port,
	}

	server.ListenAndServe()
}
