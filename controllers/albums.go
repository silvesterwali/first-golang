package controllers

import "github.com/gin-gonic/gin"


func GetAlbums(c *gin.Context) {
	c.JSON(200, gin.H{
		"albums": "albums",
	})
}


func GetAlbum(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"id": id,
	})

}


func CreateAlbum(c *gin.Context) {
	c.JSON(200, gin.H{
		"albums": "albums",
	})
}


func UpdateAlbum(c *gin.Context) {
	c.JSON(200, gin.H{
		"albums": "albums",
	})
}


func DeleteAlbum(c *gin.Context) {
	c.JSON(200, gin.H{
		"albums": "albums",
	})
}