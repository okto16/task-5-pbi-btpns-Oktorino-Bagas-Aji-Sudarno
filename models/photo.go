package models

import "time"

type Photo struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Title     string    `json:"title" binding:"required"`
    Caption   string    `json:"caption"`
    PhotoURL  string    `json:"photo_url" binding:"required"`
    UserID    uint      `json:"user_id"`
    User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
