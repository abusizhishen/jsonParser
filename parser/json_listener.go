// Code generated from Json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Json

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonListener is a complete listener for a parse tree produced by JsonParser.
type JsonListener interface {
	antlr.ParseTreeListener

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)
}
