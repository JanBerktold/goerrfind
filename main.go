package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func structNameEqual(name string, list []*ast.Field) bool {
	for _, field := range list {
		if w, ok := field.Type.(*ast.StarExpr); ok {
			return name == w.X
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

type ReturnFinder struct {
	fName string
}

func (w *ReturnFinder) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.FuncDecl:
		if funcNameEqual(w.fName, t) {
			fmt.Println(w.fName, "HELLU")
			for _, st := range t.Body.List {
				switch st.(type) {
				case *ast.ReturnStmt:
					fmt.Println(st)
				}
			}
			return nil
		}
	}
	return w
}

func NewReturnFinder(name string) *ReturnFinder {
	return &ReturnFinder{
		fName: name,
	}
}

func main() {
	set := token.NewFileSet()

	f, _ := parser.ParseDir(set, os.Getenv("GOPATH")+"/src/github.com/JanBerktold/goerrfind", nil, 0)

	for _, v := range f {
		ast.Walk(NewReturnFinder("ReturnFinder.Visit"), v)
	}
}
