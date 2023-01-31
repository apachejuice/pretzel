package pass

import (
	"github.com/apachejuice/pretzel/internal/ast"
	"github.com/apachejuice/pretzel/internal/errors"
)

// A combo pass takes multiple passes and executes them in the given order.
type ComboPass struct {
	// Errors from all the passes
	Errors []*errors.Error
	// the passes
	passes []Pass
}

func NewComboPass(passes ...Pass) *ComboPass {
	return &ComboPass{
		Errors: make([]*errors.Error, 0),
		passes: passes,
	}
}

func (c *ComboPass) Run(node ast.Node) {
	for _, pass := range c.passes {
		node.Traverse(pass)
		c.Errors = append(c.Errors, pass.Errors()...)
	}
}
