package main

import (
	"github.com/abusizhishen/jsonParser/parser"
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
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJsonParser(tokens)
	p.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(jsonListener{}, p.Json())
}

type jsonListener struct {
	*parser.BaseJsonListener
}