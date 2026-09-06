package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLanguage(t *testing.T) {

	src, err := os.ReadFile("testdata/program.goes")
	if err != nil {
		panic(err)
	}

	tgt, err := os.ReadFile("testdata/program.go")
	if err != nil {
		panic(err)
	}

	cmp, err := os.ReadFile("testdata/program.goes.ast")
	if err != nil {
		panic(err)
	}

	par := Parser{}
	prg, err := par.Parse(string(src))
	assert(t, fmt.Sprint(err), "<nil>", "compilation error")

	var out strings.Builder

	pri := AstPrinter{}
	pri.Print(prg, &out)

	got := out.String()
	exp := string(cmp)
	assert(t, got, exp, "compiled AST does not match target")

	out.Reset()

	trn := Transpiler{}
	trn.Generate(prg, &out)

	got = out.String()
	exp = string(tgt)
	assert(t, got, exp, "generated code does not match target")
}

func assert(t *testing.T, got, exp, msg string) {
	if got != exp {
		t.Errorf("\nMsg: %s\nGot:\n%v\nExp:\n%v\n", msg, got, exp)
	}
}
