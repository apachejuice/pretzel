package parser

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/ast"
	"github.com/apachejuice/pretzel/internal/errors"
	"github.com/apachejuice/pretzel/internal/lexer"
	"golang.org/x/exp/slices"
)

// Parses the input into an abstract syntax tree.
// Basic error values are not used for errors. Use `ParserError`.
type Parser struct {
	curIdx    int            // current token index
	tokens    []*lexer.Token // the tokens in the input
	nodeTypes []ast.NodeType // a stack of node types
	reportCtx *errors.ReportContext
	rootScope *ast.Scope
	scopes    []*ast.Scope

	/* Error recovery information */
	currErrorCode errors.ErrorCode
	errStartCtx   errors.SourceContext

	// Errors generated during parsing
	Errors []*errors.Error
	// The root node
	Root *ast.RootNode
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		curIdx:        0,
		tokens:        l.Tokens,
		reportCtx:     l.ReportCtx,
		rootScope:     ast.NewScope(nil),
		scopes:        make([]*ast.Scope, 0),
		nodeTypes:     make([]ast.NodeType, 0),
		currErrorCode: 0,
		errStartCtx:   errors.SourceContext{},
		Errors:        make([]*errors.Error, 0),
		Root:          nil,
	}
}

/* Helper methods */

// Return the next token from the stack.
func (p Parser) peek() *lexer.Token {
	return p.tokens[p.curIdx]
}

// Return the current source context.
func (p Parser) ctx() errors.SourceContext {
	return p.peek().Context
}

// Are we at the end of the token list?
func (p Parser) atEnd() bool {
	return p.curIdx >= (len(p.tokens) - 1)
}

// Make an error from the given info and push it onto the error stack.
func (p *Parser) makeError(nodeType ast.NodeType, end errors.SourceContext, errorCode errors.ErrorCode) {
	p.Errors = append(p.Errors, errors.NewError(errorCode, p.ctx(), end, p.reportCtx))
}

// Start an error range, ready to sync if needed.
func (p *Parser) beginErr(errorCode errors.ErrorCode) {
	p.currErrorCode = errorCode
	p.errStartCtx = p.ctx()
}

// End an error range and append the generated error to the error stack.
func (p *Parser) endErr(endLoc errors.SourceContext, args ...any) {
	p.Errors = append(p.Errors, errors.NewError(p.currErrorCode, p.errStartCtx, endLoc, p.reportCtx, args...))
}

func (p *Parser) err(errorCode errors.ErrorCode, endLoc errors.SourceContext, args ...any) {
	p.beginErr(errorCode)
	p.endErr(endLoc, args...)
}

// Append the given node type to the current parse context.
func (p *Parser) pushNodeType(nodeType ast.NodeType) {
	p.nodeTypes = append(p.nodeTypes, nodeType)
}

// Remove the last inserted node type.
func (p *Parser) popNodeType() ast.NodeType {
	if len(p.nodeTypes) > 0 {
		retval := p.nodeTypes[len(p.nodeTypes)-1]
		p.nodeTypes = p.nodeTypes[:len(p.nodeTypes)-1]
		return retval
	}

	return ast.NodeUnknown
}

func (p *Parser) pushScope() {
	parent := p.rootScope
	if len(p.scopes) > 0 {
		parent = p.scopes[len(p.scopes)-1]
	}

	p.scopes = append(p.scopes, ast.NewScope(parent))
}

func (p *Parser) getScope() *ast.Scope {
	if len(p.scopes) == 0 {
		return p.rootScope
	}

	top := p.scopes[len(p.scopes)-1]
	return top
}

func (p *Parser) popScope() {
	p.scopes = p.scopes[0 : len(p.scopes)-1]
}

// Try to survive a nuclear disaster, e.g. try to find something sensible to continue from if things go haywire.
func (p *Parser) sync() {
	if p.atEnd() {
		return
	}

	tok := p.next()
	if tok.Kind == lexer.TokenSeparator {
		return
	}

out:
	for !p.atEnd() {
		switch p.peek().Kind {
		case lexer.TokenSeparator,
			lexer.TokenClass,
			lexer.TokenUse,
			lexer.TokenIf,
			lexer.TokenFor,
			lexer.TokenWhile:
			break out
		}

		p.next()
	}
}

