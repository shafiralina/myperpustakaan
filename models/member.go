package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	u "myperpustakaan/utils"
)

type Member struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Usia   string `json:"usia"`
	Status string `json:"status"`
}

func GetAllMember(a string) []*Member {
	member := make([]*Member, 0)
	err := Connect().Table(a).Find(&member).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return member
}

func (member *Member) CreateMember(a string) map[string]interface{} {
	err := Connect().Table(a).Create(&member).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := u.Message(true, "success")
	resp["member"] = member

	return resp
}
