package ast

import (
	"fmt"
	"unicode/utf8"

	"github.com/apachejuice/pretzel/internal/errors"
)

type (
	// An enum representing node types.
	NodeType int

	// Represents the most basic of nodes at the bottom of the tree having some central properties amongst nodes.
	// This AST does not use a visitor pattern, use a Listener instead.
	Node interface {
		fmt.Stringer

		// The type of this node.
		Type() NodeType
		// Returns this node's start position.
		Begin() errors.SourceContext
		// Returns this node's end position.
		End() errors.SourceContext
		// Returns the child nodes this node has.
		Children() []Node
		// Traverse this node with the listener. The node functions as the walker, so no reflection is needed.
		Traverse(l Listener)
		// The error reporting context of this node. Should be the same reference for all nodes from the same file.
		ReportContext() *errors.ReportContext
		// This node's scope.
		Scope() *Scope
		// For subclasses of nodes, this method calls the superclass Enter methods. Same for Exit methods below.
		EnterSuperclass(l Listener)
		ExitSuperclass(l Listener)
	}

	baseNode struct {
		start, end errors.SourceContext
		reportCtx  *errors.ReportContext
		scope      *Scope
	}
)

func (b baseNode) Begin() errors.SourceContext {
	return b.start
}

func (b baseNode) End() errors.SourceContext {
	return b.end
}

func (b baseNode) ReportContext() *errors.ReportContext {
	return b.reportCtx
}

func (b baseNode) Scope() *Scope {
	return b.scope
}

func (b baseNode) rangeStr() string {
	return fmt.Sprintf("%d:%d-%d:%d", b.start.StartLine, b.start.StartColumn, b.end.EndLine, b.end.EndColumn)
}

// utf-8...
func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func nodeArrayString[N Node](nodes []N) string {
	if len(nodes) == 0 {
		return "[]"
	}

	buf := "["
	for _, n := range nodes {
		buf += n.String() + " "
	}

	return trimLastChar(buf) + "]"
}

// Constants for NodeType
const (
	NodeName NodeType = iota
	NodeQualifiedName
	NodeRootNode
	NodeUseStatement
	NodeDataType
	NodeFunctionArg
	NodeFunction
	NodeBlock
	NodeExpression
	NodeExpressionStatement
	NodeVariableDeclaration

	NodeUnknown
)

var nodeTypeNames map[NodeType]string = map[NodeType]string{
	NodeName:                "name",
	NodeQualifiedName:       "qualified name",
	NodeRootNode:            "root",
	NodeUseStatement:        "use statement",
	NodeDataType:            "data type",
	NodeFunctionArg:         "function argument",
	NodeFunction:            "function declaration",
	NodeBlock:               "block",
	NodeExpressionStatement: "expression statement",

	NodeUnknown: "unknown",
}

var _ fmt.Stringer = NodeType(0)

func (n NodeType) String() string {
	return nodeTypeNames[n]
}
