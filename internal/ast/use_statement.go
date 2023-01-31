package ast

import (
	"fmt"
)

type UseStatement struct {
	baseNode

	// The item to import.
	Imports *QualifiedName
}

var _ Statement = &UseStatement{}

func (u *UseStatement) String() string {
	return fmt.Sprintf("UseStatement(%s)", u.Imports)
}

func (u *UseStatement) Children() []Node {
	return []Node{u.Imports}
}

func (u *UseStatement) Traverse(l Listener) {
	l.EnterUseStatement(u)
	u.Imports.Traverse(l)
	l.ExitUseStatement(u)
}

func (u *UseStatement) EnterSuperclass(l Listener) {
	l.EnterNode(u)
	l.EnterStatement(u)
}

func (u *UseStatement) ExitSuperclass(l Listener) {
	l.ExitStatement(u)
	l.ExitNode(u)
}

func (u *UseStatement) IsControlFlow() bool {
	return false
}

func (UseStatement) Type() NodeType {
	return NodeUseStatement
}
