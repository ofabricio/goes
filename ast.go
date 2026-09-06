package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/ofabricio/nom"
)

type AST any

type Program struct {
	Package   Package
	Imports   []Import
	Functions []Function
	TypeDefs  []TypeDef
}

type Package struct {
	Name nom.Token
}

type Import struct {
	Path nom.Token
}

type Function struct {
	Name nom.Token
	Body []AST
}

type TypeDef struct {
	Name nom.Token
}

type AstPrinter struct{}

func (p *AstPrinter) Print(a AST, out io.Writer) {
	p.print(a, 0, out)
}

func (p *AstPrinter) print(a AST, depth int, out io.Writer) {
	switch v := a.(type) {
	case Program:
		p.print(v.Package, depth, out)
		p.print(" [\n", depth, out)
		for _, f := range v.Imports {
			p.print(f, depth+1, out)
		}
		for _, f := range v.Functions {
			p.print(f, depth+1, out)
		}
		for _, f := range v.TypeDefs {
			p.print(f, depth+1, out)
		}
		p.print("]\n", depth, out)
	case Package:
		p.print(fmt.Sprintf("Package: %s", v.Name.Text), depth, out)
	case Import:
		p.print(fmt.Sprintf("Import: \"%s\"\n", v.Path.Text), depth, out)
	case Function:
		p.print(fmt.Sprintf("Function: %s [\n", v.Name.Text), depth, out)
		for _, stmt := range v.Body {
			p.print(stmt, depth+1, out)
		}
		p.print("]\n", depth, out)
	case TypeDef:
		p.print(fmt.Sprintf("Type: %s\n", v.Name.Text), depth, out)
	default:
		fmt.Fprintf(out, "%s%v", strings.Repeat("    ", depth), v)
	}
}
