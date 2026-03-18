package main

import (
	"bufio"
	"fmt"
	"os"

	pl "github.com/okaniyoshiii/tree-sitter-programming-language/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func main() {
	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(pl.Language()))

	fmt.Println("Programming Language REPL :")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		line := reader.Bytes()

		tree := parser.Parse(line, nil)
		defer tree.Close()

		root := tree.RootNode()
		fmt.Println(root.ToSexp())
	}
}
