package ast

import "fmt"

// A variable reference.
type VariableReference struct {
	baseNode
	baseExpression

	// The variable that is referenced
	Reference *QualifiedName
}

var _ Expression = &VariableReference{}

func (v VariableReference) Children() []Node {
	return []Node{v.Reference}
}

func (VariableReference) EType() ExpressionType {
	return ExpressionVariableReference
}

func (v *VariableReference) EnterSuperclass(l Listener) {
	l.EnterNode(v)
	l.EnterExpression(v)
}

func (v *VariableReference) ExitSuperclass(l Listener) {
	l.ExitExpression(v)
	l.ExitNode(v)
}

func (v *VariableReference) Traverse(l Listener) {
	l.EnterVariableReference(v)
	v.Reference.Traverse(l)
	l.ExitVariableReference(v)
}

func (VariableReference) HasOperator() bool {
	return false
}

func (VariableReference) IsConstant() bool {
	return false
}

func (VariableReference) IsPure() bool {
	return true
}

func (VariableReference) Operands() int {
	return 1
}

func (v VariableReference) String() string {
	return fmt.Sprintf("VariableReference(%s)", v.Reference)
}
