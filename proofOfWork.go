package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const Difficulty = 1

func main() {
	t := time.Now().String()
	for i := 0; ; i++ {
		s := fmt.Sprintf("%x", i) + t
		if !isHashValid(calculateHash(s), Difficulty) {
			fmt.Println(calculateHash(s), " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(s)
			fmt.Println(calculateHash(s), " work done!")
			break
		}

	}

}

func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isHashValid(s string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(s, prefix)
}
