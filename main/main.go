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
	http.HandleFunc("/api/is_out_of_time", controller.IsOutOfTime)
	http.HandleFunc("/api/return_publication", controller.ReturnPublication)
	http.HandleFunc("/api/get_borrow_item", controller.GetBorrowItem)
	http.HandleFunc("/api/get_publication_type", controller.GetPublicationType)
	http.HandleFunc("/api/order_publication", controller.OrderPublication)
	http.HandleFunc("/api/cancel_order_item", controller.CancelOrderItem)
	http.HandleFunc("/api/borrower_get_order_item", controller.BorrowerGetOrderItem)
	http.HandleFunc("/api/librarian_get_order_item", controller.LibrarianGetOrderItem)
	http.HandleFunc("/api/is_login", controller.IsLogin)
	http.HandleFunc("/api/logout", controller.Logout)
	http.HandleFunc("/api/borrower_get_borrowed_item", controller.BorrowerGetBorrowedPublication)
	http.HandleFunc("/api/get_borrower_type", controller.GetBorrowerType)

	server.ListenAndServe()
}
