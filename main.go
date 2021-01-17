package main

import (
	"github.com/abusizhishen/jsonParser/parser"
	"github.com/abusizhishen/jsonParser/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	fileStream,err := antlr.NewFileStream("test.json")
	if err != nil{
		panic(err)
	}

	// Create the Lexer
	lexer := parser.NewjsonLexer(fileStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewjsonParser(tokens)
	p.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(&src.JsonListener{}, p.Json())
}