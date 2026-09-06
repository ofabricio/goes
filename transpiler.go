package main

import (
	"fmt"
	"io"
	"strings"
)

type Transpiler struct {
}

func (t *Transpiler) Generate(p Program, w io.Writer) {
	t.generate(p, 0, w)
}

func (t *Transpiler) generate(v any, depth int, out io.Writer) {
	switch v := v.(type) {
	case Program:
		t.generate(v.Package, depth, out)
		for _, v := range v.Imports {
			t.generate(v, depth, out)
		}
		t.generate("\n", depth, out)
		for _, v := range v.Functions {
			t.generate(v, depth, out)
		}
		for _, v := range v.TypeDefs {
			t.generate(v, depth, out)
		}
	case Package:
		t.generate(fmt.Sprintf("package %s\n\n", v.Name.Text), depth, out)
	case Import:
		t.generate(fmt.Sprintf("import \"%s\"\n", strings.TrimPrefix(v.Path.Text, "go:")), depth, out)
	case Function:
		t.generate(fmt.Sprintf("func %s() {\n", v.Name.Text), depth, out)
		t.generate("}\n", depth, out)
	case TypeDef:
		t.generate(fmt.Sprintf("\ntype %s struct{}\n", v.Name.Text), depth, out)
	default:
		fmt.Fprintf(out, "%s%v", strings.Repeat("    ", depth), v)
	}
}
