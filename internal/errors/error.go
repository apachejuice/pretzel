package errors

import (
	"fmt"
	"strings"

	"github.com/apachejuice/pretzel/internal/common/constraint"
)

type (
	// Represents an error code.
	ErrorCode int

	// Represents an error in the source code.
	Error struct {
		// The location at which the error was found at
		Begin SourceContext `json:"begin"`
		// The location at which the error ends
		End SourceContext `json:"end"`
		// The source code between begin and end.
		Source string `json:"-"`
		// The error code
		Code ErrorCode `json:"errorCode"`
		// Any format arguments
		Args []any `json:"-"`
	}

	// Carries information about error reporting such as the entire source code of a file.
	ReportContext struct {
		// The source code in the file.
		Source string
	}

	/* declared here to prevent import cycle */

	// Any location in a source file.
	SourceContext struct {
		// The relative path from the source tree.
		Path string `json:"path"`
		// A starting line number
		StartLine int `json:"startLine"`
		// An ending line number
		EndLine int `json:"endLine"`
		// A starting column number
		StartColumn int `json:"startColumn"`
		// An ending column number
		EndColumn int `json:"endColumn"`
	}
)

func NewError(code ErrorCode, begin, end SourceContext, ctx *ReportContext, args ...any) *Error {
	constraint.NotNil("ctx", ctx)

	return &Error{
		Begin:  begin,
		End:    end,
		Source: getSource(begin, end, ctx.Source),
		Code:   code,
		Args:   args,
	}
}

func getSource(start, end SourceContext, fullSource string) string {
	return strings.Split(fullSource, "\n")[start.StartLine]
}

func (e ErrorCode) IsWarning() bool {
	return e >= WarningRedundantParentheses
}

const (
	ErrorUnexpectedCharacter ErrorCode = iota
	ErrorEofAtString
	ErrorEofAtBlockComment
	ErrorIllegalEscapeChar
	ErrorIllegalEscapeLen
	ErrorUnexpectedToken
	ErrorUnexpectedEndOfInput
	ErrorExpectedIdentifier
	ErrorExpectedSemi
	ErrorExpectedFnArgList
	ErrorExpectedArrayCloseBracket
	ErrorExpectedBlockBody
	ErrorExpectedBlockEnd
	ErrorExpectedOperator
	ErrorExpectedEpression
	ErrorExpectedClosingParen
	ErrorExpectedArrayAccessCloseBracket
	ErrorReferencedEntityNotVariable
	ErrorNoSuchVariable

	WarningRedundantParentheses
	WarningRedundantParenthesesInExprStmt
	WarningIneffectiveStatement
)

var errorTexts map[ErrorCode]string = map[ErrorCode]string{
	ErrorUnexpectedCharacter:             "Unexpected token: '%s'",
	ErrorEofAtString:                     "Missing ending quote in a string literal",
	ErrorEofAtBlockComment:               "Missing */ in a block comment",
	ErrorIllegalEscapeChar:               "Illegal character after \\ in escape sequence",
	ErrorIllegalEscapeLen:                "Illegal length of hexadecimal string escape",
	ErrorUnexpectedToken:                 "Unexpected token: %s",
	ErrorUnexpectedEndOfInput:            "Unexpected EOF",
	ErrorExpectedIdentifier:              "Expected identifier, got %s",
	ErrorExpectedSemi:                    "Expected a semicolon after a statement",
	ErrorExpectedFnArgList:               "Expected function argument list",
	ErrorExpectedArrayCloseBracket:       "Expected closing ']' for array/generic type",
	ErrorExpectedBlockBody:               "Expected '{' to start a block",
	ErrorExpectedBlockEnd:                "Expected '}' to end a block",
	ErrorExpectedOperator:                "Expected an operator in this context",
	ErrorExpectedEpression:               "Expected expression",
	ErrorExpectedClosingParen:            "Expected closing parentheses",
	ErrorExpectedArrayAccessCloseBracket: "Expected closing ']' for array access",
	ErrorReferencedEntityNotVariable:     "Symbol %s is not a variable",
	ErrorNoSuchVariable:                  "Symbol %s not found in scope",

	WarningRedundantParentheses:           "Redundant parentheses around expression",
	WarningRedundantParenthesesInExprStmt: "Redundant parentheses in expression statement",
	WarningIneffectiveStatement:           "Expression statement has no effect",
}

var _ fmt.Stringer = Error{}
var _ fmt.Stringer = SourceContext{}
var _ fmt.Stringer = ErrorCode(0)

func (s SourceContext) String() string {
	return fmt.Sprintf("SourceContext(%s %d:%d-%d:%d)", s.Path, s.StartLine, s.StartColumn, s.EndLine, s.EndColumn)
}

func (e Error) Formatted() string {
	return fmt.Sprintf(e.Code.String(), e.Args...)
}

func (l Error) String() string {
	return fmt.Sprintf("Error(%s-%s '%s')", l.Begin, l.End, l.Code)
}

func (l ErrorCode) String() string {
	return errorTexts[l]
}
