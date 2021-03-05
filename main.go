package main

import (
	"log"
	"net/http"

	"github.com/revianto/goTemplate/contrloller/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// gUser is group user
	gUser := router.Group("/digileaps/v1/user")
	{
		gUser.Static("/media", "/upload/user")
		gUser.GET("/get", user.GetData)
		gUser.POST("/insert", user.AddData)
		gUser.PUT("/update", user.UpdateData)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "404", "message": "Not Found"})
	})

	log.Println("[*] Starting Serve :80 ...")
	log.Println("[+] Served ...")
	router.Run(":80")
}
