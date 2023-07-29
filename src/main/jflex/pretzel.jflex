package dev.apachejuice.pretzel.compiler.scanner;

%%

%{
    private StringBuilder string = new StringBuilder();

    private Token mkToken(TokenKind kind) {
      if (!kind.isConstant()) {
          return new Token(kind, yytext(), yyline, yycolumn);
      }

      return new Token(kind, kind.getValue(), yyline, yycolumn);
    }

    private Token mkToken(TokenKind kind, String text) {
      return new Token(kind, text, yyline, yycolumn);
    }
%}

%class PretzelScanner
%public
%final
%line
%column
%unicode
%type Token

Identifier = {Alpha}{AlphaNum}*
DecimalLiteral = [1-9]{Num}*
RealLiteral = {Num}* "." {Num}+ {RealSpec}

RealSpec = [fF] | [dD] | ( [eE] "-"? {DecimalLiteral}? )
Alpha = [:jletter:]
AlphaNum = [:jletterdigit:]
Num = [:digit:]

Ws = [ \t\n\r]

%state STRING
%%

<YYINITIAL> {
    "!="        { return mkToken(TokenKind.BANG_EQUAL); }
    "!"         { return mkToken(TokenKind.BANG); }
    "=="        { return mkToken(TokenKind.EQUAL_EQUAL); }
    "="         { return mkToken(TokenKind.EQUAL); }

    "+="        { return mkToken(TokenKind.PLUS_EQUAL); }
    "++"        { return mkToken(TokenKind.PLUS_PLUS); }
    "+"         { return mkToken(TokenKind.PLUS); }

    "-="        { return mkToken(TokenKind.MINUS_EQUAL); }
    "--"        { return mkToken(TokenKind.MINUS_MINUS); }
    "-"         { return mkToken(TokenKind.MINUS); }

    "**="       { return mkToken(TokenKind.POWER_OF_EQUAL); }
    "**"        { return mkToken(TokenKind.POWER_OF); }
    "*="        { return mkToken(TokenKind.MULTIPLY_EQUAL); }
    "*"         { return mkToken(TokenKind.MULTIPLY); }

    "/="        { return mkToken(TokenKind.DIVIDE_EQUAL); }
    "/"         { return mkToken(TokenKind.DIVIDE); }

    "%="        { return mkToken(TokenKind.MODULO_EQUAL); }
    "%"         { return mkToken(TokenKind.MODULO); }

    "&&"        { return mkToken(TokenKind.LOGICAL_AND); }
    "&"         { return mkToken(TokenKind.BITWISE_AND); }

    "||"        { return mkToken(TokenKind.LOGICAL_OR); }
    "|"         { return mkToken(TokenKind.BITWISE_OR); }

    "^"         { return mkToken(TokenKind.BITWISE_XOR); }
    "~"         { return mkToken(TokenKind.BITWISE_NEGATE); }

    "("         { return mkToken(TokenKind.OPEN_PAREN); }
    ")"         { return mkToken(TokenKind.CLOSE_PAREN); }
    "{"         { return mkToken(TokenKind.OPEN_BRACE); }
    "}"         { return mkToken(TokenKind.CLOSE_BRACE); }
    "["         { return mkToken(TokenKind.OPEN_BRACKET); }
    "]"         { return mkToken(TokenKind.CLOSE_BRACKET); }

    "yes"       { return mkToken(TokenKind.KW_YES); }
    "no"        { return mkToken(TokenKind.KW_NO); }
    "none"      { return mkToken(TokenKind.KW_NONE); }
    "func"      { return mkToken(TokenKind.KW_FUNC); }
    "let"       { return mkToken(TokenKind.KW_LET); }
    "return"    { return mkToken(TokenKind.KW_RETURN); }

    "\""        { yybegin(STRING); }

    {Identifier}      { return mkToken(TokenKind.IDENTIFIER); }
    {DecimalLiteral}  { return mkToken(TokenKind.DECIMAL_LITERAL); }
    {RealLiteral}     { return mkToken(TokenKind.REAL_LITERAL); }

    {Ws}            { /* skippity skip */ }

    .          { throw new RuntimeException("I'm not quite sure what you meant by " + yytext()); }

    <<EOF>> { return mkToken(TokenKind.EOF); }
}

<STRING> {
    "\""        { yybegin(YYINITIAL); return mkToken(TokenKind.STRING_LITERAL, string.toString()); }
    .           { string.append(yytext()); }
    {Ws}        { string.append(yytext()); }
   <<EOF>>      { throw new RuntimeException("Unexpected eof at string literal"); }
}