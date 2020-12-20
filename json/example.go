package json

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Run() {
	// Setup the input
	fileStream,err := antlr.NewFileStream("test.json")
	if err != nil{
		panic(err)
	}

	// Create the Lexer
	lexer := NewJsonLexer(fileStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewJsonParser(tokens)
	parser.BuildParseTrees = true
	tree := parser.Init()
	tree.ToStringTree()
	tree.ToStringTree(parser)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

}
