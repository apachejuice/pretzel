package ast

import (
	"fmt"
)

type (
	// Represents a single code file.
	RootNode struct {
		baseNode
		elements []Node // the top-level declarations in the file
	}
)

var _ Node = &RootNode{}

func (r RootNode) Children() []Node {
	return r.elements
}

func (r *RootNode) Traverse(l Listener) {
	l.EnterRootNode(r)
	for _, child := range r.elements {
		child.Traverse(l)
	}

	l.ExitRootNode(r)
}

func (r *RootNode) EnterSuperclass(l Listener) {
	l.EnterNode(r)
}

func (r *RootNode) ExitSuperclass(l Listener) {
	l.ExitNode(r)
}

func (r *RootNode) String() string {
	return fmt.Sprintf("RootNode(%s %s)", r.rangeStr(), nodeArrayString(r.elements))
}

func (RootNode) Type() NodeType {
	return NodeRootNode
}
