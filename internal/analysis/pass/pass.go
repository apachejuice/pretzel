//lint:file-ignore U1000 Ignore all unused code, empty implementation.
package pass

import (
	"github.com/apachejuice/pretzel/internal/ast"
	"github.com/apachejuice/pretzel/internal/errors"
)

// A pass is an implementation of a Listener that does not necessarily implement everything, so some methods are left empty.
type pass struct {
	ast.BaseListener
	errors []*errors.Error
}

// An ast.Listener with an Errors() method.
type Pass interface {
	ast.Listener

	// The errors found during this pass.
	Errors() []*errors.Error
}

var _ ast.Listener = pass{}

func newPass() pass {
	return pass{
		BaseListener: ast.NewBaseListener(),
		errors:       make([]*errors.Error, 0),
	}
}

func (pass pass) EnterNode(n ast.Node) {}

func (pass pass) ExitNode(n ast.Node) {}

func (pass pass) EnterStatement(s ast.Statement) {}

func (pass pass) ExitStatement(s ast.Statement) {}

func (pass pass) EnterDataType(d ast.DataType) {}

func (pass pass) ExitDataType(d ast.DataType) {}

func (pass pass) EnterExpression(e ast.Expression) {}

func (pass pass) ExitExpression(e ast.Expression) {}

func (pass pass) EnterLiteral(l ast.Literal) {}

func (pass pass) ExitLiteral(l ast.Literal) {}

func (pass pass) EnterRootNode(r *ast.RootNode) {}

func (pass pass) ExitRootNode(r *ast.RootNode) {}

func (pass pass) EnterUseStatement(u *ast.UseStatement) {}

func (pass pass) ExitUseStatement(u *ast.UseStatement) {}

func (pass pass) EnterName(n *ast.Name) {}

func (pass pass) ExitName(n *ast.Name) {}

func (pass pass) EnterQualifiedName(q *ast.QualifiedName) {}

func (pass pass) ExitQualifiedName(q *ast.QualifiedName) {}

func (pass pass) EnterAtomType(a *ast.AtomType) {}

func (pass pass) ExitAtomType(a *ast.AtomType) {}

func (pass pass) EnterArrayType(a *ast.ArrayType) {}

func (pass pass) ExitArrayType(a *ast.ArrayType) {}

func (pass pass) EnterGenericType(g *ast.GenericType) {}

func (pass pass) ExitGenericType(g *ast.GenericType) {}

func (pass pass) EnterNullableType(n *ast.NullableType) {}

func (pass pass) ExitNullableType(n *ast.NullableType) {}

func (pass pass) EnterNoneType(n *ast.NoneType) {}

func (pass pass) ExitNoneType(n *ast.NoneType) {}

func (pass pass) EnterFunctionArg(f *ast.FunctionArg) {}

func (pass pass) ExitFunctionArg(f *ast.FunctionArg) {}

func (pass pass) EnterFunction(f *ast.Function) {}

func (pass pass) ExitFunction(f *ast.Function) {}

func (pass pass) EnterBlock(b *ast.Block) {}

func (pass pass) ExitBlock(b *ast.Block) {}

func (pass pass) EnterDecimalLiteral(d *ast.DecimalLiteral) {}

func (pass pass) ExitDecimalLiteral(d *ast.DecimalLiteral) {}

func (pass pass) EnterHexadecimalLiteral(h *ast.HexadecimalLiteral) {}

func (pass pass) ExitHexadecimalLiteral(h *ast.HexadecimalLiteral) {}

func (pass pass) EnterBinaryLiteral(b *ast.BinaryLiteral) {}

func (pass pass) ExitBinaryLiteral(b *ast.BinaryLiteral) {}

func (pass pass) EnterOctalLiteral(o *ast.OctalLiteral) {}

func (pass pass) ExitOctalLiteral(o *ast.OctalLiteral) {}

func (pass pass) EnterRomanLiteral(r *ast.RomanLiteral) {}

func (pass pass) ExitRomanLiteral(r *ast.RomanLiteral) {}

func (pass pass) EnterStringLiteral(s *ast.StringLiteral) {}

func (pass pass) ExitStringLiteral(s *ast.StringLiteral) {}

func (pass pass) EnterTemplateStringLiteral(t *ast.TemplateStringLiteral) {}

func (pass pass) ExitTemplateStringLiteral(t *ast.TemplateStringLiteral) {}

func (pass pass) EnterBooleanLiteral(b *ast.BooleanLiteral) {}

func (pass pass) ExitBooleanLiteral(b *ast.BooleanLiteral) {}

func (pass pass) EnterBinaryExpression(b *ast.BinaryExpression) {}

func (pass pass) ExitBinaryExpression(b *ast.BinaryExpression) {}

func (pass pass) EnterVariableReference(v *ast.VariableReference) {}

func (pass pass) ExitVariableReference(v *ast.VariableReference) {}

func (pass pass) EnterParenExpression(p *ast.ParenExpression) {}

func (pass pass) ExitParenExpression(p *ast.ParenExpression) {}

func (pass pass) EnterPrefixExpression(p *ast.PrefixExpression) {}

func (pass pass) ExitPrefixExpression(p *ast.PrefixExpression) {}

func (pass pass) EnterPostfixExpression(p *ast.PostfixExpression) {}

func (pass pass) ExitPostfixExpression(p *ast.PostfixExpression) {}

func (pass pass) EnterCallExpression(c *ast.CallExpression) {}

func (pass pass) ExitCallExpression(c *ast.CallExpression) {}

func (pass pass) EnterSubscriptExpression(s *ast.SubscriptExpression) {}

func (pass pass) ExitSubscriptExpression(s *ast.SubscriptExpression) {}

func (pass pass) EnterExpressionStatement(e *ast.ExpressionStatement) {}

func (pass pass) ExitExpressionStatement(e *ast.ExpressionStatement) {}

func (pass pass) EnterReturnExpression(r *ast.ReturnExpression) {}

func (pass pass) ExitReturnExpression(r *ast.ReturnExpression) {}

func (pass pass) EnterVariableDeclaration(v *ast.VariableDeclaration) {}

func (pass pass) ExitVariableDeclaration(v *ast.VariableDeclaration) {}
