package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
