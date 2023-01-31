package ast

import "fmt"

// An expression statement encompasses an expression into a statement body.
type ExpressionStatement struct {
	baseNode

	// The inner expression
	Inner Expression
}

var _ Statement = &ExpressionStatement{}

func (ExpressionStatement) IsControlFlow() bool {
	return false
}

func (ExpressionStatement) Type() NodeType {
	return NodeExpressionStatement
}

func (e ExpressionStatement) Children() []Node {
	return []Node{e.Inner}
}

func (e *ExpressionStatement) EnterSuperclass(l Listener) {
	l.EnterNode(e)
	l.EnterStatement(e)
}

func (e *ExpressionStatement) ExitSuperclass(l Listener) {
	l.ExitStatement(e)
	l.ExitNode(e)
}

func (e *ExpressionStatement) Traverse(l Listener) {
	l.EnterExpressionStatement(e)
	e.Inner.Traverse(l)
	l.ExitExpressionStatement(e)
}

func (e ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement(%s)", e.Inner)
}
