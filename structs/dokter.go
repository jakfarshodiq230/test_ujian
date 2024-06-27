package structs

import (
	"github.com/jinzhu/gorm"
)

type Dokters struct {
	gorm.Model
	Nama          	string
	Spesialisasi  	string
	Nomor_Telepon 	string
}

type ParamsCreateDokter struct {
	ID       int    `json:"id"`
	Nama      		string `json:"nama" form:"nama" binding:"required"`
	Spesialisasi    string `json:"spesialisasi" form:"spesialisasi" binding:"required"`
	Nomor_Telepon 	string `json:"nomor_telepon" form:"nomor_telepon" binding:"required"`
	CreatedAt 		string
}

type ParamsUpdateDokter struct {
	ID      int    `json:"id" form:"id" binding:"required"`
	Nama      		string `json:"nama" form:"nama" `
	Spesialisasi    string `json:"spesialisasi" form:"spesialisasi"`
	Nomor_Telepon 	string `json:"nomor_telepon" form:"nomor_telepon"`
	UpdatedAt     	string
}

type ParamsGetListDokter struct {
	ID 	int    `json:"id" form:"id"`
	Nama   			string `json:"nama" form:"nama"`
	Limit  			*int   `json:"limit" form:"limit"`
	Offset 			*int   `json:"offset" form:"offset"`
}
type DataGetDokter struct {
	ID      int    `json:"id"`
	Nama      		string `json:"nama"`
	Spesialisasi    string `json:"spesialisasi"`
	Nomor_Telepon  	string `json:"nomor_telepon"`
	CreatedAt 		string `json:"created_at"`
}
type ParamsDokter struct {
	ID      int `json:"id" form:"id" binding:"required"`
	DeletedAt 		string
}