// Return the current token and advance one.
func (p *Parser) next() *lexer.Token {
	tok := p.peek()
	if tok.Kind != lexer.TokenEof {
		p.curIdx++
	}

	return tok
}

// Go back by one.
func (p *Parser) back() {
	p.curIdx--
}

// Return true if any of the token types given match the current token type and move one forward.
func (p *Parser) expect(kind ...lexer.TokenKind) bool {
	if slices.Contains(kind, p.peek().Kind) {
		p.next()
		return true
	}

	return false
}

// Return true if any of the token types match the current token type.
func (p *Parser) is(kind ...lexer.TokenKind) bool {
	return slices.Contains(kind, p.peek().Kind)
}

// Scream if kind doesnt equal the current token type.
func (p *Parser) assert(kind lexer.TokenKind) errors.SourceContext {
	ctx := p.ctx()
	if !p.expect(kind) {
		panic(fmt.Sprintf("Expected kind %s", kind))
	}

	return ctx
}

// Expect a semicolon to be present.
func (p *Parser) requireSemi() {
	if !p.expect(lexer.TokenSeparator) {
		p.makeError(p.nodeTypes[len(p.nodeTypes)-1], p.ctx(), errors.ErrorExpectedSemi)
	}
}

// Returns the ending and starting position of the nodes in the list
func nodeListRange[X ast.Node](list []X) (errors.SourceContext, errors.SourceContext) {
	return list[0].Begin(), list[len(list)-1].End()
}

/* Actual parsing */

// Parse the input. Returns true if there were no errors.
func (p *Parser) Parse() bool {
	nodes := make([]ast.Node, 0)
	for !p.atEnd() {
		switch p.peek().Kind {
		case lexer.TokenUse:
			nodes = append(nodes, p.parseUseStmt())
		case lexer.TokenFunc:
			nodes = append(nodes, p.parseFunction())
		case lexer.TokenLet:
			nodes = append(nodes, p.parseVariableDeclaration())
		default:
			p.sync()
		}
	}

	if len(nodes) == 0 {
		p.Root = ast.NewRootNode(nodes, p.tokens[0].Context, p.tokens[0].Context, p.reportCtx, p.getScope())
	} else {
		start, end := nodeListRange(nodes)
		p.Root = ast.NewRootNode(nodes, start, end, p.reportCtx, p.getScope())
	}
	return len(p.Errors) == 0
}

// Parse any node that works in a block.
func (p *Parser) parseBlockNode() ast.Statement {
	switch p.peek().Kind {
	case lexer.TokenLet:
		return p.parseVariableDeclaration()
	default:
		return ast.NewExpressionStatement(p.parseExpression(), p.reportCtx, p.getScope())
	}
}

// Parse a use statement:
//
//	"use" qualifiedName ";"
func (p *Parser) parseUseStmt() *ast.UseStatement {
	p.pushNodeType(ast.NodeUseStatement)

	ctx := p.assert(lexer.TokenUse)
	qn := p.parseQualifiedName()
	p.requireSemi()

	p.popNodeType()
	return ast.NewUseStatement(qn, ctx, p.ctx(), p.reportCtx, p.getScope())
}

// Parse a qualified name:
//
//	identifier ( "." identifier )*
func (p *Parser) parseQualifiedName() *ast.QualifiedName {
	p.pushNodeType(ast.NodeQualifiedName)
	names := make([]*ast.Name, 0)
	for {
		names = append(names, p.parseName())
		if !p.expect(lexer.TokenDot) {
			break
		}
	}

	start, end := nodeListRange(names)
	p.popNodeType()
	return ast.NewQualifiedName(names, start, end, p.reportCtx, p.getScope())
}

// Parse a name:
//
//	identifier
func (p *Parser) parseName() *ast.Name {
	p.pushNodeType(ast.NodeName)
	tok := p.peek()
	if tok.Kind != lexer.TokenIdentifier && tok.Kind != lexer.TokenFloor { // floors are warned about later
		p.err(errors.ErrorExpectedIdentifier, p.ctx(), tok.Text)
	}

	p.next()
	p.popNodeType()
	return ast.NewName(tok.Text, tok.Context, tok.Context, p.reportCtx, p.getScope())
}

