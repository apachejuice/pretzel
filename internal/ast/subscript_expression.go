package ast

import (
	"fmt"
)

// Represents a call expression.
type SubscriptExpression struct {
	baseNode
	baseExpression

	// The expression being subscripted
	Source Expression
	// The argument to the subscript
	Arg Expression
}

var _ Expression = &SubscriptExpression{}

func (c SubscriptExpression) Children() []Node {
	return []Node{c.Source, c.Arg}
}

func (SubscriptExpression) EType() ExpressionType {
	return ExpressionSubscript
}

func (c *SubscriptExpression) EnterSuperclass(l Listener) {
	l.EnterNode(c)
	l.EnterExpression(c)
}

func (c *SubscriptExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(c)
	l.ExitNode(c)
}

func (c *SubscriptExpression) Traverse(l Listener) {
	l.EnterSubscriptExpression(c)
	c.Source.Traverse(l)
	c.Arg.Traverse(l)

	l.ExitSubscriptExpression(c)
}

func (c SubscriptExpression) String() string {
	return fmt.Sprintf("SubscriptExpression(%s %s)", c.Source, c.Arg)
}

func (c SubscriptExpression) HasOperator() bool {
	return false
}

func (c SubscriptExpression) IsConstant() bool {
	return false
}

func (c SubscriptExpression) IsPure() bool {
	return true
}

func (c SubscriptExpression) Operands() int {
	return 2
}
