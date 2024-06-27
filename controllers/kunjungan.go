package controllers

import (
	//"github.com/dgrijalva/jwt-go"
	"net/http"

	"api/models"

	"api/structs"
	"github.com/gin-gonic/gin"
)

func GetKunjungan(c *gin.Context) {

	params := structs.ParamsGetListKunjungan{}

	response := structs.Response{}
	kunjungans := []structs.DataGetKunjungan{}

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
		kunjungans, Count, err = models.GetKunjungans(params)

		if err == nil {
			response.Count = Count
			response.Data = kunjungans

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func CreateKunjungan(c *gin.Context) {

	params := structs.ParamsCreateKunjungan{}

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
		Data, err := models.CreateKunjungans(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func UpdateKunjungan(c *gin.Context) {

	params := structs.ParamsUpdateKunjungan{}

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
		Data, err := models.UpdateKunjungans(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func DeleteKunjungan(c *gin.Context) {

	params := structs.ParamsKunjungan{}

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
		Data, err := models.DeleteKunjungans(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

