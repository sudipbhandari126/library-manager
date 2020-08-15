package main

import "fmt"

func main() {
	fmt.Print(greeter("welcome"))
}

func greeter(msg string) string {
	return fmt.Sprintf("Hello %s", msg)
}
