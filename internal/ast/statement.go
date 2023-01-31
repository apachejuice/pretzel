package ast

type (
	// A statement is a code element that cannot be used as a value. Declarations are also statements.
	// Some things that may look like statements, such as `raise` or `return` expressions are not statements.
	// This is because those things are guaranteed to end whatever is going on and as such can be relied on to produce no value.
	Statement interface {
		Node

		// Returns true if this statement is a control flow statement.
		IsControlFlow() bool
	}
)
