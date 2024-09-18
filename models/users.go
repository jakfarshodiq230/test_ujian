package models

import (
	"api/structs"
)

func GetUsers(params structs.ParamsGetListUser) ([]structs.DataGetUser, int, error) {
	var (
		data  []structs.DataGetUser
		count int
	)

	// Start the query
	query := idb.DB.Table("users").Where("deleted_at IS NULL")

	// Apply filters
	if params.Name != "" {
		query = query.Where("users.name LIKE ?", "%"+params.Name+"%")
	}

	// Count the total number of records
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	}
	if params.Offset != nil {
		query = query.Offset(*params.Offset)
	}

	// Fetch the data
	err = query.Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}


func CreateUsers(params structs.ParamsCreateUser) (structs.ParamsCreateUser, error) {
	var t = structs.Component{}
	params.CreatedAt = t.GetTimeNow()
	create := idb.DB.Table("users").Create(&params)
	err := create.Error
	return params, err
}

func UpdateUsers(params structs.ParamsUpdateUser) (structs.ParamsUpdateUser, error) {
	var t = structs.Component{}
	params.UpdatedAt = t.GetTimeNow()
	update := idb.DB.Table("users").Where("id = ?", params.ID).Updates(&params)
	err := update.Error
	return params, err
}

func DeleteUsers(params structs.ParamsUser) (structs.ParamsUser, error) {

	delete := idb.DB.Table("users").Where("id = ?", params.Id).Delete(&params)
	err := delete.Error
	return params, err
}

func Login(params structs.ParamsLogin) ([]structs.DataGetUser, error) {

	var (
		data []structs.DataGetUser
	)

	login := idb.DB.Table("users")
	login = login.Where("deleted_at is null")
	login = login.Where("users.username = ?", params.Username)
	login = login.Where("users.password = ?", params.Password)

	login = login.First(&data)

	errx := login.Error

	return data, errx
}
