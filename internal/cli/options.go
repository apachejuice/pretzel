package cli

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	// The configuration file to use
	ConfigFile string `short:"C" long:"config-file" value-name:"FILE" description:"The compiler configuration file to use" default:"pretzel.json"`
}

// Parse options from the command line
func Parse(args []string) (Options, error) {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[OPTIONS] [ARGUMENTS]"

	_, err := parser.Parse()
	if err != nil {
		return Options{}, err
	}

	return opts, nil
}
