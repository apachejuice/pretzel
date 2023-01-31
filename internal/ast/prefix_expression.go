package ast

import "fmt"

type (
	PrefixOperator int

	// Prefix expressions
	PrefixExpression struct {
		baseNode
		baseExpression

		// The operator
		Operator PrefixOperator
		// The expression the operator applies to
		Inner Expression
	}
)

var _ Expression = &PrefixExpression{}

func (p PrefixExpression) Children() []Node {
	return []Node{p.Inner}
}

func (PrefixExpression) EType() ExpressionType {
	return ExpressionPrefix
}

func (p *PrefixExpression) EnterSuperclass(l Listener) {
	l.EnterNode(p)
	l.EnterExpression(p)
}

func (p *PrefixExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(p)
	l.ExitNode(p)
}

func (p *PrefixExpression) Traverse(l Listener) {
	l.EnterPrefixExpression(p)
	p.Inner.Traverse(l)
	l.ExitPrefixExpression(p)
}

func (PrefixExpression) HasOperator() bool {
	return true
}

func (p PrefixExpression) IsConstant() bool {
	return p.Inner.IsConstant()
}

func (p PrefixExpression) IsPure() bool {
	return p.Operator.IsPure()
}

func (PrefixExpression) IsStatementLike() bool {
	return true
}

func (p PrefixExpression) Operands() int {
	return 1
}

func (p PrefixExpression) String() string {
	return fmt.Sprintf("PrefixExpression(%s %s)", p.Operator, p.Inner)
}

const (
	PrefixOperatorPlus   PrefixOperator = iota // +
	PrefixOperatorMinus                        // -
	PrefixOperatorNot                          // !
	PrefixOperatorNegate                       // ~
	PrefixOperatorInc                          // ++
	PrefixOperatorDec                          // --

	PrefixOperatorInvalid // <invalid>
)

var _ fmt.Stringer = PrefixOperator(0)

func (p PrefixOperator) String() string {
	return prefixOps[p]
}

func (p PrefixOperator) IsPure() bool {
	return p != PrefixOperatorInc && p != PrefixOperatorDec
}

var prefixOps map[PrefixOperator]string = map[PrefixOperator]string{
	PrefixOperatorPlus:   "+",
	PrefixOperatorMinus:  "-",
	PrefixOperatorNot:    "!",
	PrefixOperatorNegate: "~",
	PrefixOperatorInc:    "++",
	PrefixOperatorDec:    "--",

	PrefixOperatorInvalid: "<invalid>",
}
