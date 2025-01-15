package controllers

import (
	"net/http"
	"strconv"

	"myproject/models"
	"myproject/repositories"
	"myproject/request"
	"myproject/utils"

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
		c.JSON(500, utils.FormatDefaultError(err, "Error fetching albums"))
		return
	}

	c.JSON(200, utils.ResponseData(albums))
}

func (h *AlbumHandler) GetAlbum(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	albumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatDefaultError(err, "Invalid album ID"))
		return
	}

	album, err := h.albumRepo.GetAlbumByID(albumId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "Album not found"))
		return
	}

	c.JSON(200, utils.ResponseData(album))
}

func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var album request.AlbumDTO

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	err := h.albumRepo.CreateAlbum(&models.Album{
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	})
	if err != nil {
		c.JSON(500, utils.FormatDefaultError(err, "Error creating album"))
		return
	}

	c.JSON(200, utils.ResponseData(album))
}

func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	albumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatDefaultError(err, "Invalid album ID"))
		return
	}

	existingAlbum, err := h.albumRepo.GetAlbumByID(albumId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "Album not found"))
		return
	}

	var album request.AlbumDTO

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	existingAlbum.Title = album.Title
	existingAlbum.Artist = album.Artist
	existingAlbum.Price = album.Price

	if err := h.albumRepo.UpdateAlbum(existingAlbum); err != nil {
		c.JSON(500, utils.FormatDefaultError(err, "Error updating album"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(existingAlbum))
}

func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
	albumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatDefaultError(err, "Invalid album ID"))
		return
	}

	existingAlbum, err := h.albumRepo.GetAlbumByID(albumId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.FormatDefaultError(err, "Album not found"))
		return
	}

	if err := h.albumRepo.DeleteAlbumByID(existingAlbum); err != nil {
		c.JSON(500, utils.FormatDefaultError(err, "Error deleting album"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(existingAlbum))
}
