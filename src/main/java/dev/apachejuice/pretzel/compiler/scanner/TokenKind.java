package dev.apachejuice.pretzel.compiler.scanner;

public enum TokenKind {
    BANG("!"),
    BANG_EQUAL("!="),
    EQUAL("="),
    EQUAL_EQUAL("=="),
    PLUS("+"),
    PLUS_PLUS("++"),
    PLUS_EQUAL("+="),
    MINUS("-"),
    MINUS_MINUS("--"),
    MINUS_EQUAL("-="),
    MULTIPLY("*"),
    POWER_OF("**"),
    MULTIPLY_EQUAL("*="),
    POWER_OF_EQUAL("**="),
    DIVIDE("/"),
    DIVIDE_EQUAL("/="),
    MODULO("%"),
    MODULO_EQUAL("%="),
    BITWISE_AND("&"),
    LOGICAL_AND("&&"),
    BITWISE_OR("|"),
    LOGICAL_OR("||"),
    BITWISE_XOR("^"),
    BITWISE_NEGATE("~"),
    OPEN_PAREN("("),
    CLOSE_PAREN(")"),
    OPEN_BRACE("{"),
    CLOSE_BRACE("}"),
    OPEN_BRACKET("["),
    CLOSE_BRACKET("]"),

    STRING_LITERAL,
    IDENTIFIER,
    DECIMAL_LITERAL,
    REAL_LITERAL,

    KW_YES("yes"),
    KW_NO("no"),
    KW_NONE("none"),
    KW_FUNC("func"),
    KW_LET("let"),
    KW_RETURN("return"),

    EOF("");

    private final String defaultValue;

    TokenKind(String defaultValue) {
        this.defaultValue = defaultValue;
    }

    TokenKind() {
        this(null);
    }

    public boolean isConstant() {
        return defaultValue != null;
    }

    public String getValue() {
        return defaultValue;
    }
}
