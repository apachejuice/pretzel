package cli

import (
	"fmt"
	"os"

	"github.com/apachejuice/pretzel/internal/analysis/pass"
	"github.com/apachejuice/pretzel/internal/errors"
)

// Handle the compilation process with the given options.
func DoCompile(opts Options) bool {
	conf, err := ReadCompilerConfig(opts.ConfigFile)
	if err != nil {
		cliReportError(err)
		return false
	}

	err = conf.FindFiles()
	if err != nil {
		cliReportError(err)
		return false
	}

	if len(conf.Files) == 0 {
		cliReportError(fmt.Errorf("no source files found in directory '%s' (%d ignored)", conf.SourceDir, conf.Ignored))
		return false
	}

	f := &FileGroup{Config: conf}
	ok, errs, err := f.ParseAll()
	o := errors.NewOutputStreamPrinter(os.Stderr, false, true)

	if !ok {
		if err != nil {
			cliReportError(err)
		} else {
			for _, err := range errs {
				o.Report(err)
			}
		}

		return false
	}

	c := pass.NewComboPass(
		pass.NewScopeApplyPass(),
	)

	for _, node := range f.Output {
		c.Run(node)
	}

	for _, err := range c.Errors {
		o.Report(err)
	}

	return true
}
