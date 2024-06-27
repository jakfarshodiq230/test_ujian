package controllers

import (
	//"github.com/dgrijalva/jwt-go"
	"net/http"

	"api/models"

	"api/structs"
	"github.com/gin-gonic/gin"
)

func GetDokter(c *gin.Context) {

	params := structs.ParamsGetListDokter{}

	response := structs.Response{}
	dokters := []structs.DataGetDokter{}

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
		dokters, Count, err = models.GetDokters(params)

		if err == nil {
			response.Count = Count
			response.Data = dokters

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func CreateDokter(c *gin.Context) {

	params := structs.ParamsCreateDokter{}

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
		Data, err := models.CreateDokters(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func UpdateDokter(c *gin.Context) {

	params := structs.ParamsUpdateDokter{}

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
		Data, err := models.UpdateDokters(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func DeleteDokter(c *gin.Context) {

	params := structs.ParamsDokter{}

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
		Data, err := models.DeleteDokters(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

