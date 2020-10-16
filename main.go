package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"myperpustakaan/controllers"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/perpustakaan/get/buku", controllers.GetBukus).Methods("GET")
	router.HandleFunc("/perpustakaan/get/stokbuku", controllers.GetStokBukus).Methods("GET")
	router.HandleFunc("/perpustakaan/save/buku", controllers.CreateBuku).Methods("POST")
	router.HandleFunc("/perpustakaan/get/member", controllers.GetMembers).Methods("GET")
	router.HandleFunc("/perpustakaan/save/member", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/perpustakaan/get/laporan/pinjam", controllers.GetLaporanPinjam).Methods("GET")
	router.HandleFunc("/perpustakaan/transaksi/pinjam", controllers.TransaksiPeminjaman).Methods("POST")
	router.HandleFunc("/perpustakaan/get/laporan/kembali", controllers.GetLaporanKembali).Methods("GET")

	fmt.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
