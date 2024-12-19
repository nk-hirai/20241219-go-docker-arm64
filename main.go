package main

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
)

func main() {
	str := "Hello, world!"
	bs, err := sonic.Marshal(str)
	if err != nil {
		log.Fatalf("Failed to marshal string: %v", err)
	}

    var data string
    json := []byte("\"Hello, world!\"")
    err = sonic.Unmarshal(json, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	fmt.Println(string(bs))
}
