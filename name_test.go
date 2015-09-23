package main

import (
	"go/ast"
	"testing"
)

type NameTest struct {
	Query  string
	Func   *ast.FuncDecl
	Result bool
}

var Tests = []NameTest{
	NameTest{
		Query: "main",
		Func: &ast.FuncDecl{
			Name: &ast.Ident{
				Name:    "main",
				NamePos: 20,
			},
		},
		Result: true,
	},
	NameTest{
		Query: "help",
		Func: &ast.FuncDecl{
			Name: &ast.Ident{
				Name:    "main",
				NamePos: 20,
			},
		},
		Result: false,
	},
	/*NameTest{
		Query: "Router.Get",
		Func: &ast.FuncDecl{
			Name: &ast.Ident{
				Name:    "Get",
				NamePos: 20,
			},
			Type: &ast.FuncType{
				Func: 40,
				Params: &ast.FieldList{
					List: []*ast.Field{
					},
				},
			},
		},
		Result: false,
	},*/
}

func TestNameMatching(t *testing.T) {
	for _, test := range Tests {
		if funcNameEqual(test.Query, test.Func) != test.Result {
			t.Fatalf("Test with query %q failed. Expected: %v", test.Query, test.Result)
		}
	}
}
