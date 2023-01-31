package ast

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/common/util"
)

type QualifiedName struct {
	baseNode

	// A list of names making up the qualified name.
	// Even a single name is enough, as all references to things are resolved as QualifiedNames.
	Names []*Name
}

var _ Node = &QualifiedName{}

func (q *QualifiedName) String() string {
	return fmt.Sprintf("QualifiedName(%s)", nodeArrayString(q.Names))
}

func (q *QualifiedName) Children() []Node {
	return util.ArrayTransform[*Name, Node](q.Names) // weird go generic quirks
}

func (q *QualifiedName) Traverse(l Listener) {
	l.EnterQualifiedName(q)
	for _, name := range q.Names {
		name.Traverse(l)
	}
	l.ExitQualifiedName(q)
}

func (q *QualifiedName) EnterSuperclass(l Listener) {
	l.EnterNode(q)
}

func (q *QualifiedName) ExitSuperclass(l Listener) {
	l.ExitNode(q)
}

func (QualifiedName) Type() NodeType {
	return NodeQualifiedName
}

func (q QualifiedName) Matches(other *QualifiedName) bool {
	if len(q.Names) == len(other.Names) {
		for i := 0; i < len(q.Names); i++ {
			if q.Names[i].Name != other.Names[i].Name {
				return false
			}
		}

		return true
	}

	return false
}
