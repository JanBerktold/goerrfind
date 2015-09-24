package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"testing"
)

var testDir = os.Getenv("OS") + "/src/github.com/JanBerktold/goerrfind/internal/tests"

type NameTest struct {
	Query  string
	Result bool
}

var Tests = map[string]NameTest{}

func TestNameMatching(t *testing.T) {
	set := token.NewFileSet()
	f, _ := parser.ParseDir(set, testDir, nil, 0)

	for _, test := range Tests {
		fmt.Println(f, test)
	}
}
