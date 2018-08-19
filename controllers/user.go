package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RickyJian/gorouterpractice/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	// JSON解碼 body
	json.NewDecoder(r.Body).Decode(&u)

	u.Name = "Ricky"

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (uc *UserController) Remove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 刪除 邏輯
	w.WriteHeader(200)
}
