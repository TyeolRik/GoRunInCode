package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func main() {
	fmt.Println("This Output is from another package!!")
	fmt.Println("")
	fmt.Println("Function test :: Get Any random")
	fmt.Println(getRandom())
}

func getRandom() uint64 {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(b)
}
