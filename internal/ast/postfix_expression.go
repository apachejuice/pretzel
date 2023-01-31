package ast

import "fmt"

type (
	PostfixOperator int

	// Postfix expressions
	PostfixExpression struct {
		baseNode
		baseExpression

		// The operator
		Operator PostfixOperator
		// The expression the operator applies to
		Inner Expression
	}
)

var _ Expression = &PostfixExpression{}

func (p PostfixExpression) Children() []Node {
	return []Node{p.Inner}
}

func (PostfixExpression) EType() ExpressionType {
	return ExpressionPostfix
}

func (p *PostfixExpression) EnterSuperclass(l Listener) {
	l.EnterNode(p)
	l.EnterExpression(p)
}

func (p *PostfixExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(p)
	l.ExitNode(p)
}

func (p *PostfixExpression) Traverse(l Listener) {
	l.EnterPostfixExpression(p)
	p.Inner.Traverse(l)
	l.ExitPostfixExpression(p)
}

func (PostfixExpression) HasOperator() bool {
	return true
}

func (p PostfixExpression) IsConstant() bool {
	return p.Inner.IsConstant()
}

func (p PostfixExpression) IsPure() bool {
	return false
}

func (PostfixExpression) IsStatementLike() bool {
	return true
}

func (p PostfixExpression) Operands() int {
	return 1
}

func (p PostfixExpression) String() string {
	return fmt.Sprintf("PostfixExpression(%s %s)", p.Operator, p.Inner)
}

const (
	PostfixOperatorInc PostfixOperator = iota // ++
	PostfixOperatorDec                        // --

	PostfixOperatorInvalid // <invalid>
)

var _ fmt.Stringer = PostfixOperator(0)

func (p PostfixOperator) String() string {
	return postfixOps[p]
}

var postfixOps map[PostfixOperator]string = map[PostfixOperator]string{
	PostfixOperatorInc: "++",
	PostfixOperatorDec: "--",

	PostfixOperatorInvalid: "<invalid>",
}
