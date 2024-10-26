package utils

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestHashmap(t *testing.T) {
	sha := sha256.New()
	fmt.Println(sha.Write([]byte("andrasfdsssssssssssssssssssssss")))
}
