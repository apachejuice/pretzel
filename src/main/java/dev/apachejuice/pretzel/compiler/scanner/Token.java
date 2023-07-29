package dev.apachejuice.pretzel.compiler.scanner;

import java.util.Objects;

public class Token {
    private final String text;
    private final int startLine;
    private final int startColumn;
    private final int endLine;
    private final int endColumn;
    private final TokenKind kind;


    public Token(TokenKind kind, String text, int line, int column) {
        this.text = text;
        this.startLine = line;
        this.startColumn = column;
        this.kind = kind;

        this.endLine = (int) (line + text.chars().filter((c) -> c == '\n').count());
        int lastNewline  = text.lastIndexOf('\n');
        this.endColumn = lastNewline == -1 ? column + text.length() : text.length() - lastNewline;
    }

    public String getText() {
        return text;
    }

    public int getStartLine() {
        return startLine;
    }

    public int getStartColumn() {
        return startColumn;
    }

    public TokenKind getKind() {
        return kind;
    }

    public int getEndLine() {
        return endLine;
    }

    public int getEndColumn() {
        return endColumn;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Token token = (Token) o;
        return startLine == token.startLine && startColumn == token.startColumn && endLine == token.endLine && endColumn == token.endColumn && Objects.equals(text, token.text) && kind == token.kind;
    }

    @Override
    public int hashCode() {
        return Objects.hash(text, startLine, startColumn, endLine, endColumn, kind);
    }

    @Override
    public String toString() {
        return "Token{" +
                "text='" + text + '\'' +
                ", startLine=" + startLine +
                ", startColumn=" + startColumn +
                ", endLine=" + endLine +
                ", endColumn=" + endColumn +
                ", kind=" + kind +
                '}';
    }
}
