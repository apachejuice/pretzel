package ast

import "fmt"

// Represents a return expression.
type ReturnExpression struct {
	baseNode
	baseExpression

	// The returned value, may be nil.
	Value Expression
}

var _ Expression = &ReturnExpression{}

func (r *ReturnExpression) EType() ExpressionType {
	return ExpressionReturn
}

func (r *ReturnExpression) Operands() int {
	return 1
}

func (r *ReturnExpression) HasOperator() bool {
	return false
}

func (r *ReturnExpression) IsConstant() bool {
	return true // as an expression, never returns anything
}

func (r ReturnExpression) IsPure() bool {
	return r.Value.IsPure()
}

func (ReturnExpression) IsStatementLike() bool {
	return true
}

func (r ReturnExpression) String() string {
	return fmt.Sprintf("ReturnExpression(%s)", r.Value)
}

func (r ReturnExpression) Children() []Node {
	return []Node{r.Value}
}

func (r *ReturnExpression) EnterSuperclass(l Listener) {
	l.EnterNode(r)
	l.EnterExpression(r)
}

func (r *ReturnExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(r)
	l.ExitNode(r)
}

func (r *ReturnExpression) Traverse(l Listener) {
	l.EnterReturnExpression(r)
	r.Value.Traverse(l)
	l.ExitReturnExpression(r)
}
