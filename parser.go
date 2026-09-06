package main

import (
	"fmt"
	"strings"

	nom "github.com/ofabricio/nom"
)

type Parser struct {
	nom.Parser
}

func (p *Parser) Parse(src string) (Program, error) {
	p.Parser = nom.New(src)
	return p.Program(), p.Parser.Err
}

func (p *Parser) Program() Program {
	var g Program
	var v AST
	if p.WS() && p.Package(&v) {
		g.Package = v.(Package)
	}
	for p.WS() && p.Import(&v) {
		g.Imports = append(g.Imports, v.(Import))
	}
	for p.WS() && p.More() {
		var v AST
		switch {
		case p.Function(&v):
			g.Functions = append(g.Functions, v.(Function))
		case p.TypeDef(&v):
			g.TypeDefs = append(g.TypeDefs, v.(TypeDef))
		default:
			p.Expected(fmt.Sprintf("end of program, but found: %s", p.Char()))
			return g
		}
	}
	return g
}

func (p *Parser) Package(out *AST) bool {
	var a Package
	if p.Expect("pkg") && p.HS() && (p.MatchOut(nom.WORD, &a.Name) || p.Expected("package name")) {
		*out = a
		return true
	}
	return false
}

func (p *Parser) Import(out *AST) bool {
	var a Import
	if p.Match("use") && p.HS() && (p.Out(p.Mark(), p.String(`"`), &a.Path) || p.Expected("package path")) {
		a.Path.Text = strings.Trim(a.Path.Text, `"`)
		*out = a
		return true
	}
	return false
}

func (p *Parser) Function(out *AST) bool {
	var v Function
	if p.Match("func") &&
		p.HS() && p.MatchOut(nom.WORD, &v.Name) &&
		p.HS() && p.Match("(") && p.FunctionArgs() && p.Match(")") &&
		p.HS() && p.Match("{") && p.FunctionBody() && p.WS() && p.Expect("}") {
		*out = v
		return true
	}
	return false
}

func (p *Parser) TypeDef(out *AST) bool {
	var v TypeDef
	if p.Match("type") &&
		p.HS() && p.MatchOut(nom.WORD, &v.Name) &&
		p.HS() && p.Match("(") && p.WS() && p.Expect(")") {
		*out = v
		return true
	}
	return false
}

func (p *Parser) FunctionArgs() bool {
	return true
}

func (p *Parser) FunctionBody() bool {
	return true
}

func (p *Parser) HS() bool {
	return p.Optional(nom.ST)
}

func (p *Parser) WS() bool {
	return p.Optional(nom.WS)
}
