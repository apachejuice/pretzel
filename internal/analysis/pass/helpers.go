package pass

import (
	"github.com/apachejuice/pretzel/internal/ast"
	"github.com/apachejuice/pretzel/internal/errors"
)

func errorNode(node ast.Node, code errors.ErrorCode, args ...any) *errors.Error {
	return errors.NewError(code, node.Begin(), node.End(), node.ReportContext(), args...)
}

func varName(q *ast.QualifiedName) string {
	name := q.Names[0].Name
	for i := 1; i < len(q.Names); i++ {
		name += "." + q.Names[i].Name
	}

	return name
}
