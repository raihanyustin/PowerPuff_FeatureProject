package models

import "time"

type Review struct {
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
}
