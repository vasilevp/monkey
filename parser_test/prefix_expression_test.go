package parser_test

import (
	"testing"

	"git.exsdev.ru/ExS/monkey/ast"
	"git.exsdev.ru/ExS/monkey/lexer"
	"git.exsdev.ru/ExS/monkey/parser"
	"github.com/stretchr/testify/require"
)

func TestPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		expected interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"-b;", "-", "b"},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		require.Len(t, program.Statements, 1)
		require.IsType(t, new(ast.ExpressionStatement), program.Statements[0])

		expr := program.Statements[0].(*ast.ExpressionStatement)

		require.IsType(t, new(ast.PrefixExpression), expr.Expression)
		prefix := expr.Expression.(*ast.PrefixExpression)
		testLiteralExpression(t, prefix.Right, tt.expected)
	}
}