// Parse a variable delcaration:
//
//	"let" qualifiedName type? ( "=" expression )? ";"
func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	p.pushNodeType(ast.NodeVariableDeclaration)
	start := p.ctx()
	p.assert(lexer.TokenLet)
	name := p.parseQualifiedName()

	var dataType ast.DataType = nil
	if !p.is(lexer.TokenEq, lexer.TokenSeparator) {
		dataType = p.parseType()
	}

	var init ast.Expression = nil
	if p.expect(lexer.TokenEq) {
		init = p.parseExpression()
	}

	p.popNodeType()
	return ast.NewVariableDeclaration(name, dataType, init, start, p.ctx(), p.reportCtx, p.getScope())
}

// Parse a function:
//
//	"func" identifier "(" argumentList? ")" type? "{" blockBody "}"
func (p *Parser) parseFunction() *ast.Function {
	p.pushNodeType(ast.NodeFunction)
	start := p.ctx()
	p.assert(lexer.TokenFunc)
	name := p.parseQualifiedName()
	argList := make([]*ast.FunctionArg, 0)

	if !p.expect(lexer.TokenOpenParen) {
		p.err(errors.ErrorExpectedFnArgList, p.ctx())
	} else {
		argList = p.parseFnArgList()
	}

	var returnType ast.DataType = ast.NewNoneType(p.ctx(), p.ctx(), p.reportCtx, p.getScope())
	if p.is(lexer.TokenIdentifier, lexer.TokenOpenBracket) {
		returnType = p.parseType()
	}

	body := p.parseBlock()
	p.popNodeType()
	return ast.NewFunction(
		name, argList, body, returnType, start, p.ctx(), p.reportCtx, p.getScope(),
	)
}

// Parse function argument list:
//
//	dataType? ( ',' dataType )*
func (p *Parser) parseFnArgList() []*ast.FunctionArg {
	if p.expect(lexer.TokenCloseParen) {
		return []*ast.FunctionArg{}
	}

	args := make([]*ast.FunctionArg, 0)
	for {
		start := p.ctx()
		args = append(args, ast.NewFunctionArg(
			p.parseQualifiedName(),
			p.parseType(),
			start, p.ctx(),
			p.reportCtx, p.getScope(),
		))

		if !p.expect(lexer.TokenComma) {
			break
		}
	}

	if !p.expect(lexer.TokenCloseParen) {
		p.err(errors.ErrorExpectedClosingParen, p.ctx())
	}

	return args
}

// Parse a block:
//
//	"{" node* "}"
func (p *Parser) parseBlock() *ast.Block {
	p.pushNodeType(ast.NodeBlock)
	p.pushScope()

	if !p.expect(lexer.TokenOpenBrace) {
		p.err(errors.ErrorExpectedBlockBody, p.ctx())
		p.sync()
	}

	start := p.ctx()
	if p.expect(lexer.TokenCloseBrace) {
		return ast.NewBlock([]ast.Statement{}, start, p.ctx(), p.reportCtx, p.getScope())
	}

	nodes := make([]ast.Statement, 0)
	for !p.expect(lexer.TokenCloseBrace) {
		if p.atEnd() {
			p.err(errors.ErrorExpectedBlockEnd, p.ctx())
			break
		}

		nodes = append(nodes, p.parseBlockNode())
		p.requireSemi()
	}

	var startCtx, endCtx errors.SourceContext
	if len(nodes) == 0 {
		startCtx = start
		endCtx = start
	} else {
		startCtx, endCtx = nodeListRange(nodes)
	}

	p.popNodeType()
	scope := p.getScope()
	p.popScope()

	return ast.NewBlock(nodes, startCtx, endCtx, p.reportCtx, scope)
}

// Parse a datatype:
//
//	atomType | arrayType | genericType | nullableType
func (p *Parser) parseType() ast.DataType {
	tok := p.peek()
	p.pushNodeType(ast.NodeDataType)
	var t ast.DataType
	if tok.Kind == lexer.TokenOpenBracket {
		t = p.parseArrayType()
	} else {
		t = p.parseSingletType()
	}

	if p.expect(lexer.TokenQuery) {
		t = ast.NewNullableType(t, tok.Context, p.ctx(), p.reportCtx, p.getScope())
	}

	p.popNodeType()
	return t
}

