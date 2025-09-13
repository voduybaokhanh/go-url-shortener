package models

import "time"

type Link struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"index"`
    OriginalURL string    `gorm:"size:2048"`
    ShortCode   string    `gorm:"uniqueIndex;size:64"`
    CreatedAt   time.Time
}
