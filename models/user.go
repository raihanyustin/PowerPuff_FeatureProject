package models

type User struct {
    IdUser   int64    `gorm:"primaryKey;autoIncrement;type:BIGINT" json:"id"`
    NamaUser string   `gorm:"type:varchar(50)" json:"username"`
    Reviews  []Review `gorm:"foreignKey:IdUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
}