// Parse an array type:
//
//	"["+ dataType "]"+
func (p *Parser) parseArrayType() *ast.ArrayType {
	start := p.ctx()
	p.beginErr(errors.ErrorExpectedArrayCloseBracket)

	brackCount := 0
	for p.expect(lexer.TokenOpenBracket) {
		brackCount++
	}

	inner := p.parseType()
	closeBrackCount := brackCount
	for p.expect(lexer.TokenCloseBracket) {
		closeBrackCount--
	}

	if closeBrackCount != 0 {
		p.endErr(p.ctx())
	}

	return ast.NewArrayType(inner, start, p.ctx(), p.reportCtx, p.getScope())
}

// Parse a singlet type, e.g. something that isnt an array or a none type.
func (p *Parser) parseSingletType() ast.DataType {
	start := p.ctx()
	// The typename
	name := p.parseQualifiedName()

	if p.expect(lexer.TokenOpenBracket) {
		inners := make([]ast.DataType, 0)

		for {
			inners = append(inners, p.parseType())
			if !p.expect(lexer.TokenComma) {
				break
			}
		}

		if !p.expect(lexer.TokenCloseBracket) {
			p.err(errors.ErrorExpectedArrayCloseBracket, p.ctx())
		}

		var t ast.DataType = ast.NewGenericType(name, inners, start, p.ctx(), p.reportCtx, p.getScope())
		if p.expect(lexer.TokenQuery) {
			t = ast.NewNullableType(t, start, p.ctx(), p.reportCtx, p.getScope())
		}

		return t
	}

	return ast.NewAtomType(name, start, p.ctx(), p.reportCtx, p.getScope())
}

/* Expression parsing */
/* Helpers */

var binaryOps map[lexer.TokenKind]ast.BinaryOperator = map[lexer.TokenKind]ast.BinaryOperator{
	lexer.TokenEqEq:             ast.BinaryOperatorEquals,
	lexer.TokenBangEq:           ast.BinaryOperatorNotEquals,
	lexer.TokenOpenArrowEq:      ast.BinaryOperatorLessThanOrEqual,
	lexer.TokenCloseArrowEq:     ast.BinaryOperatorMoreThanOrEqual,
	lexer.TokenWallWall:         ast.BinaryOperatorOr,
	lexer.TokenStarStar:         ast.BinaryOperatorPow,
	lexer.TokenAmpAmp:           ast.BinaryOperatorAnd,
	lexer.TokenUpUp:             ast.BinaryOperatorXor,
	lexer.TokenDotDot:           ast.BinaryOperatorChainAccess,
	lexer.TokenOpenArrow:        ast.BinaryOperatorLessThan,
	lexer.TokenCloseArrow:       ast.BinaryOperatorMoreThan,
	lexer.TokenWall:             ast.BinaryOperatorBitOr,
	lexer.TokenPercent:          ast.BinaryOperatorMod,
	lexer.TokenUp:               ast.BinaryOperatorBitXor,
	lexer.TokenAmp:              ast.BinaryOperatorBitAnd,
	lexer.TokenMinus:            ast.BinaryOperatorMinus,
	lexer.TokenPlus:             ast.BinaryOperatorPlus,
	lexer.TokenStar:             ast.BinaryOperatorMul,
	lexer.TokenSlash:            ast.BinaryOperatorDiv,
	lexer.TokenDot:              ast.BinaryOperatorAccess,
	lexer.TokenAs:               ast.BinaryOperatorAs,
	lexer.TokenIs:               ast.BinaryOperatorIs,
	lexer.TokenTo:               ast.BinaryOperatorTo,
	lexer.TokenIn:               ast.BinaryOperatorIn,
	lexer.TokenDoubleOpenArrow:  ast.BinaryOperatorLeftShift,
	lexer.TokenDoubleCloseArrow: ast.BinaryOperatorRightShift,
}

var prefixOps map[lexer.TokenKind]ast.PrefixOperator = map[lexer.TokenKind]ast.PrefixOperator{
	lexer.TokenPlus:       ast.PrefixOperatorPlus,
	lexer.TokenMinus:      ast.PrefixOperatorMinus,
	lexer.TokenTilde:      ast.PrefixOperatorNegate,
	lexer.TokenBang:       ast.PrefixOperatorNot,
	lexer.TokenPlusPlus:   ast.PrefixOperatorInc,
	lexer.TokenMinusMinus: ast.PrefixOperatorDec,
}

