package main

import (
	"fmt"
	"os"

	actions "github.com/prasangmisra/geektrust1/actions"
)

func main() {
	fmt.Println("Hi")
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Some problem here")
		return
	}
	fmt.Println("All good, the param is:", args[1])
	actions.Init(args[1])
}
