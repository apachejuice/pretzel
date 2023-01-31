package lexer

import (
	"unicode"

	"github.com/apachejuice/pretzel/internal/errors"
	"golang.org/x/exp/slices"
)

// The actual lexer object. Produces a list of tokens from string input
type Lexer struct {
	/* context variables */
	curLine, curCol, curPos int
	inputName               string
	input                   []rune

	/* source context helpers */
	ctxStartLine, ctxStartCol int
	inCtxMode                 bool

	/* backup variables */
	buLine, buCol, buPos int

	// A hard limit of errors at which to stop. If 0, 15 is used.
	// If -1, there is no limit.
	// Note that the lexer may return more errors than that,
	// if the scanning of one token produces multiple errors.
	hardErrorLimit int

	// The report context passed to the parser
	ReportCtx *errors.ReportContext
	// The tokens this lexer produced.
	Tokens []*Token
	// The errors this lexer encountered.
	Errors []*errors.Error
	// Did the lexer stop early due to the hard error limit?
	FailureStop bool
}

func NewLexer(input, file string, hardErrorLimit int) *Lexer {
	return &Lexer{
		curLine:        0,
		curCol:         0,
		curPos:         0,
		input:          []rune(input),
		inputName:      file,
		hardErrorLimit: hardErrorLimit,
		ReportCtx:      &errors.ReportContext{Source: input},
		Tokens:         make([]*Token, 0),
		Errors:         make([]*errors.Error, 0),
	}
}

/* Helper methods */

func (l Lexer) peek() rune {
	return l.input[l.curPos]
}

func (l Lexer) atEnd() bool {
	return l.curPos >= len(l.input)
}

func (l *Lexer) backup() {
	l.buLine = l.curLine
	l.buCol = l.curCol
	l.buPos = l.curPos
}

func (l *Lexer) restore() {
	l.curLine = l.buLine
	l.curCol = l.buCol
	l.curPos = l.buPos
}

func (l *Lexer) next() rune {
	r := l.peek()
	if r == '\n' {
		l.curLine++
		l.curCol = 0
	} else {
		l.curCol++
	}
	l.curPos++

	return r
}

func (l *Lexer) match(want rune) bool {
	r := l.peek()
	if r == want {
		l.next()
		return true
	}

	return false
}

func (l *Lexer) matchMany(want string) bool {
	left := len(l.input) - l.curPos
	if len(want) > left {
		return false
	}

	l.backup()
	for _, r := range want {
		if !l.match(r) {
			l.restore()
			return false
		}
	}

	return true
}

/* Lexing helpers */
func (l *Lexer) begin() {
	l.inCtxMode = true
	l.ctxStartLine = l.curLine
	l.ctxStartCol = l.curCol
}

func (l *Lexer) steal() errors.SourceContext { // end() but doesnt stop the range
	return errors.SourceContext{
		StartLine:   l.ctxStartLine,
		StartColumn: l.ctxStartCol,
		EndLine:     l.curLine,
		EndColumn:   l.curCol,
		Path:        l.inputName,
	}
}

func (l *Lexer) end() errors.SourceContext {
	l.inCtxMode = false
	return errors.SourceContext{
		StartLine:   l.ctxStartLine,
		StartColumn: l.ctxStartCol,
		EndLine:     l.curLine,
		EndColumn:   l.curCol,
		Path:        l.inputName,
	}
}

func (l *Lexer) ignore() errors.SourceContext {
	ctx := l.end()
	l.next()
	return ctx
}

func (l *Lexer) addToken(kind TokenKind, literal string) {
	ctx := l.end()
	l.Tokens = append(l.Tokens, &Token{
		Kind:    kind,
		Text:    literal,
		Context: ctx,
	})
}

var noErrors []*errors.Error = make([]*errors.Error, 0)