var prefixOpTypes []lexer.TokenKind = func() []lexer.TokenKind {
	keys := make([]lexer.TokenKind, len(prefixOps))
	i := 0
	for key := range prefixOps {
		keys[i] = key
		i++
	}

	return keys
}()

var postfixOps map[lexer.TokenKind]ast.PostfixOperator = map[lexer.TokenKind]ast.PostfixOperator{
	lexer.TokenPlusPlus: ast.PostfixOperatorInc,
	lexer.TokenMinus:    ast.PostfixOperatorDec,
}

var assignmentOps map[lexer.TokenKind]ast.BinaryOperator = map[lexer.TokenKind]ast.BinaryOperator{
	lexer.TokenEq:                 ast.BinaryOperatorAssign,
	lexer.TokenSlashEq:            ast.BinaryOperatorDivAssign,
	lexer.TokenStarEq:             ast.BinaryOperatorMulAssign,
	lexer.TokenPlusEq:             ast.BinaryOperatorPlusAssign,
	lexer.TokenMinusEq:            ast.BinaryOperatorMinusAssign,
	lexer.TokenAmpEq:              ast.BinaryOperatorBitAndAssign,
	lexer.TokenUpEq:               ast.BinaryOperatorBitXorAssign,
	lexer.TokenPercentEq:          ast.BinaryOperatorModAssign,
	lexer.TokenWallEq:             ast.BinaryOperatorBitOrAssign,
	lexer.TokenUpUpEq:             ast.BinaryOperatorXorAssign,
	lexer.TokenAmpAmpEq:           ast.BinaryOperatorAndAssign,
	lexer.TokenStarStarEq:         ast.BinaryOperatorPowAssign,
	lexer.TokenWallWallEq:         ast.BinaryOperatorOrAssign,
	lexer.TokenDoubleOpenArrowEq:  ast.BinaryOperatorLeftShiftAssign,
	lexer.TokenDoubleCloseArrowEq: ast.BinaryOperatorRightShiftAssign,
}

var assignmentOpTypes []lexer.TokenKind = func() []lexer.TokenKind {
	keys := make([]lexer.TokenKind, len(assignmentOps))
	i := 0
	for key := range assignmentOps {
		keys[i] = key
		i++
	}

	return keys
}()

func (p *Parser) binaryOp() ast.BinaryOperator {
	p.beginErr(errors.ErrorExpectedOperator)
	tok := p.next()
	val, ok := binaryOps[tok.Kind]
	if !ok {
		p.endErr(p.ctx())
		return ast.BinaryOperatorInvalid
	}

	return val
}

func (p *Parser) prefixOp() ast.PrefixOperator {
	p.beginErr(errors.ErrorExpectedOperator)
	tok := p.next()
	val, ok := prefixOps[tok.Kind]
	if !ok {
		p.endErr(p.ctx())
		return ast.PrefixOperatorInvalid
	}

	return val
}

func (p *Parser) postfixOp() ast.PostfixOperator {
	p.beginErr(errors.ErrorExpectedOperator)
	tok := p.next()
	val, ok := postfixOps[tok.Kind]
	if !ok {
		p.endErr(p.ctx())
		return ast.PostfixOperatorInvalid
	}

	return val
}

func (p *Parser) assignmentOp() (ast.BinaryOperator, bool) {
	if slices.Contains(assignmentOpTypes, p.peek().Kind) {
		t := p.next().Kind
		return assignmentOps[t], true
	}

	return 0, false
}

/* Parsing methods */

func (p *Parser) parseExpression() ast.Expression {
	return p.parseAssignment()
}

