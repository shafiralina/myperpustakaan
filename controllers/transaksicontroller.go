package controllers

import (
	"encoding/json"
	"myperpustakaan/models"
	u "myperpustakaan/utils"
	"net/http"
)

var GetLaporanPinjam = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetLaporanPinjam("historypinjam")
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetLaporanKembali = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetLaporanKembali("historypengembalian")
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var TransaksiPeminjaman = func(w http.ResponseWriter, r *http.Request) {
	pinjam := &models.Pinjam{}

	err := json.NewDecoder(r.Body).Decode(pinjam)
	if err != nil {
		u.Respond(w, u.Message(false, "error while decoding"))
		return
	}
	resp := pinjam.TransaksiPeminjaman("historypinjam")
	u.Respond(w, resp)
}
