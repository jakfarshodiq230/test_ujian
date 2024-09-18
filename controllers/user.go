package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"api/models"

	"api/structs"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	params := structs.ParamsGetListUser{}

	response := structs.Response{}
	users := []structs.DataGetUser{}

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
		users, Count, err = models.GetUsers(params)

		if err == nil {
			response.Count = Count
			response.Data = users

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func CreateUser(c *gin.Context) {

	params := structs.ParamsCreateUser{}

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
		Data, err := models.CreateUsers(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func UpdateUser(c *gin.Context) {

	params := structs.ParamsUpdateUser{}

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
		Data, err := models.UpdateUsers(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func DeleteUser(c *gin.Context) {

	params := structs.ParamsUser{}

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
		Data, err := models.DeleteUsers(params)

		if err == nil {
			response.Data = Data

			c.JSON(http.StatusOK, response)
		} else {
			c.JSON(500, response)
		}
	}

}

func Login(c *gin.Context) {
    var params structs.ParamsLogin
    var response structs.Response

    // Bind request parameters to the struct
    if err := c.ShouldBind(&params); err != nil {
        response.Message = err.Error()
        c.JSON(http.StatusBadRequest, response)
        return
    }

    // Attempt to log in
    Data, err := models.Login(params)
    if err != nil {
        response.Message = "Internal server error"
        c.JSON(http.StatusInternalServerError, response)
        return
    }

    // Check if any data was returned
    if len(Data) > 0 {
        sign := jwt.New(jwt.GetSigningMethod("HS256"))
        token, err := sign.SignedString([]byte("jakfarshodiq")) // Consider using an environment variable for the secret
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
            return
        }

        response.Message = token
        response.Data = Data
        c.JSON(http.StatusOK, response)
    } else {
        response.Message = "Username / Password Salah"
        c.JSON(http.StatusUnauthorized, response)
    }
}

