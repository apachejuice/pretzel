package ast

import "fmt"

// A parenthesized expression.
type ParenExpression struct {
	baseNode
	baseExpression

	// The inner expression
	Inner Expression
}

var _ Expression = &ParenExpression{}

func (p ParenExpression) Children() []Node {
	return []Node{p.Inner}
}

func (ParenExpression) EType() ExpressionType {
	return ExpressionParen
}

func (p *ParenExpression) EnterSuperclass(l Listener) {
	l.EnterNode(p)
	l.EnterExpression(p)
}

func (p *ParenExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(p)
	l.ExitNode(p)
}

func (p *ParenExpression) Traverse(l Listener) {
	l.EnterParenExpression(p)
	p.Inner.Traverse(l)
	l.ExitParenExpression(p)
}

func (ParenExpression) HasOperator() bool {
	return false
}

func (p ParenExpression) IsConstant() bool {
	return p.Inner.IsConstant()
}

func (p ParenExpression) IsPure() bool {
	return p.Inner.IsPure()
}

func (p ParenExpression) IsStatementLike() bool {
	return p.Inner.IsStatementLike()
}

func (ParenExpression) Operands() int {
	return 1
}

func (p ParenExpression) String() string {
	return fmt.Sprintf("ParenExpression(%s)", p.Inner)
}
