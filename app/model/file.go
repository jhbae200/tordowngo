package model

import (
	"time"
)

type File struct {
	FileNum    int    `gorm:"primary_key;AUTO_INCREMENT"`
	FileName   string `gorm:"size:255"`
	FileType   string `gorm:"size:255"`
	FileSize   int64
	FileStatus int       `gorm:"default:0"`
	FileDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
