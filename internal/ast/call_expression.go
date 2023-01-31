package ast

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/common/util"
)

// Represents a call expression.
type CallExpression struct {
	baseNode
	baseExpression

	// The expression being called
	Callee Expression
	// The arguments to the call
	Args []Expression
}

var _ Expression = &CallExpression{}

func (c CallExpression) Children() []Node {
	return append([]Node{c.Callee}, util.ArrayTransform[Expression, Node](c.Args)...)
}

func (CallExpression) EType() ExpressionType {
	return ExpressionCall
}

func (c *CallExpression) EnterSuperclass(l Listener) {
	l.EnterNode(c)
	l.EnterExpression(c)
}

func (c *CallExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(c)
	l.ExitNode(c)
}

func (c *CallExpression) Traverse(l Listener) {
	l.EnterCallExpression(c)
	c.Callee.Traverse(l)
	for _, arg := range c.Args {
		arg.Traverse(l)
	}

	l.ExitCallExpression(c)
}

func (c CallExpression) String() string {
	return fmt.Sprintf("CallExpression(%s %s)", c.Callee, nodeArrayString(c.Args))
}

func (c CallExpression) HasOperator() bool {
	return false
}

func (c CallExpression) IsConstant() bool {
	return false
}

func (c CallExpression) IsPure() bool {
	return false
}

func (CallExpression) IsStatementLike() bool {
	return true
}

func (c CallExpression) Operands() int {
	return len(c.Args) + 1
}
