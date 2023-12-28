package controllers

import (
	"golang-api/database"
	"golang-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePhoto
func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	// Mengikat data permintaan ke struktur photo tanpa melakukan validasi atribut User
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menetapkan UserID berdasarkan user_id dari permintaan
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	photo.UserID = userID.(uint) // Mengasumsikan userID bertipe uint

	// Menyimpan foto ke dalam database
	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat foto"})
		return
	}

	// Pesan Sukses
	c.JSON(http.StatusCreated, gin.H{"message": "Foto berhasil dibuat"})
}

// GetPhotos
func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	// Dapatkan daftar foto beserta informasi pengguna dari database
	if err := database.DB.Preload("User").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photos"})
		return
	}
	// Pesan Sukses
	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto
func UpdatePhoto(c *gin.Context) {
	// Mendapatkan ID foto dari URL parameter
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	var updatedPhoto models.Photo

	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memeriksa apakah pengguna memiliki izin untuk mengupdate foto
	userID, _ := c.Get("userID")

	// Mengambil foto yang akan diupdate
	existingPhoto := models.Photo{}
	if err := database.DB.First(&existingPhoto, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Memeriksa apakah pengguna memiliki izin untuk mengupdate foto
	if existingPhoto.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	// Mengupdate foto
	if err := database.DB.Model(&models.Photo{}).Where("id = ? AND user_id = ?", photoID, userID).Updates(&updatedPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}
	// Pesan Sukses
	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto
func DeletePhoto(c *gin.Context) {
	// Mendapatkan ID foto dari URL parameter
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	// Memeriksa apakah pengguna memiliki izin untuk menghapus foto
	userID, _ := c.Get("userID")

	// Mengambil foto yang akan dihapus
	existingPhoto := models.Photo{}
	if err := database.DB.First(&existingPhoto, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}
	// Memeriksa apakah pengguna memiliki izin untuk menghapus foto
	if existingPhoto.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}
	// Menghapus foto
	if err := database.DB.Where("id = ? AND user_id = ?", photoID, userID).Delete(&models.Photo{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}
	// Pesan Sukses
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
