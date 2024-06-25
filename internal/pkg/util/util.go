package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

//FileExist ...
func FileExist(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//Sha1 ...
func Sha1(data string) string {
	hash := sha1.New()
	io.WriteString(hash, data)
	return fmt.Sprintf(`%x`, hash.Sum(nil))
}
