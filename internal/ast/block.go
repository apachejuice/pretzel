package ast

import (
	"fmt"

	"github.com/apachejuice/pretzel/internal/common/util"
)

type (
	// A block is a collection of code delimited by "{}".
	Block struct {
		baseNode

		// The nodes in this block.
		Nodes []Statement
	}
)

var _ Node = &Block{}

func (b *Block) String() string {
	return fmt.Sprintf("Block(%s)", nodeArrayString(b.Nodes))
}

func (b *Block) Type() NodeType {
	return NodeBlock
}

func (b *Block) Children() []Node {
	return util.ArrayTransform[Statement, Node](b.Nodes)
}

func (b *Block) Traverse(l Listener) {
	l.EnterBlock(b)
	for _, node := range b.Nodes {
		node.Traverse(l)
	}

	l.ExitBlock(b)
}

func (b *Block) EnterSuperclass(l Listener) {
	l.EnterNode(b)
}

func (b *Block) ExitSuperclass(l Listener) {
	l.ExitNode(b)
}
