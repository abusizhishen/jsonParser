// Code generated from json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // json

import "github.com/antlr/antlr4/runtime/Go/antlr"

// jsonListener is a complete listener for a parse tree produced by jsonParser.
type jsonListener interface {
	antlr.ParseTreeListener

	// EnterSingleValue is called when entering the SingleValue production.
	EnterSingleValue(c *SingleValueContext)

	// EnterARRAY is called when entering the ARRAY production.
	EnterARRAY(c *ARRAYContext)

	// EnterOBJECT is called when entering the OBJECT production.
	EnterOBJECT(c *OBJECTContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// ExitSingleValue is called when exiting the SingleValue production.
	ExitSingleValue(c *SingleValueContext)

	// ExitARRAY is called when exiting the ARRAY production.
	ExitARRAY(c *ARRAYContext)

	// ExitOBJECT is called when exiting the OBJECT production.
	ExitOBJECT(c *OBJECTContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)
}
