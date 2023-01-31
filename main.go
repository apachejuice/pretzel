package main

import (
	"os"

	"github.com/apachejuice/pretzel/internal/cli"
)

func main() {
	os.Exit(cli.Main(os.Args))
}
