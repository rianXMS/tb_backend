package main

import (
	"fmt"
	"tb_backend/Features"
)

func main() {

	var a = Features.Load()

	for i := range a {
		fmt.Printf("Executing %d\n", i)
		a[i]()
	}

}
func test() {
	fmt.Printf("Test")
}
