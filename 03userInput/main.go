package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	welcome := "Welcome to Go Programming"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello, ", name)
}
