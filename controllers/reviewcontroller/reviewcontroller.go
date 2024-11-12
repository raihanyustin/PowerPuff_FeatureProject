package reviewcontroller

import (
	"net/http"
	"PowerPuff_ReviewBarang/models"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var reviews []models.Review

	models.DB.Find(&reviews)
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})

}

func Show(c *gin.Context) {
    var reviews []models.Review
    productName := c.Param("ProductName")

    // Query berdasarkan ProductName
    if err := models.DB.Debug().Where("product_name = ?", productName).Find(&reviews).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
            return
        }
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    // Jika tidak ada hasil
    if len(reviews) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }

    // Kembalikan hasil dalam format JSON
    c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func Create(c *gin.Context) {

	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&review)
	c.JSON(http.StatusOK, gin.H{"review": review})
}

func Update(c *gin.Context) {
    var review models.Review
    id := c.Param("ID")

    // Bind JSON ke struct review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    result := models.DB.Model(&models.Review{}).Where("id = ?", id).Updates(review)
    if result.Error != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
        return
    }
    if result.RowsAffected == 0 {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate review"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}