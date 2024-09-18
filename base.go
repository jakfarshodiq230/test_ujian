package main

import (
    "api/controllers"
    "api/middleware"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
	auth := middleware.Auth

	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"}, // Or specify allowed origins
        AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
	// users
	router.POST("/user", controllers.CreateUser)
	router.GET("/users", auth, controllers.GetUsers)
	router.PUT("/user", auth, controllers.UpdateUser)
	router.DELETE("/user", auth, controllers.DeleteUser)
	router.POST("/login", controllers.Login)

	// pasien
	router.POST("/pasien", auth, controllers.CreatePasien)
	router.GET("/pasien", auth, controllers.GetPasien)
	router.PUT("/pasien", auth, controllers.UpdatePasien)
	router.DELETE("/pasien", auth, controllers.DeletePasien)

	//dokter
	router.POST("/dokter", auth, controllers.CreateDokter)
	router.GET("/dokter", auth, controllers.GetDokter)
	router.PUT("/dokter", auth, controllers.UpdateDokter)
	router.DELETE("/dokter", auth, controllers.DeleteDokter)

	//kunjungan
	router.POST("/kunjungan", auth, controllers.CreateKunjungan)
	router.GET("/kunjungan", auth, controllers.GetKunjungan)
	router.PUT("/kunjungan", auth, controllers.UpdateKunjungan)
	router.DELETE("/kunjungan", auth, controllers.DeleteKunjungan)

	router.Run(":4000")
}
