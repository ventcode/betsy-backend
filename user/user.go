package user

import "gorm.io/gorm"

type User struct {
    gorm.Model
    ExternalId string `gorm:"not null;unique"`
    MoneyAmount int `gorm:"not null;default:0"`
}
