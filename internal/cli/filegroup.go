package cli

import (
	"os"

	"github.com/apachejuice/pretzel/internal/ast"
	"github.com/apachejuice/pretzel/internal/errors"
	"github.com/apachejuice/pretzel/internal/lexer"
	"github.com/apachejuice/pretzel/internal/parser"
)

// A FileGroup handles the lexing and parsing process for a predetermined set of files.
type FileGroup struct {
	// The compiler configuration
	Config *CompilerConfig
	// The generated RootNodes
	Output []*ast.RootNode
}

// Parse all files in the file group. Returns OK status and errors, in addition to the error outside of the parsing/lexing.
func (f *FileGroup) ParseAll() (bool, []*errors.Error, error) {
	err := f.Config.FindFiles()
	if err != nil {
		return false, []*errors.Error{}, err
	}

	for _, file := range f.Config.Files {
		content, err := os.ReadFile(file)
		if err != nil {
			return false, []*errors.Error{}, err
		}

		l := lexer.NewLexer(string(content), file, f.Config.ErrorLimit)
		if !l.GetTokens() {
			return false, l.Errors, nil
		}

		p := parser.NewParser(l)
		if !p.Parse() {
			return false, p.Errors, nil
		}

		f.Output = append(f.Output, p.Root)
	}

	return true, []*errors.Error{}, nil
}
