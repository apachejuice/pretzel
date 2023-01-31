package ast

import "github.com/apachejuice/pretzel/internal/common/constraint"

type (
	// A hook called upon an Enter or Exit function.
	ListenerHook func(n Node)

	// Represents a listener that traverses through nodes.
	Listener interface {
		// Register a hook called by entering any node.
		RegisterEnterHook(id string, hook ListenerHook)
		// Register a hook called by exiting any node.
		RegisterExitHook(id string, hook ListenerHook)

		/* The basic Node */

		// Called when entering every single type of node.
		EnterNode(n Node)
		// Called when exiting every single type of node.
		ExitNode(n Node)

		/* Abstract classes - interfaces OR structs that embed baseNode but do not implement everything.
		These are necessarily not Nodes due to Go's weird inheritance issues. */

		// Called when entering a statement.
		EnterStatement(s Statement)
		// Called when exiting a statement.
		ExitStatement(s Statement)

		// Called when entering a data type.
		EnterDataType(d DataType)
		// Called when exiting a data type.
		ExitDataType(d DataType)

		// Called when entering an expression
		EnterExpression(e Expression)
		// Called when exiting an expression
		ExitExpression(e Expression)

		// Called when entering a literal expression
		EnterLiteral(l Literal)
		// Called when exiting a literal
		ExitLiteral(l Literal)

		/* Concrete implementations */

		// Called when entering a root node.
		EnterRootNode(r *RootNode)
		// Called when exiting a root node.
		ExitRootNode(r *RootNode)

		// Called when entering a use statement.
		EnterUseStatement(u *UseStatement)
		// Called when exiting a use statement.
		ExitUseStatement(u *UseStatement)

		// Called when entering a name.
		EnterName(n *Name)
		// Called when exiting a name.
		ExitName(n *Name)

		// Called when entering a qualified name.
		EnterQualifiedName(q *QualifiedName)
		// Called when exiting a qualified name.
		ExitQualifiedName(q *QualifiedName)

		// Called when entering an atom type.
		EnterAtomType(a *AtomType)
		// Called when exiting an atom type.
		ExitAtomType(a *AtomType)

		// Called when entering an array type.
		EnterArrayType(a *ArrayType)
		// Called when exiting an array type.
		ExitArrayType(a *ArrayType)

		// Called when entering a generic type.
		EnterGenericType(g *GenericType)
		// Called when exiting a generic type.
		ExitGenericType(g *GenericType)

		// Called when entering a nullable type.
		EnterNullableType(n *NullableType)
		// Called when exiting a nullable type.
		ExitNullableType(n *NullableType)

		// Called when entering a none type.
		EnterNoneType(n *NoneType)
		// Called when exiting a none type.
		ExitNoneType(n *NoneType)

		// Called when entering a function argument.
		EnterFunctionArg(f *FunctionArg)
		// Called when exiting a function argument.
		ExitFunctionArg(f *FunctionArg)

		// Called when entering a function declaration.
		EnterFunction(f *Function)
		// Called when exiting a function declaration.
		ExitFunction(f *Function)

		// Called when entering a block.
		EnterBlock(b *Block)
		// Called when exiting a block.
		ExitBlock(b *Block)

		// Called when entering a decimal literal.
		EnterDecimalLiteral(d *DecimalLiteral)
		// Called when exiting a decimal literal.
		ExitDecimalLiteral(d *DecimalLiteral)

		// Called when entering a hexadecimal literal.
		EnterHexadecimalLiteral(h *HexadecimalLiteral)
		// Called when exiting a hexadecimal literal.
		ExitHexadecimalLiteral(h *HexadecimalLiteral)

		// Called when entering a binary literal.
		EnterBinaryLiteral(b *BinaryLiteral)
		// Called when exiting a binary literal.
		ExitBinaryLiteral(b *BinaryLiteral)

		// Called when entering an octal literal.
		EnterOctalLiteral(o *OctalLiteral)
		// Called when exiting an octal literal.
		ExitOctalLiteral(o *OctalLiteral)

		// Called when entering a roman literal.
		EnterRomanLiteral(r *RomanLiteral)
		// Called when exiting a roman literal.
		ExitRomanLiteral(r *RomanLiteral)

		// Called when entering a string literal.
		EnterStringLiteral(s *StringLiteral)
		// Called when exiting a string literal.
		ExitStringLiteral(s *StringLiteral)

		// Called when entering a template string literal.
		EnterTemplateStringLiteral(t *TemplateStringLiteral)
		// Called when exiting a template string literal.
		ExitTemplateStringLiteral(t *TemplateStringLiteral)

		// Called when entering a boolean literal.
		EnterBooleanLiteral(b *BooleanLiteral)
		// Called when exiting a boolean literal.
		ExitBooleanLiteral(b *BooleanLiteral)

		// Called when entering a binary expression.
		EnterBinaryExpression(b *BinaryExpression)
		// Called when exiting a binary expression.
		ExitBinaryExpression(b *BinaryExpression)

		// Called when entering a variable reference.
		EnterVariableReference(v *VariableReference)
		// Called when exiting a variable reference.
		ExitVariableReference(v *VariableReference)

		// Called when entering a parenthesized expression.
		EnterParenExpression(p *ParenExpression)
		// Called when exiting a parenthesized expression.
		ExitParenExpression(p *ParenExpression)

		// Called when entering a prefix expression.
		EnterPrefixExpression(p *PrefixExpression)
		// Called when exiting a prefix expression.
		ExitPrefixExpression(p *PrefixExpression)

		// Called when entering a postfix expression.
		EnterPostfixExpression(p *PostfixExpression)
		// Called when exiting a postfix expression.
		ExitPostfixExpression(p *PostfixExpression)

		// Called when entering a call expression.
		EnterCallExpression(c *CallExpression)
		// Called when exiting a call expression
		ExitCallExpression(c *CallExpression)

		// Called when entering a subscript expression.
		EnterSubscriptExpression(s *SubscriptExpression)
		// Called when exiting a subscript expression.
		ExitSubscriptExpression(s *SubscriptExpression)

		// Called when entering an expression statement
		EnterExpressionStatement(e *ExpressionStatement)
		// Called when exiting an expression statement
		ExitExpressionStatement(e *ExpressionStatement)

		// Called when entering a return expression
		EnterReturnExpression(r *ReturnExpression)
		// Called when exiting a return expression
		ExitReturnExpression(r *ReturnExpression)

		// Called when entering a variable declaration
		EnterVariableDeclaration(v *VariableDeclaration)
		// Called when exiting a variable declaration
		ExitVariableDeclaration(v *VariableDeclaration)
	}

	// A basic listener all listeners should embed.
	BaseListener struct {
		EnterHooks map[string]ListenerHook
		ExitHooks  map[string]ListenerHook
	}
)

// Use this when embedding
func NewBaseListener() BaseListener {
	return BaseListener{
		EnterHooks: make(map[string]ListenerHook),
		ExitHooks:  make(map[string]ListenerHook),
	}
}

func (b BaseListener) RegisterEnterHook(id string, hook ListenerHook) {
	constraint.NotNil("hook", hook)
	b.EnterHooks[id] = hook
}

func (b BaseListener) RegisterExitHook(id string, hook ListenerHook) {
	constraint.NotNil("hook", hook)
	b.ExitHooks[id] = hook
}
