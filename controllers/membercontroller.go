package controllers

import (
	"encoding/json"
	"myperpustakaan/models"
	u "myperpustakaan/utils"
	"net/http"
)

var GetMembers = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllMember("member")
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateMember = func(w http.ResponseWriter, r *http.Request) {
	member := &models.Member{}

	err := json.NewDecoder(r.Body).Decode(member)
	if err != nil {
		u.Respond(w, u.Message(false, "error while decoding"))
		return
	}
	resp := member.CreateMember("member")
	u.Respond(w, resp)
}
