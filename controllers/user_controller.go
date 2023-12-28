package controllers

import (
	"golang-api/database"
	"golang-api/helpers"
    "golang-api/app"
	"golang-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user app.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password user sebelum disimpan ke dalam database
	hashedPassword, err := app.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi kata sandi"})
		return
	}
	user.Password = hashedPassword

	// Membuat data pengguna baru di dalam database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pengguna"})
		return
	}

	// Menghasilkan token JWT untuk pengguna yang terdaftar
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token"})
		return
	}
	// Membuat Token
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
	// Mendapatkan ID pengguna dari URL parameter
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Memeriksa apakah pengguna yang sedang mencoba mengupdate adalah pemilik akun
	authUserID, exists := c.Get("userID")
	if !exists || authUserID.(uint) != uint(userID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memperbarui data pengguna dalam database
	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(&updatedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	// Pesan Sukses
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	// Mendapatkan ID pengguna dari URL parameter
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Memeriksa apakah pengguna yang sedang mencoba menghapus adalah pemilik akun
	authUserID, exists := c.Get("userID")
	if !exists || authUserID.(uint) != uint(userID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Menghapus pengguna dari database
	if err := database.DB.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	// Pesan Sukses
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
func GetUser(c *gin.Context) {
	user := models.User{}

	if err := database.DB.Preload("Photos").First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photos"})
		return
	}

	// Pesan Sukses
	c.JSON(http.StatusOK, user)
}
