package ast

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/lexer"
)

type (
	// Literal type enum
	LiteralType int

	// A literal of some kind - an integer, a string etc
	Literal interface {
		Expression

		// The type of the literal
		LType() LiteralType
		// The string value of the literal
		Value() string
	}

	// A base literal all literals should embed.
	baseLiteral struct {
		baseExpression
		baseNode

		// The literal token
		token *lexer.Token
	}

	// A decimal number literal.
	DecimalLiteral struct{ baseLiteral }
	// A hexadecimal number literal.
	HexadecimalLiteral struct{ baseLiteral }
	// A binary number literal.
	BinaryLiteral struct{ baseLiteral }
	// An octal number literal.
	OctalLiteral struct{ baseLiteral }
	// A roman number literal
	RomanLiteral struct{ baseLiteral }
	// A string literal
	StringLiteral struct{ baseLiteral }
	// A template string literal
	TemplateStringLiteral struct{ baseLiteral }
	// A boolean literal
	BooleanLiteral struct{ baseLiteral }
)

func (b baseLiteral) Value() string {
	return b.token.Text
}

func (baseLiteral) Children() []Node {
	return []Node{}
}

func (baseLiteral) EType() ExpressionType {
	return ExpressionLiteral
}

func (baseLiteral) HasOperator() bool {
	return false
}

func (baseLiteral) IsPure() bool {
	return true
}

func (baseLiteral) Operands() int {
	return 1
}

var _ Literal = &DecimalLiteral{}
var _ Literal = &HexadecimalLiteral{}
var _ Literal = &BinaryLiteral{}
var _ Literal = &OctalLiteral{}
var _ Literal = &RomanLiteral{}
var _ Literal = &StringLiteral{}
var _ Literal = &TemplateStringLiteral{}
var _ Literal = &BooleanLiteral{}

func (d *DecimalLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(d)
	l.EnterExpression(d)
	l.EnterLiteral(d)
}

func (d *DecimalLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(d)
	l.ExitExpression(d)
	l.ExitNode(d)
}

func (DecimalLiteral) IsConstant() bool {
	return true
}

func (DecimalLiteral) LType() LiteralType {
	return LiteralDecimal
}

func (d DecimalLiteral) String() string {
	return fmt.Sprintf("DecimalLiteral(%s)", d.token.Text)
}

func (d *DecimalLiteral) Traverse(l Listener) {
	l.EnterDecimalLiteral(d)
	l.ExitDecimalLiteral(d)
}

func (h *HexadecimalLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(h)
	l.EnterExpression(h)
	l.EnterLiteral(h)
}

func (h *HexadecimalLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(h)
	l.ExitExpression(h)
	l.ExitNode(h)
}

func (HexadecimalLiteral) IsConstant() bool {
	return true
}

func (HexadecimalLiteral) LType() LiteralType {
	return LiteralHexadecimal
}

func (h HexadecimalLiteral) String() string {
	return fmt.Sprintf("HexadecimalLiteral(%s)", h.token.Text)
}

func (h *HexadecimalLiteral) Traverse(l Listener) {
	l.EnterHexadecimalLiteral(h)
	l.ExitHexadecimalLiteral(h)
}

func (b *BinaryLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
	l.EnterLiteral(b)
}

func (b *BinaryLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(b)
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (BinaryLiteral) IsConstant() bool {
	return true
}

func (BinaryLiteral) LType() LiteralType {
	return LiteralBinary
}

func (b BinaryLiteral) String() string {
	return fmt.Sprintf("BinaryLiteral(%s)", b.token.Text)
}

func (b *BinaryLiteral) Traverse(l Listener) {
	l.EnterBinaryLiteral(b)
	l.ExitBinaryLiteral(b)
}

func (b *OctalLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
	l.EnterLiteral(b)
}

func (b *OctalLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(b)
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (OctalLiteral) IsConstant() bool {
	return true
}

func (OctalLiteral) LType() LiteralType {
	return LiteralOctal
}

func (b OctalLiteral) String() string {
	return fmt.Sprintf("OctalLiteral(%s)", b.token.Text)
}

func (b *OctalLiteral) Traverse(l Listener) {
	l.EnterOctalLiteral(b)
	l.ExitOctalLiteral(b)
}

func (b *RomanLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
	l.EnterLiteral(b)
}

func (b *RomanLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(b)
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (RomanLiteral) IsConstant() bool {
	return true
}

func (RomanLiteral) LType() LiteralType {
	return LiteralRoman
}

func (b RomanLiteral) String() string {
	return fmt.Sprintf("RomanLiteral(%s)", b.token.Text)
}

func (b *RomanLiteral) Traverse(l Listener) {
	l.EnterRomanLiteral(b)
	l.ExitRomanLiteral(b)
}

func (b *StringLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
	l.EnterLiteral(b)
}

func (b *StringLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(b)
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (StringLiteral) IsConstant() bool {
	return true
}

func (StringLiteral) LType() LiteralType {
	return LiteralString
}

func (b StringLiteral) String() string {
	return fmt.Sprintf("StringLiteral(%s)", b.token.Text)
}

func (b *StringLiteral) Traverse(l Listener) {
	l.EnterStringLiteral(b)
	l.ExitStringLiteral(b)
}

func (b *TemplateStringLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
	l.EnterLiteral(b)
}

func (b *TemplateStringLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(b)
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (TemplateStringLiteral) IsConstant() bool {
	return false
}

func (TemplateStringLiteral) LType() LiteralType {
	return LiteralTemplateString
}

func (b TemplateStringLiteral) String() string {
	return fmt.Sprintf("TemplateStringLiteral(%s)", b.token.Text)
}

func (b *TemplateStringLiteral) Traverse(l Listener) {
	l.EnterTemplateStringLiteral(b)
	l.ExitTemplateStringLiteral(b)
}

func (b *BooleanLiteral) EnterSuperclass(l Listener) {
	l.EnterNode(b)
	l.EnterExpression(b)
	l.EnterLiteral(b)
}

func (b *BooleanLiteral) ExitSuperclass(l Listener) {
	l.ExitLiteral(b)
	l.ExitExpression(b)
	l.ExitNode(b)
}

func (BooleanLiteral) IsConstant() bool {
	return true
}

func (BooleanLiteral) LType() LiteralType {
	return LiteralBoolean
}

func (b BooleanLiteral) String() string {
	return fmt.Sprintf("BooleanLiteral(%s)", b.token.Text)
}

func (b *BooleanLiteral) Traverse(l Listener) {
	l.EnterBooleanLiteral(b)
	l.ExitBooleanLiteral(b)
}

const (
	LiteralDecimal LiteralType = iota
	LiteralHexadecimal
	LiteralBinary
	LiteralOctal
	LiteralRoman
	LiteralString
	LiteralTemplateString
	LiteralBoolean
)
