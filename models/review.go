package models

type Review struct {
    IdReview     int64   `gorm:"primaryKey;autoIncrement;type:BIGINT" json:"id_review"`
    IdProduct    int64   `gorm:"not null" json:"id_product"`
    IdUser       int64   `gorm:"not null" json:"id_user"`
    ReviewProduk string  `gorm:"type:text" json:"ulasan"`
    Product      Product `gorm:"foreignKey:IdProduct;references:IdProduct"`
    User         User    `gorm:"foreignKey:IdUser;references:IdUser"`
}
