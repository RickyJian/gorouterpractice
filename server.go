package main

import (
	"net/http"

	"github.com/RickyJian/gorouterpractice/controllers"
	"github.com/julienschmidt/httprouter"
)

var userController = controllers.NewUserController()

func main() {
	// 開啟 HTTPRouter
	mux := httprouter.New()
	// GET 方法
	mux.GET("/user/:name", userController.GetUserInfo)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
