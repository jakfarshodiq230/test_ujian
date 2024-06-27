package models

import (
	"api/structs"
)

func GetPasiens(params structs.ParamsGetListPasien) ([]structs.DataGetPasien, int, error) {

	var (
		data []structs.DataGetPasien
		//t      pasien.Component
		Count int
	)

	err := idb.DB.Table("pasiens")
	err = err.Where("deleted_at is null")

	if params.Nama != "" {
		err = err.Where("pasiens.nama like ?", "%"+params.Nama+"%")
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

func CreatePasiens(params structs.ParamsCreatePasien) (structs.ParamsCreatePasien, error) {
	var t = structs.Component{}
	params.CreatedAt = t.GetTimeNow()
	create := idb.DB.Table("pasiens").Create(&params)
	err := create.Error
	return params, err
}

func UpdatePasiens(params structs.ParamsUpdatePasien) (structs.ParamsUpdatePasien, error) {
	var t = structs.Component{}
	params.UpdatedAt = t.GetTimeNow()
	update := idb.DB.Table("pasiens").Where("id = ?", params.ID).Updates(&params)
	err := update.Error
	return params, err
}

func DeletePasiens(params structs.ParamsPasien) (structs.ParamsPasien, error) {

	delete := idb.DB.Table("pasiens").Where("id = ?", params.ID).Delete(&params)
	err := delete.Error
	return params, err
}
