package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RickyJian/gorouterpractice/models"
	"github.com/julienschmidt/httprouter"
)

func get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   p.ByName("name"),
		Gender: "male",
		Age:    18,
	}
	// JSON 編碼
	userJSON, _ := json.Marshal(u)
	// 設定回傳 Header
	w.Header().Set("Content-Type", "application/json")
	// 設定 HTTP Status
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", userJSON)
}

func main() {
	// 開啟 HTTPRouter
	mux := httprouter.New()
	// GET 方法
	mux.GET("/hello/:name", get)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
