package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/nu7hatch/gouuid"
	u "myperpustakaan/utils"
	"time"
)

type Pinjam struct {
	Id            string `json:"id"`
	MemberId      string `json:"idMember"`
	BukuId        string `json:"idBuku"`
	TanggalPinjam string `json:"tanggalPinjam"`
}

type Kembali struct {
	Id             string `json:"id"`
	MemberId       string `json:"memberId"`
	BukuId         string `json:"bukuId"`
	TanggalKembali string `json:"tanggalKembali"`
}

func GetLaporanPinjam(a string) []*Pinjam {
	pinjam := make([]*Pinjam, 0)
	err := Connect().Table(a).Find(&pinjam).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pinjam
}

func GetLaporanKembali(a string) []*Kembali {
	kembali := make([]*Kembali, 0)
	err := Connect().Table(a).Find(&kembali).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return kembali
}

func (pinjam *Pinjam) TransaksiPeminjaman(a string) map[string]interface{} {
	t := time.Now()
	id, _ := uuid.NewV4()
	data := Pinjam{Id: id.String(), MemberId: pinjam.MemberId, BukuId: pinjam.BukuId, TanggalPinjam: t.Format("2006-01-02")}

	err := Connect().Table(a).Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp := u.Message(true, "success")
	resp["pinjam"] = pinjam

	return resp
}
