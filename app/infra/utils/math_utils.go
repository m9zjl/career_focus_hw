package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToInt(str string) int64 {
	parsed, _ := strconv.ParseInt(strings.Replace(str, "0x", "", -1), 16, 32)
	return parsed
}

func IntToHex(num int64) string {
	return fmt.Sprintf("0x%x", num)
}
