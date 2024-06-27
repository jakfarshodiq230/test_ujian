package models

import (
	"api/structs"
)

func GetKunjungans(params structs.ParamsGetListKunjungan) ([]structs.DataGetKunjungan, int, error) {
	var (
		data  []structs.DataGetKunjungan
		count int
	)

	query := db.Table("kunjungans").
		Select("kunjungans.id_dokter, kunjungans.id_pasien, kunjungans.id, dokters.nama as nama_dokter, pasiens.nama as nama_pasien, kunjungans.tanggal_kunjungan, kunjungans.diagnosa, kunjungans.created_at").
		Joins("left join dokters on kunjungans.id_dokter = dokters.id").
		Joins("left join pasiens on kunjungans.id_pasien = pasiens.id").
		Where("kunjungans.deleted_at IS NULL")

	if !params.Tanggal_Kunjungan.IsZero() {
		tanggalKunjunganStr := params.Tanggal_Kunjungan.Format("2006-01-02")
		query = query.Where("kunjungans.tanggal_kunjungan LIKE ?", "%"+tanggalKunjunganStr+"%")
	}

	query = query.Count(&count)
	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	}
	if params.Offset != nil {
		query = query.Offset(*params.Offset)
	}

	err := query.Find(&data).Error
	return data, count, err
}

func CreateKunjungans(params structs.ParamsCreateKunjungan) (structs.ParamsCreateKunjungan, error) {
	var t = structs.Component{}
	params.CreatedAt = t.GetTimeNow()
	create := idb.DB.Table("kunjungans").Create(&params)
	err := create.Error
	return params, err
}

func UpdateKunjungans(params structs.ParamsUpdateKunjungan) (structs.ParamsUpdateKunjungan, error) {
	var t = structs.Component{}
	params.UpdatedAt = t.GetTimeNow()
	update := idb.DB.Table("kunjungans").Where("id = ?", params.ID).Updates(&params)
	err := update.Error
	return params, err
}

func DeleteKunjungans(params structs.ParamsKunjungan) (structs.ParamsKunjungan, error) {

	delete := idb.DB.Table("kunjungans").Where("id = ?", params.ID).Delete(&params)
	err := delete.Error
	return params, err
}
