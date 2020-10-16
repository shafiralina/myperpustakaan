package controllers

import (
	"encoding/json"
	"myperpustakaan/models"
	u "myperpustakaan/utils"
	"net/http"
)

var GetBukus = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllBuku("buku")
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetStokBukus = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetStokBuku("buku")
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateBuku = func(w http.ResponseWriter, r *http.Request) {
	buku := &models.Buku{}

	err := json.NewDecoder(r.Body).Decode(buku)
	if err != nil {
		u.Respond(w, u.Message(false, "error while decoding"))
		return
	}

	resp := buku.CreateBuku("buku")
	u.Respond(w, resp)
}
