package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

//https://www.cnblogs.com/double12gzh/p/13632267.html

func main() {
	expr := `a == 1 && b == 2`
	fset := token.NewFileSet()
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Print(fset, exprAst)

}
