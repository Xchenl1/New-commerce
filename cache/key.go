package cache

import (
	"fmt"
	"strconv"
)

const (
	RanKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
