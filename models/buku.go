package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	u "myperpustakaan/utils"
)

type Buku struct {
	Id           int    `json:"id"`
	Judul        string `json:"judul"`
	Penulis      string `json:"penulis"`
	StokAsli     string `json:"stokAsli"`
	StokSekarang string `json:"stokSekarang"`
}

type StokBuku struct {
	Judul        string `json:"judul"`
	StokSekarang string `json:"stokSekarang"`
}

func GetAllBuku(a string) []*Buku {
	buku := make([]*Buku, 0)
	err := Connect().Table(a).Find(&buku).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buku
}

func GetStokBuku(a string) []*StokBuku {
	buku := []*StokBuku{}
	err := Connect().Table(a).Select([]string{"judul", "stok_sekarang"}).Find(&buku).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buku
}

func (buku *Buku) CreateBuku(a string) map[string]interface{} {
	err := Connect().Table(a).Create(&buku).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := u.Message(true, "success")
	resp["buku"] = buku

	return resp
}
