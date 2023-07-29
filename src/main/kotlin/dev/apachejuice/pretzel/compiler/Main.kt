package dev.apachejuice.pretzel.compiler

import dev.apachejuice.pretzel.compiler.scanner.PretzelScanner
import java.io.StringReader

fun main(args: Array<String>) {
    val s = PretzelScanner(StringReader(args[0]));
    while (!s.yyatEOF()) {
        println(s.yylex())
    }
}