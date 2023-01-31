package ast

type (
	// Expression type enum
	ExpressionType int

	// The basic expression interface.
	Expression interface {
		Node

		// The type of the expression.
		EType() ExpressionType
		// The amount of operands taken by the expression
		Operands() int
		// Does this expression have an operator?
		HasOperator() bool
		// Is this expression constant?
		IsConstant() bool
		// Is this expression pure?
		IsPure() bool
		// Is this a statement-like expression?
		IsStatementLike() bool
	}

	// A base expression all expressions should embed.
	baseExpression int

	// An invalid expression.
	InvalidExpression struct {
		baseNode
		baseExpression
	}
)

func (baseExpression) Type() NodeType {
	return NodeExpression
}

func (baseExpression) IsStatementLike() bool {
	return false
}

const (
	ExpressionLiteral ExpressionType = iota
	ExpressionBinary
	ExpressionVariableReference
	ExpressionParen
	ExpressionPrefix
	ExpressionPostfix
	ExpressionCall
	ExpressionSubscript
	ExpressionReturn

	ExpressionInvalid
)

var _ Expression = &InvalidExpression{}

func (InvalidExpression) Children() []Node           { return []Node{} }
func (InvalidExpression) EType() ExpressionType      { return ExpressionInvalid }
func (InvalidExpression) EnterSuperclass(l Listener) {}
func (InvalidExpression) ExitSuperclass(l Listener)  {}
func (InvalidExpression) HasOperator() bool          { return false }
func (InvalidExpression) IsConstant() bool           { return true }
func (InvalidExpression) IsPure() bool               { return true }
func (InvalidExpression) IsStatementLike() bool      { return false }
func (InvalidExpression) Operands() int              { return 0 }
func (InvalidExpression) String() string             { return "InvalidExpression()" }
func (InvalidExpression) Traverse(l Listener)        {}
