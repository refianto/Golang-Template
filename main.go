package main

import (
	"log"
	"net/http"

	"github.com/refianto/goTemplate/contrloller/kategori"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	gKategori := router.Group("/digileaps/v1/kategori")
	{
		gKategori.Static("/media", "/upload/kategori")
		gKategori.GET("/get", kategori.GetKategori)
		gKategori.POST("/insert", kategori.AddKategori)
		gKategori.PUT("/update", kategori.UpdateKategori)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "404", "message": "Not Found"})
	})

	log.Println("[*] Starting Serve :80 ...")
	log.Println("[+] Served ...")
	router.Run(":80")
}
