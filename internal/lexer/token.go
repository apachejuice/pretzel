package lexer

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/errors"
)

type (
	// Specifies the type of a token.
	TokenKind int

	// Represents a single token found in the source code.
	Token struct {
		// The type of the token.
		Kind TokenKind `json:"kind"`
		// A string containing the text of the token.
		Text string `json:"text"`
		// A context object containing the location of the token.
		Context errors.SourceContext `json:"context"`
	}
)

var _ fmt.Stringer = Token{}
var _ fmt.Stringer = TokenKind(0)

func (t Token) String() string {
	return fmt.Sprintf("Token(%s '%s' %s)", t.Kind, t.Text, t.Context)
}

func (k TokenKind) String() string {
	return tokenNames[k]
}

// Constants for TokenKind
const (
	/* 1 char */
	TokenSeparator    TokenKind = iota // ;
	TokenOpenBrace                     // {
	TokenCloseBrace                    // }
	TokenOpenParen                     // (
	TokenCloseParen                    // )
	TokenOpenBracket                   // [
	TokenCloseBracket                  // ]
	TokenDot                           // .
	TokenComma                         // ,
	TokenColon                         // :
	TokenSlash                         // /
	TokenStar                          // *
	TokenPlus                          // +
	TokenMinus                         // -
	TokenAmp                           // &
	TokenUp                            // ^
	TokenPercent                       // %
	TokenAt                            // @
	TokenHash                          // #
	TokenBang                          // !
	TokenQuery                         // ?
	TokenWall                          // |
	TokenFloor                         // _
	TokenOpenArrow                     // <
	TokenCloseArrow                    // >
	TokenTilde                         // ~
	TokenEq                            // =

	/* 2 chars */
	TokenSlashEq            // /=
	TokenStarEq             // *=
	TokenPlusEq             // +=
	TokenMinusEq            // -=
	TokenAmpEq              // &=
	TokenUpEq               // ^=
	TokenPercentEq          // %=
	TokenWallEq             // |=
	TokenUpUpEq             // ^^=
	TokenAmpAmpEq           // &&=
	TokenStarStarEq         // **=
	TokenWallWallEq         // ||=
	TokenDoubleOpenArrowEq  // <<=
	TokenDoubleCloseArrowEq // >>=
	TokenDotDot             // ..
	TokenPlusPlus           // ++
	TokenStarStar           // **
	TokenAmpAmp             // &&
	TokenUpUp               // ^^
	TokenWallWall           // ||
	TokenMinusMinus         // --
	TokenCloseArrowEq       // >=
	TokenOpenArrowEq        // <=
	TokenEqEq               // ==
	TokenBangEq             // !=
	TokenAs                 // as
	TokenIs                 // is
	TokenTo                 // to
	TokenIn                 // in
	TokenIf                 // if
	TokenNo                 // no
	TokenDoubleOpenArrow    // <<
	TokenDoubleCloseArrow   // >>

	/* 3 chars */
	TokenDotDotDot // ...
	TokenFor       // for
	TokenUse       // use
	TokenYes       // yes
	TokenLet       // let

	/* 4 chars */
	TokenFrom // from
	TokenFunc // func
	TokenElse // else

	/* 5 chars */
	TokenWhile // while
	TokenClass // class

	/* variable */
	TokenString      // "string in quotes" 'another one'
	TokenChar        // `a`
	TokenIdentifier  // this_is_an_identifier
	TokenNumber      // 8293742984
	TokenHexNumber   // 0xAB23478BBFC
	TokenBinNumber   // 0b101010010010111
	TokenOctNumber   // 0o172536771526
	TokenRomanNumber // 0rIIICXVI

	TokenEof
)

var tokenNames map[TokenKind]string = map[TokenKind]string{
	TokenSeparator:    "separator",
	TokenOpenBrace:    "opening brace",
	TokenCloseBrace:   "closing brace",
	TokenOpenParen:    "opening paren",
	TokenCloseParen:   "closing paren",
	TokenOpenBracket:  "opening bracket",
	TokenCloseBracket: "closing bracket",
	TokenDot:          "dot",
	TokenComma:        "comma",
	TokenColon:        "colon",
	TokenSlash:        "slash",
	TokenStar:         "star",
	TokenPlus:         "plus",
	TokenMinus:        "minus",
	TokenAmp:          "ampersand",
	TokenUp:           "up arrow",
	TokenPercent:      "percent",
	TokenAt:           "at",
	TokenHash:         "hash",
	TokenBang:         "bang",
	TokenQuery:        "question mark",
	TokenWall:         "wall",
	TokenFloor:        "floor",
	TokenOpenArrow:    "open arrow",
	TokenCloseArrow:   "close arrow",
	TokenTilde:        "tilde",
	TokenEq:           "assignment",
	TokenDotDot:       "double dot",
	TokenPlusPlus:     "double plus",
	TokenStarStar:     "double star",
	TokenAmpAmp:       "double ampersand",
	TokenUpUp:         "double up arrow",
	TokenWallWall:     "double wall",
	TokenMinusMinus:   "double minus",
	TokenOpenArrowEq:  "less or equal",
	TokenCloseArrowEq: "more or equal",
	TokenEqEq:         "equal",
	TokenBangEq:       "not equal",
	TokenAs:           "kw_as",
	TokenIs:           "kw_is",
	TokenTo:           "kw_to",
	TokenIn:           "kw_in",
	TokenIf:           "kw_if",
	TokenDotDotDot:    "triple dot",
	TokenFor:          "kw_for",
	TokenUse:          "kw_use",
	TokenFrom:         "kw_from",
	TokenFunc:         "kw_func",
	TokenElse:         "kw_else",
	TokenWhile:        "kw_while",
	TokenClass:        "kw_class",
	TokenYes:          "kw_yes",
	TokenNo:           "kw_no",
	TokenString:       "string",
	TokenChar:         "char",
	TokenIdentifier:   "identifier",

	TokenEof: "EOF",
}
