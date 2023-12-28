package models

import "time"

type User struct {
    ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    Username  string     `json:"username" grom:"required"`
    Email     string     `json:"email" grom:"required,unique"`
    Password  string     `json:"password" gorm:"required,min=6"`
    Photos    []Photo    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
}