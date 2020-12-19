package main

import (
	"fmt"
	"github.com/abusizhishen/dlang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// Setup the input
	fileStream,err := antlr.NewFileStream("test.json")
	if err != nil{
		panic(err)
	}

	// Create the Lexer
	lexer := parser.NewJsonLexer(fileStream)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
