package ast

import (
	"fmt"
)

// A variable declaration declares a variable.
type VariableDeclaration struct {
	baseNode

	// The name of the variable to be declared
	Name *QualifiedName
	// The data type declared (inference done later)
	DeclaredType DataType
	// The initial value, if any.
	Init Expression
}

var _ Statement = &VariableDeclaration{}

func (v *VariableDeclaration) String() string {
	return fmt.Sprintf("VariableDeclaration(%s %s %s)", v.Name, v.DeclaredType, v.Init)
}

func (v *VariableDeclaration) Type() NodeType {
	return NodeVariableDeclaration
}

func (v *VariableDeclaration) Children() []Node {
	return []Node{v.Name, v.DeclaredType, v.Init}
}

func (v *VariableDeclaration) Traverse(l Listener) {
	l.EnterVariableDeclaration(v)
	v.Name.Traverse(l)

	if v.DeclaredType != nil {
		v.DeclaredType.Traverse(l)
	}

	if v.Init != nil {
		v.Init.Traverse(l)
	}

	l.ExitVariableDeclaration(v)
}

func (v *VariableDeclaration) EnterSuperclass(l Listener) {
	l.EnterNode(v)
	l.EnterStatement(v)
}

func (v *VariableDeclaration) ExitSuperclass(l Listener) {
	l.ExitStatement(v)
	l.ExitNode(v)
}

func (VariableDeclaration) IsControlFlow() bool {
	return false
}
