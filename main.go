package main

import (
	"go/parser"
	"go/token"
	"fmt"
	"os"
	"go/ast"
	"reflect"
)

type WriteVisitor struct {}

func (w *WriteVisitor) Visit(node ast.Node) ast.Visitor {
	fmt.Println(reflect.ValueOf(node).Elem().Kind())
	return nil
}

func main() {
	set := token.NewFileSet()

	f, err := parser.ParseDir(set, os.Getenv("GOPATH") + "/src/github.com/JanBerktold/JarvisClient/console", nil, 0)
	fmt.Println(err)

	for _, v := range f {
		fmt.Println(v.Scope)
		ast.Walk(new(WriteVisitor), v)
	}
}

