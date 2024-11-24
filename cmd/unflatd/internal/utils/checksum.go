package utils

import (
	"fmt"
	"hash/crc32"
)

func Checksum(input string) string {
	hash := crc32.ChecksumIEEE([]byte(input))
	return fmt.Sprintf("%08x", hash)
}
