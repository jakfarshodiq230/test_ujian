package structs

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Pasiens struct {
	gorm.Model
	Nama          string
	Alamat        string
	Tanggal_Lahir time.Time
	Nomor_Telepon string
}

type ParamsCreatePasien struct {
	ID      int    		`json:"id"`
	Nama      		string 		`json:"nama" form:"nama" binding:"required"`
	Alamat     		string 		`json:"alamat" form:"alamat" binding:"required"`
	Tanggal_Lahir   time.Time 	`json:"tanggal_lahir" form:"tanggal_lahir" binding:"required"`
	Nomor_Telepon  	string 		`json:"nomor_telepon" form:"nomor_telepon" binding:"required"`
	CreatedAt 		string
}

type ParamsUpdatePasien struct {
	ID      int    		`json:"id" form:"id" binding:"required"`
	Nama      		string 		`json:"nama" form:"nama"`
	Alamat     		string 		`json:"alamat" form:"alamat"`
	Tanggal_Lahir   time.Time 	`json:"tanggal_lahir" form:"tanggal_lahir"`
	Nomor_Telepon  	string 		`json:"nomor_telepon" form:"nomor_telepon"`
	UpdatedAt     	string
}

type ParamsGetListPasien struct {
	ID    	int    `json:"id" form:"id"`
	Nama   			string `json:"nama" form:"nama"`
	Limit  			*int   `json:"limit" form:"limit"`
	Offset 			*int   `json:"offset" form:"offset"`
}
type DataGetPasien struct {
	ID      int    		`json:"id"`
	Nama      		string 		`json:"nama"`
	Alamat     		string 		`json:"alamat"`
	Tanggal_Lahir   time.Time 	`json:"tanggal_lahir"`
	Nomor_Telepon  	string 		`json:"nomor_telepon"`
	CreatedAt 		string 		`json:"created_at"`
}
type ParamsPasien struct {
	ID       int 	`json:"id" form:"id" binding:"required"`
	DeletedAt 		string
}

