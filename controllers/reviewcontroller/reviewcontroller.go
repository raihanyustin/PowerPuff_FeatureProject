package reviewcontroller

import (
	"net/http"
	"PowerPuff_ReviewBarang/models"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
    "strconv"
)

// Struktur Stack untuk menyimpan ulasan
type Stack []models.Review

func (s *Stack) Push(review models.Review) {
	*s = append(*s, review)
}

// Variabel global untuk stack
var reviewStack = make(Stack, 0)
func (s *Stack) Peek() (models.Review, bool) {
	if len(*s) == 0 {
		return models.Review{}, false
	}
	return (*s)[len(*s)-1], true
}

// Fungsi untuk menambahkan data ke stack
func PushToStack(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data tidak valid", "error": err.Error()})
		return
	}
    
	reviewStack.Push(review)
	c.JSON(http.StatusOK, gin.H{"message": "Ulasan berhasil ditambahkan ke stack", "review": review})
    
    models.DB.Create(&review)       
    c.JSON(http.StatusOK, gin.H{"review": review})
}

func PeekStack(c *gin.Context) {
	review, ok := reviewStack.Peek()
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Stack kosong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ulasan teratas di stack", "review":review})
}

func Index(c *gin.Context) {

	var reviews []models.Review

	models.DB.Find(&reviews)
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})

}

// Fungsi untuk menampilkan semua ulasan berdasarkan stack (Last in FIrst Out)
func GetAllFromStack(c *gin.Context) {
    var reviewStack Stack
    var reversedStack Stack
    models.DB.Find(&reviewStack)

    if len(reviewStack) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Data Review kosong"})
        return
    }

    for i := len(reviewStack) - 1; i >= 0; i-- {
        reversedStack = append(reversedStack, reviewStack[i])
    }

    c.JSON(http.StatusOK, gin.H{
        "stack": reversedStack,
    })
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

    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func SearchByProductAndRating(c *gin.Context) {
	namaProduk := c.DefaultQuery("product_name", "")
	rating := c.DefaultQuery("rating", "")

	if namaProduk == "" || rating == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nama produk dan rating wajib diisi"})
		return
	}

	ratingInt, err := strconv.Atoi(rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format rating tidak valid"})
		return
	}

	var ulasan []models.Review
	result := models.DB.Where("product_name = ? AND rating = ?", namaProduk, ratingInt).Find(&ulasan)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}

	if len(ulasan) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": ulasan})
}