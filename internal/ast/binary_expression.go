package ast

import "fmt"

type (
	// A binary expression operator
	BinaryOperator int

	// A binary expression type holding two operands and an operator in the form <operand1> <operator> <operand2>.
	BinaryExpression struct {
		baseExpression
		baseNode

		// The left side operand
		Left Expression
		// The operator
		Operator BinaryOperator
		// The right side operand
		Right Expression
	}
)

var _ Expression = &BinaryExpression{}

func (b BinaryExpression) Children() []Node {
	return []Node{b.Left, b.Right}
}

func (b BinaryExpression) EType() ExpressionType {
	return ExpressionBinary
}

func (b *BinaryExpression) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
}

func (b *BinaryExpression) ExitSuperclass(l Listener) {
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (b *BinaryExpression) Traverse(l Listener) {
	l.EnterBinaryExpression(b)
	b.Left.Traverse(l)
	b.Right.Traverse(l)
	l.ExitBinaryExpression(b)
}

func (BinaryExpression) HasOperator() bool {
	return true
}

func (b BinaryExpression) IsConstant() bool {
	return b.Left.IsConstant() && b.Right.IsConstant()
}

func (b BinaryExpression) IsPure() bool {
	return b.Left.IsPure() && b.Right.IsPure()
}

func (b BinaryExpression) IsStatementLike() bool {
	return b.Operator >= BinaryOperatorAssign && b.Operator <= BinaryOperatorRightShiftAssign
}

func (BinaryExpression) Operands() int {
	return 2
}

func (b BinaryExpression) String() string {
	return fmt.Sprintf("BinaryExpression(%s %s %s)", b.Left, b.Operator, b.Right)
}

const (
	BinaryOperatorEquals          BinaryOperator = iota // ==
	BinaryOperatorNotEquals                             // !=
	BinaryOperatorLessThanOrEqual                       // <=
	BinaryOperatorMoreThanOrEqual                       // >=
	BinaryOperatorOr                                    // ||
	BinaryOperatorPow                                   // **
	BinaryOperatorAnd                                   // &&
	BinaryOperatorXor                                   // ^^
	BinaryOperatorChainAccess                           // ..
	BinaryOperatorLessThan                              // <
	BinaryOperatorMoreThan                              // >
	BinaryOperatorBitOr                                 // |
	BinaryOperatorMod                                   // %
	BinaryOperatorBitXor                                // ^
	BinaryOperatorBitAnd                                // &
	BinaryOperatorMinus                                 // -
	BinaryOperatorPlus                                  // +
	BinaryOperatorMul                                   // *
	BinaryOperatorDiv                                   // /
	BinaryOperatorAccess                                // .
	BinaryOperatorAs                                    // as
	BinaryOperatorIs                                    // is
	BinaryOperatorTo                                    // to
	BinaryOperatorIn                                    // in
	BinaryOperatorLeftShift                             // <<
	BinaryOperatorRightShift                            // >>

	BinaryOperatorAssign           // =
	BinaryOperatorDivAssign        // /=
	BinaryOperatorMulAssign        // *=
	BinaryOperatorPlusAssign       // +=
	BinaryOperatorMinusAssign      // -=
	BinaryOperatorBitAndAssign     // &=
	BinaryOperatorBitXorAssign     // ^=
	BinaryOperatorModAssign        // %=
	BinaryOperatorBitOrAssign      // |=
	BinaryOperatorXorAssign        // ^^=
	BinaryOperatorAndAssign        // &&=
	BinaryOperatorPowAssign        // **=
	BinaryOperatorOrAssign         // ||=
	BinaryOperatorLeftShiftAssign  // <<=
	BinaryOperatorRightShiftAssign // >>=

	BinaryOperatorInvalid // <invalid>
)

var _ fmt.Stringer = BinaryOperator(0)
var binaryOps map[BinaryOperator]string = map[BinaryOperator]string{
	BinaryOperatorEquals:          "==",
	BinaryOperatorNotEquals:       "!=",
	BinaryOperatorLessThanOrEqual: "<=",
	BinaryOperatorMoreThanOrEqual: ">=",
	BinaryOperatorOr:              "||",
	BinaryOperatorPow:             "**",
	BinaryOperatorAnd:             "&&",
	BinaryOperatorXor:             "^^",
	BinaryOperatorChainAccess:     "..",
	BinaryOperatorLessThan:        "<",
	BinaryOperatorMoreThan:        ">",
	BinaryOperatorBitOr:           "|",
	BinaryOperatorMod:             "%",
	BinaryOperatorBitXor:          "^",
	BinaryOperatorBitAnd:          "&",
	BinaryOperatorMinus:           "-",
	BinaryOperatorPlus:            "+",
	BinaryOperatorMul:             "*",
	BinaryOperatorDiv:             "/",
	BinaryOperatorAccess:          ".",
	BinaryOperatorAs:              "as",
	BinaryOperatorIs:              "is",
	BinaryOperatorTo:              "to",
	BinaryOperatorIn:              "in",

	BinaryOperatorAssign:           "=",
	BinaryOperatorDivAssign:        "/=",
	BinaryOperatorMulAssign:        "*=",
	BinaryOperatorPlusAssign:       "+=",
	BinaryOperatorMinusAssign:      "-=",
	BinaryOperatorBitAndAssign:     "&=",
	BinaryOperatorBitXorAssign:     "^=",
	BinaryOperatorModAssign:        "%=",
	BinaryOperatorBitOrAssign:      "|=",
	BinaryOperatorXorAssign:        "^^=",
	BinaryOperatorAndAssign:        "&&=",
	BinaryOperatorPowAssign:        "**=",
	BinaryOperatorOrAssign:         "||=",
	BinaryOperatorLeftShiftAssign:  "<<=",
	BinaryOperatorRightShiftAssign: ">>=",

	BinaryOperatorInvalid: "<invalid>",
}

func (b BinaryOperator) String() string {
	return binaryOps[b]
}
