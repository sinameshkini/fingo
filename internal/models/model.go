package models

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

type ID snowflake.ID

func (i ID) String() string {
	return fmt.Sprintf("%d", i)
}

func ParseID(in string) (ID, error) {
	id, err := snowflake.ParseString(in)
	return ID(id), err
}

func ParseIDf(in string) ID {
	id, _ := snowflake.ParseString(in)
	return ID(id)
}

type Model struct {
	ID        ID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
