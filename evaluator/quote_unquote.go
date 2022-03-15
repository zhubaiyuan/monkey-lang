package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	return &object.Quote{Node: node}
}
