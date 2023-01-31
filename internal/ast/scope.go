package ast

import "fmt"

// The item type of a name in a scope.
type ScopeRefType int

const (
	ScopeRefVariable ScopeRefType = iota
	ScopeRefFunction
)

// A scope that contains names that can be referenced from child scopes.
type Scope struct {
	// The names.
	Names map[*QualifiedName]ScopeRefType
	// Possible parent scope.
	Parent *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Names:  make(map[*QualifiedName]ScopeRefType),
		Parent: parent,
	}
}

func (s Scope) canFind(name *QualifiedName) (ScopeRefType, bool) {
	for key, value := range s.Names {
		if key.Matches(name) {
			return value, true
		}
	}

	return -1, false
}

// Is the given name findable in this scope?. Return the ref type if so.
func (s Scope) Exists(name *QualifiedName) (ScopeRefType, bool) {
	if val, ok := s.canFind(name); ok {
		return val, true
	}

	if s.Parent != nil {
		return s.Parent.Exists(name)
	}

	return -1, false
}

// Add the given name to the scope.
func (s Scope) Add(name *QualifiedName, refType ScopeRefType) {
	s.Names[name] = refType
}

func (s Scope) Print() {
	for key, val := range s.Names {
		fmt.Printf("Scope: %s to %d\n", key, val)
	}

	if s.Parent != nil {
		s.Parent.Print()
	}
}
