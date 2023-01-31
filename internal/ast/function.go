package ast

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/common/util"
)

type (
	// An argument to a function delcaration
	FunctionArg struct {
		baseNode

		// The name of the argument
		Name *QualifiedName
		// The type of the argument
		ArgType DataType
	}

	// A function declaration
	Function struct {
		baseNode

		// The name of the function
		Name *QualifiedName
		// Arguments to the function
		Args []*FunctionArg
		// The function body
		Body *Block
		// Return type
		ReturnType DataType
	}
)

var _ Node = &FunctionArg{}
var _ Node = &Function{}

func (f FunctionArg) Children() []Node {
	return []Node{f.Name, f.ArgType}
}

func (f FunctionArg) String() string {
	return fmt.Sprintf("FunctionArg(%s %s)", f.Name, f.ArgType)
}

func (f FunctionArg) Type() NodeType {
	return NodeFunctionArg
}

func (f *FunctionArg) EnterSuperclass(l Listener) {
	l.EnterNode(f)
}

func (f *FunctionArg) ExitSuperclass(l Listener) {
	l.ExitNode(f)
}

func (f *FunctionArg) Traverse(l Listener) {
	l.EnterFunctionArg(f)
	f.Name.Traverse(l)
	f.ArgType.Traverse(l)
	l.ExitFunctionArg(f)
}

func (f Function) Children() []Node {
	return append([]Node{f.Name, f.Body, f.ReturnType}, util.ArrayTransform[*FunctionArg, Node](f.Args)...)
}

func (f Function) String() string {
	return fmt.Sprintf("Function(%s %s %s)", f.Name, nodeArrayString(f.Args), f.ReturnType)
}

func (f Function) Type() NodeType {
	return NodeFunction
}

func (f *Function) EnterSuperclass(l Listener) {
	l.EnterNode(f)
}

func (f *Function) ExitSuperclass(l Listener) {
	l.ExitNode(f)
}

func (f *Function) Traverse(l Listener) {
	l.EnterFunction(f)
	f.Name.Traverse(l)
	for _, arg := range f.Args {
		arg.Traverse(l)
	}

	f.Body.Traverse(l)
	f.ReturnType.Traverse(l)
	l.ExitFunction(f)
}
