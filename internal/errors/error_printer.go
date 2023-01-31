package errors

import (
	"fmt"
	"io"
	"strings"

	"github.com/apachejuice/pretzel/internal/common/util"
	"github.com/fatih/color"
)

type (
	// An error printer that displays errors to any kind of output channel.
	ErrorPrinter interface {
		// Returns the amount of errors, warnings and informational messages reported.
		Stats() (err, warn, info int)
		// Report the given error code.
		Report(e *Error) error
	}

	// An error printer to an output stream
	OutputStreamPrinter struct {
		// The output stream
		Output io.Writer
		// Whether or not to compact errors
		DoCompact bool
		// Whether or not to use colored output
		DoColor bool

		// the amount of errors, warnings, infos reported
		errs, warns, infos int
		// color instance
		c *color.Color
	}
)

// Creates a new output stream printer
func NewOutputStreamPrinter(writer io.Writer, doCompact, doColor bool) *OutputStreamPrinter {
	color := color.New()
	if !doColor {
		color.DisableColor()
	}

	return &OutputStreamPrinter{
		Output:    writer,
		DoCompact: doCompact,
		DoColor:   doColor,
		errs:      0,
		warns:     0,
		infos:     0,
		c:         color,
	}
}

var _ ErrorPrinter = &OutputStreamPrinter{}

func (o OutputStreamPrinter) Stats() (int, int, int) {
	return o.errs, o.warns, o.infos
}

func (o OutputStreamPrinter) printSubjectLine(colour color.Attribute, tag string, e *Error) {
	o.c.Fprint(o.Output, "[")
	o.c.Add(colour)
	o.c.Fprint(o.Output, tag)
	o.c.Add(color.Reset)
	o.c.Fprintf(o.Output, "]: %s", e.Formatted())

	if o.DoCompact {
		o.c.Fprintf(o.Output, " (in %q %d:%d-%d:%d)", e.Begin.Path, e.Begin.StartLine, e.Begin.StartColumn, e.End.EndLine, e.End.EndColumn)
	}

	o.c.Fprint(o.Output, "\n")
}

func (o OutputStreamPrinter) underline(e *Error) {
	width := util.NumberLen(e.Begin.StartLine) + 3
	o.printf(strings.Repeat(" ", width+e.Begin.StartColumn))
	diff := e.End.EndColumn - e.Begin.StartColumn

	if diff < 3 {
		o.printf("^")
	} else {
		o.printf("^")
		for i := 0; i < diff-2; i++ {
			o.printf("~")
		}

		o.printf("^")
	}

	o.printf("\n")
}

func (o OutputStreamPrinter) printf(msg string, args ...any) {
	o.Output.Write([]byte(fmt.Sprintf(msg, args...)))
}

func (o *OutputStreamPrinter) Report(e *Error) error {
	if e.Code.IsWarning() {
		return o.reportWarning(e)
	} else {
		return o.reportError(e)
	}
}

func (o *OutputStreamPrinter) reportError(e *Error) error {
	o.printSubjectLine(color.FgRed, "Error", e)
	if !o.DoCompact {
		beginLine := e.Begin.StartLine + 1 // human readable 1-based
		o.printf("%d | %s\n", beginLine, e.Source)
		o.underline(e)
	}

	return nil
}

func (o *OutputStreamPrinter) reportWarning(e *Error) error {
	o.printSubjectLine(color.FgYellow, "Warning", e)
	if !o.DoCompact {
		beginLine := e.Begin.StartLine + 1 // human readable 1-based
		o.printf("%d | %s\n", beginLine, e.Source)
		o.underline(e)
	}

	return nil
}

func (o *OutputStreamPrinter) reportInfo(e *Error) error {
	return nil
}
