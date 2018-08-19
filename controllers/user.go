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

func (uc *UserController) GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
