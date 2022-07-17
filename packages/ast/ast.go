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

type Ident struct {
	NamePos token.Pos
	Name    string
}

type Var struct {
	Name *Ident
	Type Expr
}

func (v Var) Pos() token.Pos {
	return v.Name.NamePos
}

func (v Var) End() token.Pos {
	return v.Type.End()
}

func (v Var) stmtNode() {}

// DefStmt def foo(a int) a+1 end
type DefStmt struct {
	Func *Proto
	Body *Block
}

func (d DefStmt) Pos() token.Pos {
	return d.Func.Pos()
}

func (d DefStmt) End() token.Pos {
	return d.Body.End()
}

func (d DefStmt) exprNode() {}

type Proto struct {
	Name      *Ident
	LParenPos token.Pos
	RParenPos token.Pos
	Args      []*Var
}

func (p Proto) Pos() token.Pos {
	return p.Name.NamePos
}

func (p Proto) End() token.Pos {
	//if len(p.Args) > 0 {
	//	return p.Args[len(p.Args)-1].End()
	//}
	return p.RParenPos
}

func (p Proto) exprNode() {}

type Block struct {
	DoPos  token.Pos
	Stmts  []Stmt
	EndPos token.Pos
}

func (b Block) Pos() token.Pos {
	return b.DoPos
}

func (b Block) End() token.Pos {
	return b.EndPos
}

func (b Block) exprNode() {}

type Call struct {
	Callee *Proto
}

func (c Call) Pos() token.Pos {
	return c.Callee.Pos()
}

func (c Call) End() token.Pos {
	return c.Callee.RParenPos
}

func (c Call) exprNode() {}

type ExprStmt struct {
	Expr Expr
}

func (s ExprStmt) Pos() token.Pos {
	return s.Pos()
}

func (s ExprStmt) End() token.Pos {
	return s.End()
}

func (s ExprStmt) stmtNode() {}
