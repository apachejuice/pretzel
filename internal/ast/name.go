package ast

import "fmt"

// A single part of a name.
type Name struct {
	baseNode

	// The identifier making up the name.
	Name string
}

var _ Node = &Name{}

func (n *Name) String() string {
	return fmt.Sprintf("Name(%s)", n.Name)
}

func (n *Name) Children() []Node {
	return []Node{}
}

func (n *Name) Traverse(l Listener) {
	l.EnterName(n)
	l.ExitName(n)
}

func (n *Name) EnterSuperclass(l Listener) {
	l.EnterNode(n)
}

func (n *Name) ExitSuperclass(l Listener) {
	l.ExitNode(n)
}

func (Name) Type() NodeType {
	return NodeName
}