// A map of keywords to token types
var keywords map[string]TokenKind = map[string]TokenKind{
	"while": TokenWhile,
	"class": TokenClass,
	"else":  TokenElse,
	"from":  TokenFrom,
	"func":  TokenFunc,
	"for":   TokenFor,
	"use":   TokenUse,
	"yes":   TokenYes,
	"let":   TokenLet,
	"if":    TokenIf,
	"in":    TokenIn,
	"to":    TokenTo,
	"is":    TokenIs,
	"as":    TokenAs,
	"no":    TokenNo,
}

func (l *Lexer) identifierOrKeyword() []*errors.Error {
	buf := ""
	for !l.atEnd() && ((unicode.IsLetter(l.peek()) || l.peek() == '_') || (buf != "" && unicode.IsNumber(l.peek()))) {
		buf += string(l.next())
	}

	if kw, ok := keywords[buf]; ok {
		l.addToken(kw, buf)
	} else {
		l.addToken(TokenIdentifier, buf)
	}

	return noErrors
}

func (l *Lexer) stringLiteral() []*errors.Error {
	q := l.next() // stores the type of the quote so there's no need to switch between " and '
	buf := ""
	errs := make([]*errors.Error, 0)

	for {
		if l.atEnd() {
			errs = append(errs, errors.NewError(errors.ErrorEofAtString, l.end(), l.end(), l.ReportCtx))
			break
		}

		r := l.next()
		if r == q {
			break
		}

		if r == '\\' {
			// TODO \xXX \uXXXX and \UXXXXXXXX
			switch l.next() {
			case 'n':
				buf += "\n"
			case 'r':
				buf += "\r"
			case 'b':
				buf += "\b"
			case 't':
				buf += "\t"
			case '\\':
				buf += "\\"
			case q:
				buf += string(q)
			default:
				errs = append(errs, errors.NewError(errors.ErrorIllegalEscapeChar, l.steal(), l.steal(), l.ReportCtx))
			}
		} else {
			buf += string(r)
		}
	}

	l.addToken(TokenString, buf)
	return errs
}

var hexDigits []rune = []rune{
	'0', '1', '2', '3', '4', '5', '6',
	'7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
}

var romanDigits []rune = []rune{
	'i', 'v', 'x', 'l', 'c', 'd', 'm', '_',
}

func (l *Lexer) number() []*errors.Error {
	begin := l.next()
	buf := string(begin)

	l.backup()
	cur := unicode.ToLower(l.next())
	switch {
	case cur == 'x':
		for slices.Contains(hexDigits, unicode.ToLower(l.peek())) {
			buf += string(l.next())
		}

		l.addToken(TokenHexNumber, buf)
	case cur == 'b':
		for l.peek() == '0' || l.peek() == '1' {
			buf += string(l.next())
		}

		l.addToken(TokenBinNumber, buf)
	case cur == 'o':
		for l.peek() >= '0' && l.peek() <= '7' {
			buf += string(l.next())
		}

		l.addToken(TokenOctNumber, buf)
	case cur == 'r':
		for slices.Contains(romanDigits, unicode.ToLower(l.peek())) {
			buf += string(l.next())
		}

		l.addToken(TokenRomanNumber, buf)
	case unicode.IsNumber(cur):
		for unicode.IsNumber(l.peek()) {
			buf += string(l.next())
		}

		l.addToken(TokenNumber, buf)
	default:
		l.restore()
		l.addToken(TokenNumber, buf)
	}

	return noErrors
}

