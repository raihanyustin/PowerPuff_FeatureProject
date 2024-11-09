package models

type Review struct {
	IdReview 			int64 `gorm:"primaryKey" json:"id_review"`
	IdUser				int64 `gorm:"not null" json:"id_user"`
	IdProduct			int64 `gorm:"not null" json:"id_product"`
	ReviewProduk		string `gorm:"type:text" json:"ulasan"`
	User 				User `gorm:"foreignKey:id_user"`
	Product 			Product `gorm:"foreignKey:id_product"`
}