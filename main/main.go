package main

import (
	"fmt"
	"library-manager/library"
)

func main() {
	fmt.Print(summary())

}

func summary() string {
	return library.Summary()
}
