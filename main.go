package main

import (
	"fmt"

	"github.com/neermitt/terraform-config-inspect/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