func (p *Parser) parseAssignment() ast.Expression {
	left := p.parseEquality()

	if op, ok := p.assignmentOp(); ok {
		right := p.parseAssignment()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseEquality() ast.Expression {
	left := p.parseComparison()

	for p.is(lexer.TokenEqEq, lexer.TokenBangEq) {
		op := p.binaryOp()
		right := p.parseComparison()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseComparison() ast.Expression {
	left := p.parseAddition()

	for p.is(lexer.TokenOpenArrow, lexer.TokenOpenArrowEq,
		lexer.TokenCloseArrow, lexer.TokenCloseArrowEq) {
		op := p.binaryOp()
		right := p.parseAddition()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseAddition() ast.Expression {
	left := p.parseMultiplication()

	for p.is(lexer.TokenPlus, lexer.TokenMinus) {
		op := p.binaryOp()
		right := p.parseMultiplication()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseMultiplication() ast.Expression {
	left := p.parsePowers()

	for p.is(lexer.TokenStar, lexer.TokenSlash) {
		op := p.binaryOp()
		right := p.parsePowers()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parsePowers() ast.Expression {
	left := p.parsePrefix()
	for p.is(lexer.TokenStarStar) {
		op := p.binaryOp()
		right := p.parsePowers()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parsePrefix() ast.Expression {
	if slices.Contains(prefixOpTypes, p.peek().Kind) {
		ctx := p.ctx()
		op := p.prefixOp()
		right := p.parsePrefix()
		return ast.NewPrefixExpression(right, op, ctx, p.ctx(), p.reportCtx, p.getScope())
	}

	return p.parsePostfix()
}

func (p *Parser) parsePostfix() ast.Expression {
	left := p.parseFnCall()
	for p.is(lexer.TokenPlusPlus, lexer.TokenMinusMinus) {
		op := p.postfixOp()
		left = ast.NewPostfixExpression(left, op, left.Begin(), p.ctx(), p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseFnCall() ast.Expression {
	left := p.parseObjectSubscript()
	for p.expect(lexer.TokenOpenParen) {
		args := make([]ast.Expression, 0)
		if !p.expect(lexer.TokenCloseParen) {
			for {
				args = append(args, p.parseExpression())
				if !p.expect(lexer.TokenComma) {
					break
				}
			}

			if !p.expect(lexer.TokenCloseParen) {
				p.err(errors.ErrorExpectedClosingParen, p.ctx())
			}
		}

		left = ast.NewCallExpression(left, args, left.Begin(), p.ctx(), p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseObjectSubscript() ast.Expression {
	left := p.parseArraySubscript()
	for p.is(lexer.TokenDot, lexer.TokenDotDot) {
		op := p.binaryOp()
		right := p.parseArraySubscript()
		left = ast.NewBinaryExpression(left, right, op, p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseArraySubscript() ast.Expression {
	left := p.parseAtom()
	for p.expect(lexer.TokenOpenBracket) {
		right := p.parseAtom()
		if !p.expect(lexer.TokenCloseBracket) {
			p.err(errors.ErrorExpectedArrayAccessCloseBracket, p.ctx())
		}

		left = ast.NewSubscriptExpression(left, right, left.Begin(), p.ctx(), p.reportCtx, p.getScope())
	}

	return left
}

func (p *Parser) parseAtom() ast.Expression {
	tok := p.next()
	switch tok.Kind {
	case lexer.TokenNumber:
		return ast.NewDecimalLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenHexNumber:
		return ast.NewHexadecimalLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenBinNumber:
		return ast.NewBinaryLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenOctNumber:
		return ast.NewOctalLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenRomanNumber:
		return ast.NewRomanLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenString:
		return ast.NewStringLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenYes, lexer.TokenNo:
		return ast.NewBooleanLiteral(tok, p.reportCtx, p.getScope())
	case lexer.TokenIdentifier:
		return ast.NewVariableReference(ast.NewQualifiedName(
			[]*ast.Name{ast.NewName(tok.Text, tok.Context, tok.Context, p.reportCtx, p.getScope())},
			tok.Context, tok.Context, p.reportCtx, p.getScope(),
		), p.reportCtx, p.getScope())
	case lexer.TokenOpenParen:
		if p.expect(lexer.TokenCloseParen) {
			p.err(errors.ErrorExpectedEpression, p.ctx())
			return ast.NewInvalidExpression(p.ctx(), p.ctx(), p.reportCtx, p.getScope())
		}

		expr := ast.NewParenExpression(
			p.parseExpression(),
			tok.Context, p.ctx(),
			p.reportCtx, p.getScope(),
		)

		if !p.expect(lexer.TokenCloseParen) {
			p.err(errors.ErrorExpectedClosingParen, p.ctx())
		}

		return expr
	default:
		p.back()
		p.err(errors.ErrorExpectedEpression, p.ctx())
		p.sync()
		return ast.NewInvalidExpression(tok.Context, tok.Context, p.reportCtx, p.getScope())
	}
}
