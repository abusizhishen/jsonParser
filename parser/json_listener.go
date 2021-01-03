// Code generated from Json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Json

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonListener is a complete listener for a parse tree produced by JsonParser.
type JsonListener interface {
	antlr.ParseTreeListener

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterNUMBER is called when entering the NUMBER production.
	EnterNUMBER(c *NUMBERContext)

	// EnterSTRING is called when entering the STRING production.
	EnterSTRING(c *STRINGContext)

	// EnterBOOL is called when entering the BOOL production.
	EnterBOOL(c *BOOLContext)

	// EnterNULL is called when entering the NULL production.
	EnterNULL(c *NULLContext)

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

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitNUMBER is called when exiting the NUMBER production.
	ExitNUMBER(c *NUMBERContext)

	// ExitSTRING is called when exiting the STRING production.
	ExitSTRING(c *STRINGContext)

	// ExitBOOL is called when exiting the BOOL production.
	ExitBOOL(c *BOOLContext)

	// ExitNULL is called when exiting the NULL production.
	ExitNULL(c *NULLContext)

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
