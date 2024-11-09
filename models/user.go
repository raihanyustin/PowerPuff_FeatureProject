package models

type User struct {
	IdUser				int64	`gorm:"primaryKey" json:"id_user"`
	NamaUser			string	`gorm:"type:varchar(50)" json:"username"`
	Reviews				[]Review `gorm:"foreignKey:id_user" json:"reviews"`
}