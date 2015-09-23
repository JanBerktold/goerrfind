package main

import (
	"go/ast"
	"testing"
)

type NameTest struct {
	Name   string
	Func   *ast.Ident
	Result bool
}

var Tests = []NameTest{
	NameTest{
		Name: "main",
		Func: &ast.Ident{
			Name: "main",
		},
		Result: true,
	},
}

func TestNameMatching(t *testing.T) {
	for _, test := range Tests {
		if funcNameEqual(test.Name, test.Func) != test.Result {
			t.Fatalf("Test with name %q failed. Expected: %v", test.Name, test.Result)
		}
	}
}
