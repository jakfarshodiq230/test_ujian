package models

import (
	"api/structs"
)

func GetDokters(params structs.ParamsGetListDokter) ([]structs.DataGetDokter, int, error) {

	var (
		data []structs.DataGetDokter
		//t      Dokter.Component
		Count int
	)

	err := idb.DB.Table("dokters")
	err = err.Where("deleted_at is null")

	if params.Nama != "" {
		err = err.Where("dokters.nama like ?", "%"+params.Nama+"%")
	}
	err = err.Count(&Count)
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Limit)
	}

	err = err.Find(&data)
	errx := err.Error

	return data, Count, errx
}

func CreateDokters(params structs.ParamsCreateDokter) (structs.ParamsCreateDokter, error) {
	var t = structs.Component{}
	params.CreatedAt = t.GetTimeNow()
	create := idb.DB.Table("dokters").Create(&params)
	err := create.Error
	return params, err
}

func UpdateDokters(params structs.ParamsUpdateDokter) (structs.ParamsUpdateDokter, error) {
	var t = structs.Component{}
	params.UpdatedAt = t.GetTimeNow()
	update := idb.DB.Table("dokters").Where("id = ?", params.ID).Updates(&params)
	err := update.Error
	return params, err
}

func DeleteDokters(params structs.ParamsDokter) (structs.ParamsDokter, error) {

	delete := idb.DB.Table("dokters").Where("d = ?", params.ID).Delete(&params)
	err := delete.Error
	return params, err
}