// Get a single token from the input.
func (l *Lexer) getToken() []*errors.Error {
	r := l.peek()
	l.begin()

	switch {
	// First, anything skippable
	case unicode.IsSpace(r):
		l.ignore()
		return noErrors

	// Longer to shorter tokens (not keywords or '_')
	case l.matchMany("..."):
		l.addToken(TokenDotDotDot, "...")
	case l.matchMany("!="):
		l.addToken(TokenBangEq, "!=")
	case l.matchMany("=="):
		l.addToken(TokenEqEq, "==")
	case l.matchMany("<="):
		l.addToken(TokenOpenArrowEq, "<=")
	case l.matchMany(">="):
		l.addToken(TokenCloseArrowEq, ">=")
	case l.matchMany("--"):
		l.addToken(TokenMinusMinus, "--")
	case l.matchMany("||"):
		l.addToken(TokenWallWall, "||")
	case l.matchMany("^^"):
		l.addToken(TokenUpUp, "^^")
	case l.matchMany("&&"):
		l.addToken(TokenAmpAmp, "&&")
	case l.matchMany("**"):
		l.addToken(TokenStarStar, "**")
	case l.matchMany("++"):
		l.addToken(TokenPlusPlus, "++")
	case l.matchMany(".."):
		l.addToken(TokenDotDot, "..")
	case l.matchMany("<<"):
		l.addToken(TokenDoubleOpenArrow, "<<")
	case l.matchMany(">>"):
		l.addToken(TokenDoubleCloseArrow, ">>")

	// Single character
	case l.match('_'):
		l.addToken(TokenFloor, "_")
	case l.match('='):
		l.addToken(TokenEq, "=")
	case l.match('~'):
		l.addToken(TokenTilde, "~")
	case l.match('>'):
		l.addToken(TokenCloseArrow, ">")
	case l.match('<'):
		l.addToken(TokenOpenArrow, "<")
	case l.match('|'):
		l.addToken(TokenWall, "|")
	case l.match('!'):
		l.addToken(TokenBang, "!")
	case l.match('?'):
		l.addToken(TokenQuery, "?")
	case l.match('#'):
		l.addToken(TokenHash, "#")
	case l.match('@'):
		l.addToken(TokenAt, "@")
	case l.match('%'):
		l.addToken(TokenPercent, "%")
	case l.match('^'):
		l.addToken(TokenUp, "^")
	case l.match('&'):
		l.addToken(TokenAmp, "&")
	case l.match('-'):
		l.addToken(TokenMinus, "-")
	case l.match('+'):
		l.addToken(TokenPlus, "+")
	case l.match('*'):
		l.addToken(TokenStar, "*")
	case l.match('/'):
		l.addToken(TokenSlash, "/")
	case l.match(':'):
		l.addToken(TokenColon, ":")
	case l.match(','):
		l.addToken(TokenComma, ",")
	case l.match('.'):
		l.addToken(TokenDot, ".")
	case l.match(']'):
		l.addToken(TokenCloseBracket, "]")
	case l.match('['):
		l.addToken(TokenOpenBracket, "[")
	case l.match(')'):
		l.addToken(TokenCloseParen, ")")
	case l.match('('):
		l.addToken(TokenOpenParen, "(")
	case l.match('{'):
		l.addToken(TokenOpenBrace, "{")
	case l.match('}'):
		l.addToken(TokenCloseBrace, "}")
	case l.match(';'):
		l.addToken(TokenSeparator, ";")

	case unicode.IsLetter(r):
		return l.identifierOrKeyword()
	case r == '"' || r == '\'':
		return l.stringLiteral()
	case unicode.IsNumber(r):
		return l.number()

	default:
		ctx := l.ignore()
		return []*errors.Error{errors.NewError(errors.ErrorUnexpectedCharacter, ctx, l.end(), l.ReportCtx, string(r))}
	}

	return noErrors
}

// Processes the input and stores the result in the Tokens variable.
// Returns true on success with no errors
func (l *Lexer) GetTokens() bool {
	neverStop := l.hardErrorLimit == -1
	if l.hardErrorLimit == 0 {
		l.hardErrorLimit = 15
	}

	for !l.atEnd() {
		errors := l.getToken()
		l.Errors = append(l.Errors, errors...)
		if !neverStop && len(l.Errors) >= l.hardErrorLimit {
			return false
		}
	}

	l.Tokens = append(l.Tokens, &Token{
		Kind:    TokenEof,
		Text:    "",
		Context: l.Tokens[len(l.Tokens)-1].Context,
	})
	return len(l.Errors) == 0
}
