package cli

import "fmt"

// Report a normal error value.
func cliReportError(err error) {
	fmt.Printf("Error: %s\n", err)
}
