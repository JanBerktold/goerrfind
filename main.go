package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type ReturnFinder struct {
	fName  string
	rType  string
	pIndex int
}

func matchingExpression(stat *ast.ReturnStmt, pIndex int) int {
	realIndex := 0
	resultIndex := 0
	for realIndex < pIndex {
		switch stat.Results[resultIndex].(type) {
		case *ast.CallExpr:
			panic("not implemented yet")
		default:
			realIndex++
		}
		resultIndex++
	}
	return realIndex
}

func (w *ReturnFinder) handleStatement(st ast.Stmt) {
	switch stat := st.(type) {
	case *ast.ReturnStmt:
		i := matchingExpression(stat, w.pIndex)
		expr := stat.Results[i]
		switch ex := expr.(type) {
		case *ast.Ident:
			fmt.Println("IDENT", ex)
		case *ast.CallExpr:
			fmt.Println("CALL", ex)
		}
	case *ast.SwitchStmt:
		for _, block := range stat.Body.List {
			if clause, ok := block.(*ast.CaseClause); ok {
				if clause.Body != nil {
					for _, statement := range clause.Body {
						w.handleStatement(statement)
					}
				}
			}
		}
	case *ast.IfStmt:
		for _, s := range stat.Body.List {
			w.handleStatement(s)
		}
		if stat.Else != nil {
			w.handleStatement(stat.Else)
		}
	case *ast.BlockStmt:
		for _, s := range stat.List {
			w.handleStatement(s)
		}
	}
}

func (w *ReturnFinder) Visit(node ast.Node) ast.Visitor {
	if t, ok := node.(*ast.FuncDecl); ok && funcNameEqual(w.fName, t) {
		w.pIndex = paramIndexByType(t.Type, w.rType)
		for _, st := range t.Body.List {
			w.handleStatement(st)
		}
		return nil
	}
	return w
}

func NewReturnFinder(name, typ string) *ReturnFinder {
	return &ReturnFinder{
		fName: name,
		rType: typ,
	}
}

//usage: goerrfind main
//usage: goerrfind ReturnFinder.Visit
//usage: goerrfind github.com/JanBerktold/Hello ReturnFinder.Visit
func main() {
	set := token.NewFileSet()

	dir := os.Getenv("GOPATH") + "/src/"
	fun := ""
	switch len(os.Args) {
	case 3:
		dir += os.Args[1]
		fun = os.Args[2]
	case 2:
		fun = os.Args[1]
	default:
		fmt.Println("some usage info")
	}

	f, _ := parser.ParseDir(set, dir, nil, 0)

	for _, v := range f {
		ast.Walk(NewReturnFinder(fun, "error"), v)
	}
}
