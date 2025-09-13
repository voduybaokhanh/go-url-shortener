package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Email        string    `gorm:"uniqueIndex;size:255"`
    PasswordHash string    `gorm:"size:255"`
    CreatedAt    time.Time
    Links        []Link    `gorm:"foreignKey:UserID"`
}
