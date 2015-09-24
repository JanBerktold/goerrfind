package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
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
			// case: return x, abc(), d
			panic("not implemented yet")
		default:
			realIndex++
		}
		resultIndex++
	}
	return realIndex
}

func (w *ReturnFinder) logResult(r string, p token.Pos) {
	fmt.Printf("%s %v\n", r, p)
}

// return types
// 	const Exported (done)
// 	const Not exported
//	var n-assignments
//	var 1-assignment (exported?)
//	call expr
func (w *ReturnFinder) handleStatement(st ast.Stmt) {
	switch stat := st.(type) {
	case *ast.ReturnStmt:
		i := matchingExpression(stat, w.pIndex)
		expr := stat.Results[i]
		switch ex := expr.(type) {
		case *ast.Ident:
			if ex.Obj != nil {
				switch ex.Obj.Kind {
				case ast.Var:
					fmt.Println("VAR", reflect.TypeOf(ex.Obj.Decl), ex.Obj.Decl)
					if assign, ok := ex.Obj.Decl.(*ast.AssignStmt); ok {
						fmt.Println(assign)
					}
				case ast.Con:
					if value, ok := ex.Obj.Decl.(*ast.ValueSpec); ok {
						if value.Names[0].IsExported() {
							w.logResult(value.Names[0].Name, stat.Pos())
						} else {
							panic("not implemented yet")
						}
					}
				default:
					panic("not implemented obj.Kind")
				}
			} else {
				// really only nil?
				fmt.Println("IDENT", ex)
			}
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
		if w.rType == "" {
			w.pIndex = paramIndexByType(t.Type, w.rType)
		}
		for _, st := range t.Body.List {
			w.handleStatement(st)
		}
		return nil
	}
	return w
}

func NewReturnFinderByType(name, typ string) *ReturnFinder {
	return &ReturnFinder{
		fName: name,
		rType: typ,
	}
}

func NewReturnFinderByPos(name string, pos int) *ReturnFinder {
	return &ReturnFinder{
		fName:  name,
		pIndex: pos,
	}
}

func workFuncMethod(set *token.FileSet, dir, typ, name string) {
	f, _ := parser.ParseDir(set, dir, nil, 0)
	finder := NewReturnFinderByType(name, typ)

	for _, v := range f {
		ast.Walk(finder, v)
	}
}

var (
	packageFlag = flag.String("pkg", "", "Package to search for func")
)

func printUsageInfo() {
	os.Exit(1)
}

//usage: goerrfind main
//usage: goerrfind ReturnFinder.Visit
//usage: goerrfind -pkg=github.com/JanBerktold/Hello ReturnFinder.Visit
func main() {
	flag.Parse()
	set := token.NewFileSet()

	directory := ""
	if *packageFlag != "" {
		directory = fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), *packageFlag)
	} else {
		if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
			panic(err)
		} else {
			directory = dir
		}
	}

	if len(flag.Args()) != 1 {
		printUsageInfo()
	}

	workFuncMethod(set, directory, "error", flag.Args()[0])
}
