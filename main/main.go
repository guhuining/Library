package main

import (
	"library/config"
	controller "library/server"
	"net/http"
)

func main() {
	c := &config.Config{}
	serverConfig := c.GetServer()
	server := http.Server{
		Addr: serverConfig.Host + ":" + serverConfig.Port,
	}

	http.HandleFunc("/api/create_administrator", controller.CreateAdministrator)
	http.HandleFunc("/api/login_administrator", controller.LoginAdministrator)
	http.HandleFunc("/api/create_librarian", controller.CreateLibrarian)
	http.HandleFunc("/api/login_librarian", controller.LoginLibrarian)
	http.HandleFunc("/api/create_borrower", controller.CreateBorrower)
	http.HandleFunc("/api/login_borrower", controller.LoginBorrower)
	http.HandleFunc("/api/add_publication_type", controller.AddPublicationType)
	http.HandleFunc("/api/delete_publication_type", controller.DeletePublicationType)
	http.HandleFunc("/api/add_publication", controller.AddPublication)
	http.HandleFunc("/api/delete_publication", controller.DeletePublication)
	http.HandleFunc("/api/delete_librarian", controller.DeleteLibrarian)
	http.HandleFunc("/api/bind_card", controller.BindCard)
	http.HandleFunc("/api/delete_card", controller.DeleteCard)
	http.HandleFunc("/api/get_publication_by_name", controller.GetPublicationByName)
	http.HandleFunc("/api/borrow_publication", controller.BorrowPublication)

	server.ListenAndServe()
}
