package controllers

import (
	"fmt"
	"strconv"

	"myproject/models"
	"myproject/repositories"
	"myproject/request"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	albumRepo *repositories.AlbumRepository
}

func NewAlbumHandler() *AlbumHandler {
	return &AlbumHandler{
		albumRepo: repositories.NewAlbumRepository(),
	}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := h.albumRepo.GetAllAlbums()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error fetching albums",
		})
		return
	}

	c.JSON(200, gin.H{
		"albums": albums,
	})
}

func (h *AlbumHandler) GetAlbum(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	albumId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	album, err := h.albumRepo.GetAlbumByID(albumId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Album not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"album": album,
	})
}

func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var album request.AlbumDTO

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := h.albumRepo.CreateAlbum(&models.Album{
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error creating album",
		})
		return
	}

	c.JSON(200, gin.H{
		"albums": "albums",
	})
}

func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	albumId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	existingAlbum, err := h.albumRepo.GetAlbumByID(albumId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Album not found",
		})
		return
	}

	var album request.AlbumDTO

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	existingAlbum.Title = album.Title
	existingAlbum.Artist = album.Artist
	existingAlbum.Price = album.Price

	if err := h.albumRepo.UpdateAlbum(existingAlbum); err != nil {
		c.JSON(500, gin.H{
			"error": "Error updating album",
		})
		return
	}

	c.JSON(200, gin.H{
		"albums": "albums",
	})
}

func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	albumId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	existingAlbum, err := h.albumRepo.GetAlbumByID(albumId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Album not found",
		})
		return
	}

	if err := h.albumRepo.DeleteAlbumByID(existingAlbum); err != nil {
		c.JSON(500, gin.H{
			"error": "Error deleting album",
		})
		return
	}

	c.JSON(200, gin.H{
		"albums": "albums",
	})
}
