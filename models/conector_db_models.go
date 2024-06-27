package models

import (

	"api/config"
)

var (
	db  = config.DBInit()
	idb = config.InDB{DB: db}
	// idb = inDB

)
