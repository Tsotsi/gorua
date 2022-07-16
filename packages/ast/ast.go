package ast

import "xlin/learn-llvm/packages/token"

type Node interface {
	Pos() token.Pos
	End() token.Pos
}

type Expr interface {
	Node
	exprNode()
}
type Stmt interface {
	Node
	stmtNode()
}

type DefStmt struct {
	Name  string
	Value interface{}
}
