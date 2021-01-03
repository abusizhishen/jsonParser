// Code generated from Json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Json

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJsonListener is a complete listener for a parse tree produced by JsonParser.
type BaseJsonListener struct{}

var _ JsonListener = &BaseJsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseJsonListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseJsonListener) ExitNumber(ctx *NumberContext) {}

// EnterNUMBER is called when production NUMBER is entered.
func (s *BaseJsonListener) EnterNUMBER(ctx *NUMBERContext) {}

// ExitNUMBER is called when production NUMBER is exited.
func (s *BaseJsonListener) ExitNUMBER(ctx *NUMBERContext) {}

// EnterSTRING is called when production STRING is entered.
func (s *BaseJsonListener) EnterSTRING(ctx *STRINGContext) {}

// ExitSTRING is called when production STRING is exited.
func (s *BaseJsonListener) ExitSTRING(ctx *STRINGContext) {}

// EnterBOOL is called when production BOOL is entered.
func (s *BaseJsonListener) EnterBOOL(ctx *BOOLContext) {}

// ExitBOOL is called when production BOOL is exited.
func (s *BaseJsonListener) ExitBOOL(ctx *BOOLContext) {}

// EnterNULL is called when production NULL is entered.
func (s *BaseJsonListener) EnterNULL(ctx *NULLContext) {}

// ExitNULL is called when production NULL is exited.
func (s *BaseJsonListener) ExitNULL(ctx *NULLContext) {}

// EnterARRAY is called when production ARRAY is entered.
func (s *BaseJsonListener) EnterARRAY(ctx *ARRAYContext) {}

// ExitARRAY is called when production ARRAY is exited.
func (s *BaseJsonListener) ExitARRAY(ctx *ARRAYContext) {}

// EnterOBJECT is called when production OBJECT is entered.
func (s *BaseJsonListener) EnterOBJECT(ctx *OBJECTContext) {}

// ExitOBJECT is called when production OBJECT is exited.
func (s *BaseJsonListener) ExitOBJECT(ctx *OBJECTContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseJsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseJsonListener) ExitPair(ctx *PairContext) {}

// EnterObject is called when production object is entered.
func (s *BaseJsonListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseJsonListener) ExitObject(ctx *ObjectContext) {}

// EnterArray is called when production array is entered.
func (s *BaseJsonListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseJsonListener) ExitArray(ctx *ArrayContext) {}

// EnterJson is called when production json is entered.
func (s *BaseJsonListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseJsonListener) ExitJson(ctx *JsonContext) {}
