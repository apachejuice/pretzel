package cli

import (
	"fmt"
)

// The main function.
func Main(args []string) int {
	opts, err := Parse(args)
	if err != nil {
		fmt.Println("Exiting due to above error.")
		return 1
	}

	ok := DoCompile(opts)
	if !ok {
		return 1
	}

	return 0
}
