package models

type Product struct {
    IdProduct       int64    `gorm:"primaryKey;autoIncrement;type:BIGINT" json:"id"`
    NamaProduk      string   `gorm:"type:varchar(300)" json:"nama_product"`
    DeskripsiProduk string   `gorm:"type:text" json:"deskripsi"`
    Reviews         []Review `gorm:"foreignKey:IdProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"` 
}