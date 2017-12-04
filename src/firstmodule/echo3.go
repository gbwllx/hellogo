package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], ""))
	fmt.Printf("Hello!%s\n", "harry")
	fmt.Println("Hello!")
}
