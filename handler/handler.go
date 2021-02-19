package handler

import (
	"encoding/json"
	"fmt"
	// "internal/syscall/windows"
	"net/http"

	model "github.com/rishi-org-stack/pz/model"
)

func Shout(w http.ResponseWriter, r *http.Request) {
	var database = &model.DB{}
	body := r.Body
	var inputuser = &model.User{}
	json.NewDecoder(body).Decode(inputuser)
	database = database.ConnectTodb()
	database.T = inputuser
	err := database.CreateaTable()
	if err != nil {
		fmt.Println(err)
	}
	inputuser.Save(database)
	println(inputuser.Email)
	// http.Redirect(w, r, "http://localhost:8081/#/list", 1)
}

func Listall(w http.ResponseWriter, r *http.Request){
	var database = &model.DB{}
	database = database.ConnectTodb()
	var inputuser = &model.User{}
	database.T = inputuser
	all := inputuser.GetallUser(database)
	json.NewEncoder(w).Encode(all)
}