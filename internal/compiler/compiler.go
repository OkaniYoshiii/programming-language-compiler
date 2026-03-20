package compiler

import (
	"github.com/okaniyoshiii/programming-language/internal/bytecode"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

const (
	SourceFile string = "source_file"
	VarStatement string = "var_statement"
	ConstStatement string = "const_statement"
	Identifier string = "identifier"
	Keyword string = "keyword"
	ConstantExpression = "constant_expression"
)

func Compile(source []byte, tree *tree_sitter.Node) (error, []byte) {
	builder := bytecode.BytecodeBuilder{}

	cursor := tree.Walk()
	defer cursor.Close()

	if cursor.Node().Kind() == SourceFile {
		cursor.GotoFirstChild()
	}

	for {
		parent := cursor.Node()
		switch parent.Kind() {
			// (var_statement (var_keyword) (identifier) (expression))
			case VarStatement:
				identifier := parent.NamedChild(1)
				_ = string(source[identifier.StartByte():identifier.EndByte()])

				expression := parent.NamedChild(2)
				_ = CompileExpression(source, &builder, expression)
		}

		if !cursor.GotoNextSibling() {
			break
		}
	}

	return nil, builder.Build()
}

func CompileExpression(source []byte, builder *bytecode.BytecodeBuilder, expression *tree_sitter.Node) (error) {
	switch expression.Kind() {
		case ConstantExpression:
			litteral := string(source[expression.StartByte():expression.EndByte()])
			index := builder.AddConstant([]byte(litteral))

			instruction := bytecode.MakeOpConst(uint32(index))
			builder.AddInstruction(instruction.OpCode, instruction.Operands)
	}

	return nil
}
