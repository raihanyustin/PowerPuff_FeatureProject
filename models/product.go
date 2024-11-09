package models

type Product struct {
	IdProduct			int64	`gorm:"primaryKey" json:"id_product"`
	NamaProduk 			string	`gorm:"type:varchar(300)" json:"nama_product"`
	DeskripsiProduk		string	`gorm:"type:text" json:"deskripsi"`
	Reviews				[]Review `gorm:"foreignKey:id_product" json:"reviews"`
}