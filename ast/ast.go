package ast

import (
	"github.com/PonzyPon/monkey/token"
)

// Node は、AST(Abstract syntax tree:抽象構文木)の要素を表します。
type Node interface {
	TokenLiteral() string
}

// Statement は、文を表します。文とは一般に手続きを表します。
type Statement interface {
	Node
	statementNode()
}

// Expression は、式を表します。式とは、一般に評価された値を持ちます。
type Expression interface {
	Node
	expressionNode()
}

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

type LetStatement struct {
	Token token.Token //token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
