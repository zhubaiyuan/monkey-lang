package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func quote(node ast.Node) object.Object {
	node = evalUnquoteCalls(node)
	return &object.Quote{Node: node}
}

func evalUnquoteCalls(quoted ast.Node) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquoteCall(node) {
			return node
		}
		call, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}
		if len(call.Arguments) != 1 {
			return node
		}
		return node
	})
}

func isUnquoteCall(node ast.Node) bool {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}
	return callExpression.Function.TokenLiteral() == "unquote"
}
