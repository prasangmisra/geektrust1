package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi")
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Some problem here")
	} else {
		fmt.Println("All good, the param is:", args[1])
	}

}
