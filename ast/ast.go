package ast

// Node is the base node interface in the AST.
type Node interface {
	TokenLiteral() string
}

// Statement is the interface for all statement nodes.
// Statements are nodes that do not produce a value but do have side effects.
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface for all expression nodes.
// Expressions are nodes that produce a value.
type Expression interface {
	Node
	expressionNode()
}

// The root node of every AST the parser produces.
// Every valid Monkey program is a series of statements contained in Program.Statements
// which is just a slice of AST nodes that implement the Statement interface.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
