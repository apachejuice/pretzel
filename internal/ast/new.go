package ast

import (
	"github.com/apachejuice/pretzel/internal/errors"
	"github.com/apachejuice/pretzel/internal/lexer"
)

// Creation methods for all nodes.

func NewRootNode(elements []Node, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *RootNode {
	return &RootNode{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		elements: elements,
	}
}

func NewName(name string, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *Name {
	return &Name{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Name: name,
	}
}

func NewQualifiedName(names []*Name, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *QualifiedName {
	return &QualifiedName{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Names: names,
	}
}

func NewUseStatement(imports *QualifiedName, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *UseStatement {
	return &UseStatement{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Imports: imports,
	}
}

func NewFunctionArg(name *QualifiedName, argType DataType, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *FunctionArg {
	return &FunctionArg{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Name:    name,
		ArgType: argType,
	}
}

func NewFunction(name *QualifiedName, args []*FunctionArg, body *Block, returnType DataType, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *Function {
	return &Function{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Name:       name,
		Args:       args,
		Body:       body,
		ReturnType: returnType,
	}
}

func NewNoneType(start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *NoneType {
	return &NoneType{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
	}
}

func NewArrayType(inner DataType, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *ArrayType {
	return &ArrayType{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Inner: inner,
	}
}

func NewGenericType(container *QualifiedName, inners []DataType, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *GenericType {
	return &GenericType{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Reference: container,
		Args:      inners,
	}
}

func NewAtomType(name *QualifiedName, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *AtomType {
	return &AtomType{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Reference: name,
	}
}

func NewNullableType(inner DataType, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *NullableType {
	return &NullableType{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Inner: inner,
	}
}

func NewBlock(nodes []Statement, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *Block {
	return &Block{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Nodes: nodes,
	}
}

func baselit(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) baseLiteral {
	return baseLiteral{
		baseNode: baseNode{
			start:     token.Context,
			end:       token.Context,
			reportCtx: reportCtx,
			scope:     scope,
		},
		token: token,
	}
}

func NewDecimalLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *DecimalLiteral {
	return &DecimalLiteral{baselit(token, reportCtx, scope)}
}

func NewHexadecimalLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *HexadecimalLiteral {
	return &HexadecimalLiteral{baselit(token, reportCtx, scope)}
}

func NewBinaryLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *BinaryLiteral {
	return &BinaryLiteral{baselit(token, reportCtx, scope)}
}

func NewOctalLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *OctalLiteral {
	return &OctalLiteral{baselit(token, reportCtx, scope)}
}

func NewRomanLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *RomanLiteral {
	return &RomanLiteral{baselit(token, reportCtx, scope)}
}

func NewStringLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *StringLiteral {
	return &StringLiteral{baselit(token, reportCtx, scope)}
}

func NewTemplateStringLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *TemplateStringLiteral {
	return &TemplateStringLiteral{baselit(token, reportCtx, scope)}
}

func NewBooleanLiteral(token *lexer.Token, reportCtx *errors.ReportContext, scope *Scope) *BooleanLiteral {
	return &BooleanLiteral{baselit(token, reportCtx, scope)}
}

func NewBinaryExpression(left, right Expression, operator BinaryOperator, reportCtx *errors.ReportContext, scope *Scope) *BinaryExpression {
	return &BinaryExpression{
		baseNode: baseNode{
			start:     left.Begin(),
			end:       right.End(),
			reportCtx: reportCtx,
			scope:     scope,
		},
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}

func NewVariableReference(name *QualifiedName, reportCtx *errors.ReportContext, scope *Scope) *VariableReference {
	return &VariableReference{
		baseNode: baseNode{
			start:     name.start,
			end:       name.end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Reference: name,
	}
}

func NewInvalidExpression(start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *InvalidExpression {
	return &InvalidExpression{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
	}
}

func NewParenExpression(inner Expression, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *ParenExpression {
	return &ParenExpression{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Inner: inner,
	}
}

func NewPrefixExpression(inner Expression, operator PrefixOperator, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *PrefixExpression {
	return &PrefixExpression{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Inner:    inner,
		Operator: operator,
	}
}

func NewPostfixExpression(inner Expression, operator PostfixOperator, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *PostfixExpression {
	return &PostfixExpression{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Inner:    inner,
		Operator: operator,
	}
}

func NewCallExpression(callee Expression, args []Expression, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *CallExpression {
	return &CallExpression{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Callee: callee,
		Args:   args,
	}
}

func NewSubscriptExpression(source, arg Expression, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *SubscriptExpression {
	return &SubscriptExpression{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Source: source,
		Arg:    arg,
	}
}

func NewExpressionStatement(inner Expression, reportCtx *errors.ReportContext, scope *Scope) *ExpressionStatement {
	return &ExpressionStatement{
		baseNode: baseNode{
			start:     inner.Begin(),
			end:       inner.End(),
			reportCtx: reportCtx,
			scope:     scope,
		},
		Inner: inner,
	}
}

func NewVariableDeclaration(name *QualifiedName, dataType DataType, init Expression, start, end errors.SourceContext, reportCtx *errors.ReportContext, scope *Scope) *VariableDeclaration {
	return &VariableDeclaration{
		baseNode: baseNode{
			start:     start,
			end:       end,
			reportCtx: reportCtx,
			scope:     scope,
		},
		Name:         name,
		DeclaredType: dataType,
		Init:         init,
	}
}
