package ast

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/common/util"
)

type (
	// Categories of datatypes.
	DataTypeType int

	// A data type.
	DataType interface {
		Node

		// The type of the datatype
		DType() DataTypeType
	}

	// A base data type.
	baseDataType int

	// An atom type, having a qualified name.
	AtomType struct {
		baseNode
		baseDataType

		// The reference to the type this type identifies.
		Reference *QualifiedName
	}

	// An array type.
	// Arrays of arrays are not collapsed.
	ArrayType struct {
		baseNode
		baseDataType

		// The underlying type.
		Inner DataType
	}

	// A generic type having a variable amount of type arguments.
	GenericType struct {
		baseNode
		baseDataType

		// The qualified name representing the main type
		Reference *QualifiedName
		// A list of type arguments
		Args []DataType
	}

	// A nullable type assignable to the nothing value.
	NullableType struct {
		baseNode
		baseDataType

		// The underlying type
		Inner DataType
	}

	// No type at all
	NoneType struct {
		baseNode
		baseDataType
	}
)

func (baseDataType) Type() NodeType {
	return NodeDataType
}

var _ fmt.Stringer = DataTypeType(0)

func (d DataTypeType) HasOtherTypes() bool {
	return d == DataTypeArray || d == DataTypeGeneric
}

func (d DataTypeType) String() string {
	return dataTypeTypeNames[d]
}

var dataTypeTypeNames map[DataTypeType]string = map[DataTypeType]string{
	DataTypeAtom:     "atom",
	DataTypeArray:    "array",
	DataTypeGeneric:  "generic",
	DataTypeNullable: "nullable",
	DataTypeNone:     "none",
}

const (
	DataTypeAtom DataTypeType = iota
	DataTypeArray
	DataTypeGeneric
	DataTypeNullable
	DataTypeNone
)

// Data type implementations: AtomType

var _ DataType = &AtomType{}

func (a AtomType) DType() DataTypeType {
	return DataTypeAtom
}

func (a AtomType) Children() []Node {
	return []Node{a.Reference}
}

func (a AtomType) String() string {
	return fmt.Sprintf("AtomType(%s)", a.Reference)
}

func (a *AtomType) EnterSuperclass(l Listener) {
	l.EnterNode(a)
	l.EnterDataType(a)
}

func (a *AtomType) ExitSuperclass(l Listener) {
	l.ExitDataType(a)
	l.ExitNode(a)
}

func (a *AtomType) Traverse(l Listener) {
	l.EnterAtomType(a)
	a.Reference.Traverse(l)
	l.ExitAtomType(a)
}

// ArrayType

var _ DataType = &ArrayType{}

func (a ArrayType) DType() DataTypeType {
	return DataTypeArray
}

func (a ArrayType) Children() []Node {
	return []Node{a.Inner}
}

func (a ArrayType) String() string {
	return fmt.Sprintf("ArrayType(%s)", nodeArrayString(a.Children()))
}

func (a *ArrayType) EnterSuperclass(l Listener) {
	l.EnterNode(a)
	l.EnterDataType(a)
}

func (a *ArrayType) ExitSuperclass(l Listener) {
	l.ExitDataType(a)
	l.ExitNode(a)
}

func (a *ArrayType) Traverse(l Listener) {
	l.EnterArrayType(a)
	a.Inner.Traverse(l)
	l.ExitArrayType(a)
}

// GenericType

func (g GenericType) DType() DataTypeType {
	return DataTypeGeneric
}

func (g GenericType) Children() []Node {
	return append([]Node{g.Reference}, util.ArrayTransform[DataType, Node](g.Args)...)
}

func (g GenericType) String() string {
	return fmt.Sprintf("GenericType(%s %s)", g.Reference, nodeArrayString(g.Args))
}

func (g *GenericType) EnterSuperclass(l Listener) {
	l.EnterNode(g)
	l.EnterDataType(g)
}

func (g *GenericType) ExitSuperclass(l Listener) {
	l.ExitDataType(g)
	l.ExitNode(g)
}

func (g *GenericType) Traverse(l Listener) {
	l.EnterGenericType(g)
	g.Reference.Traverse(l)
	for _, child := range g.Args {
		child.Traverse(l)
	}

	l.ExitGenericType(g)
}

// NullableType

func (n NullableType) DType() DataTypeType {
	return DataTypeGeneric
}

func (n NullableType) Children() []Node {
	return []Node{n.Inner}
}

func (n NullableType) String() string {
	return fmt.Sprintf("NullableType(%s)", n.Inner)
}

func (n *NullableType) EnterSuperclass(l Listener) {
	l.EnterNode(n)
	l.EnterDataType(n)
}

func (n *NullableType) ExitSuperclass(l Listener) {
	l.ExitDataType(n)
	l.ExitNode(n)
}

func (n *NullableType) Traverse(l Listener) {
	l.EnterNullableType(n)
	n.Inner.Traverse(l)
	l.ExitNullableType(n)
}

// NoneType

func (n NoneType) DType() DataTypeType {
	return DataTypeNone
}

func (n NoneType) Children() []Node {
	return []Node{}
}

func (n NoneType) String() string {
	return "NoneType()"
}

func (n *NoneType) EnterSuperclass(l Listener) {
	l.EnterNode(n)
	l.EnterDataType(n)
}

func (n *NoneType) ExitSuperclass(l Listener) {
	l.ExitDataType(n)
	l.ExitNode(n)
}

func (n *NoneType) Traverse(l Listener) {
	l.EnterNoneType(n)
	l.ExitNoneType(n)
}
