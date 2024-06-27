package structs

import (
	"time"
    "gorm.io/gorm"
)

type Kunjungans struct {
	gorm.Model
	ID_Dokter          	int
	ID_Pasien  			int
	Tanggal_Kunjungan 	time.Time
	Diagnosa 			string
}

type CustomTime struct {
	time.Time
}

type ParamsCreateKunjungan struct {
	ID		int 	  `json:"id"`
	ID_Dokter       	int       `json:"id_dokter" form:"id_dokter" binding:"required"`
	ID_Pasien      		int 	  `json:"id_pasien" form:"id_pasien" binding:"required"`
	Tanggal_Kunjungan   time.Time `json:"tanggal_kunjungan" form:"tanggal_kunjungan" binding:"required"`
	Diagnosa 			string    `json:"diagnosa" form:"diagnosa" binding:"required"`
	CreatedAt 			string
}

type ParamsUpdateKunjungan struct {
	ID		int 	  `json:"id" form:"id" binding:"required"`
	ID_Dokter       	int       `json:"id_dokter" form:"id_dokter" binding:"required"`
	ID_Pasien      		int		  `json:"id_pasien" form:"id_pasien" binding:"required"`
	Tanggal_Kunjungan   time.Time `json:"tanggal_kunjungan" form:"tanggal_kunjungan" binding:"required"`
	Diagnosa 			string    `json:"diagnosa" form:"diagnosa" binding:"required"`
	UpdatedAt     		string
}

type ParamsGetListKunjungan struct {
	ID     	int       `json:"id"`
	ID_Dokter       	int       `json:"id_dokter"`
	ID_Pasien      		int 	  `json:"id_pasien"`
    Nama_Dokter      	string    `json:"nama_dokter"`
    Nama_Pasien      	string    `json:"nama_pasien"`
    Tanggal_Kunjungan 	time.Time `json:"tanggal_kunjungan"`
    Diagnosa         	string    `json:"diagnosa"`
	Limit  				*int   	  `json:"limit" form:"limit"`
	Offset 				*int   	  `json:"offset" form:"offset"`
}

type DataGetKunjungan struct {
	ID     	int       `json:"id"`
	ID_Dokter       	int    	  `json:"id_dokter"`
	ID_Pasien      		int 	  `json:"id_pasien"`
    Nama_Dokter      	string    `json:"nama_dokter"`
    Nama_Pasien      	string    `json:"nama_pasien"`
    Tanggal_Kunjungan 	time.Time `json:"tanggal_kunjungan"`
    Diagnosa         	string    `json:"diagnosa"`
    CreatedAt        	string    `json:"created_at"`
}
type ParamsKunjungan struct {
	ID       	int 	  `json:"id" form:"id" binding:"required"`
	DeletedAt 			string
}
