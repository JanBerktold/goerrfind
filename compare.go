package main

import (
	"go/ast"
	"strings"
)

func structNameEqual(name string, list []*ast.Field) bool {
	for _, field := range list {
		if w, ok := field.Type.(*ast.StarExpr); ok {
			if ident, ok := w.X.(*ast.Ident); ok {
				return name == ident.Name
			}
		}
	}
	return false
}

func funcNameEqual(name string, fun *ast.FuncDecl) bool {
	if i := strings.Index(name, "."); i > 0 {
		structName := name[0:i]
		methodName := name[i+1 : len(name)]
		return fun.Recv != nil &&
			fun.Name.Name == methodName &&
			structNameEqual(structName, fun.Recv.List)
	}
	return fun.Name.Name == name && fun.Recv == nil
}

func paramIndexByType(f *ast.FuncType, t string) int {
	for index, field := range f.Results.List {
		if ident, ok := field.Type.(*ast.Ident); ok {
			if ident.Name == t {
				return index
			}
		}
	}
	panic("attempt to search missing type")
}
