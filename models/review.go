package models

import "time"

type Review struct {
<<<<<<< HEAD
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UsernameID  int       `json:"username_id" binding:"required"`
	Username    string    `json:"username" binding:"required"`
	IsAnonymous bool      `json:"is_anonymous"`
	ProductID   int       `json:"product_id" binding:"required"`
	ProductName string    `json:"product_name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Rating      int       `json:"rating" binding:"gte=1,lte=5"`
	TextReview  string    `json:"text_review" binding:"required"`
	Likes       int       `json:"likes" binding:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" binding:"required"`
=======
    ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
    UsernameID    int       `json:"username_id"`
    Username      string    `json:"username"`
    IsAnonymous   bool      `json:"is_anonymous"`
    ProductID     int       `json:"product_id"`
    ProductName   string    `json:"product_name"`
    Category      string    `json:"category"`
    Rating        int       `json:"rating"`
    TextReview    string    `json:"text_review"`
    Likes         int       `json:"likes"`
    CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
>>>>>>> 969d07304c1a8a45f0327a10ae136286f7566c58
}
