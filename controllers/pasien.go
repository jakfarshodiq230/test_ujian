package controllers

import (
	//"github.com/dgrijalva/jwt-go"
	"net/http"

	"api/models"

	"api/structs"
	"github.com/gin-gonic/gin"
)

func GetPasien(c *gin.Context) {

	params := structs.ParamsGetListPasien{}

	response := structs.Response{}
	pasiens := []structs.DataGetPasien{}

	err := c.ShouldBind(&params)
	var Count int

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = "validation " + mess
		c.JSON(400, response)
	} else {
		pasiens, Count, err = models.GetPasiens(params)

		if err == nil {
			response.Count = Count
			response.Data = pasiens

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func CreatePasien(c *gin.Context) {

	params := structs.ParamsCreatePasien{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.CreatePasiens(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func UpdatePasien(c *gin.Context) {

	params := structs.ParamsUpdatePasien{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.UpdatePasiens(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func DeletePasien(c *gin.Context) {

	params := structs.ParamsPasien{}

	response := structs.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.Message = mess
		c.JSON(400, response)
	} else {
		Data, err := models.DeletePasiens(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

