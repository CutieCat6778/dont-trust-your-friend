package lib

import (
	"github.com/bwmarrin/snowflake"
)

var (
	CurrentIndex int64 = 0
)

func GenerateID() (*[]byte, *CustomError) {
	node, err := snowflake.NewNode(CurrentIndex)
	if err != nil {
		return nil, NewError("Snowflake Node Error", 500, SnowflakeService)
	}
	id := node.Generate().Bytes()
	return &id, nil
}
