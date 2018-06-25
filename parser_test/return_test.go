package parser_test

import (
	"testing"

	"git.exsdev.ru/ExS/gop/ast"
	"git.exsdev.ru/ExS/gop/lexer"
	"git.exsdev.ru/ExS/gop/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)

	assert.Equal(t, 3, len(program.Statements), "program should have 3 statements")

	for _, stmt := range program.Statements {
		assert.Equal(t, stmt.TokenLiteral(), "return")
		assert.IsType(t, new(ast.ReturnStatement), stmt)
	}
}
