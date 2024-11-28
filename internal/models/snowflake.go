package models

import (
	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func InitID(machineID int64) {
	snowflakeNode, _ = snowflake.NewNode(machineID)
}

func Next() ID {
	return ID(snowflakeNode.Generate())
}
