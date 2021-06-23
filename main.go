package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")
	fmt.Printf("Hello %s, how is it going?\n", user)
}
