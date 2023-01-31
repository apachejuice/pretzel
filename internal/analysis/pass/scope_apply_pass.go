package pass

import (
	"github.com/apachejuice/pretzel/internal/ast"
	"github.com/apachejuice/pretzel/internal/errors"
)

// A pass for applying variables to scopes.
// Also brings in scopes from other files to build a tree.
type ScopeApplyPass struct {
	pass
	scope *ast.Scope
}

var _ Pass = &ScopeApplyPass{}

func NewScopeApplyPass() *ScopeApplyPass {
	return &ScopeApplyPass{
		pass:  newPass(),
		scope: nil,
	}
}

func (ScopeApplyPass) Errors() []*errors.Error {
	return []*errors.Error{}
}

func (pass *ScopeApplyPass) EnterRootNode(r *ast.RootNode) {
	pass.scope = r.Scope()
}

func (pass *ScopeApplyPass) EnterVariableDeclaration(v *ast.VariableDeclaration) {
	pass.scope.Add(v.Name, ast.ScopeRefVariable)
}

func (pass *ScopeApplyPass) EnterFunction(f *ast.Function) {
	pass.scope.Add(f.Name, ast.ScopeRefFunction)
}

func (pass *ScopeApplyPass) ExitVariableDeclaration(v *ast.VariableDeclaration) {
	v.ExitSuperclass(pass)
}

func (pass *ScopeApplyPass) ExitFunction(f *ast.Function) {
	f.ExitSuperclass(pass)
}
