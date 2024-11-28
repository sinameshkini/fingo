package models

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

type ID snowflake.ID

type Model struct {
	ID        ID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
