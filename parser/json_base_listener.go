// Code generated from json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // json

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasejsonListener is a complete listener for a parse tree produced by jsonParser.
type BasejsonListener struct{}

var _ jsonListener = &BasejsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingleValue is called when production SingleValue is entered.
func (s *BasejsonListener) EnterSingleValue(ctx *SingleValueContext) {}

// ExitSingleValue is called when production SingleValue is exited.
func (s *BasejsonListener) ExitSingleValue(ctx *SingleValueContext) {}

// EnterARRAY is called when production ARRAY is entered.
func (s *BasejsonListener) EnterARRAY(ctx *ARRAYContext) {}

// ExitARRAY is called when production ARRAY is exited.
func (s *BasejsonListener) ExitARRAY(ctx *ARRAYContext) {}

// EnterOBJECT is called when production OBJECT is entered.
func (s *BasejsonListener) EnterOBJECT(ctx *OBJECTContext) {}

// ExitOBJECT is called when production OBJECT is exited.
func (s *BasejsonListener) ExitOBJECT(ctx *OBJECTContext) {}

// EnterPair is called when production pair is entered.
func (s *BasejsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BasejsonListener) ExitPair(ctx *PairContext) {}

// EnterObject is called when production object is entered.
func (s *BasejsonListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BasejsonListener) ExitObject(ctx *ObjectContext) {}

// EnterArray is called when production array is entered.
func (s *BasejsonListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BasejsonListener) ExitArray(ctx *ArrayContext) {}

// EnterJson is called when production json is entered.
func (s *BasejsonListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BasejsonListener) ExitJson(ctx *JsonContext) {}
