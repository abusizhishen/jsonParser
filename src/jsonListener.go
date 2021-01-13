package src

import (
	"fmt"
	"github.com/abusizhishen/jsonParser/parser"
)

type JsonListener struct {
	*parser.BaseJsonListener
	stack []string
}

func (l *JsonListener)pop() string {
	var length = len(l.stack)

	if length == 0 {
		panic("stack is empty")
	}
	
	s := l.stack[length-1]
	l.stack = l.stack[0:length-1]
	
	return s
}

func (l *JsonListener)push(s string){
	l.stack = append(l.stack, s)
}

func (l *JsonListener)ExitARRAY(ctx *parser.ARRAYContext)  {
	fmt.Println(ctx.GetText())
}

func (l *JsonListener)ExitOBJECT(ctx *parser.OBJECTContext)  {
	fmt.Println(ctx.GetText())
}

func (l *JsonListener)ExitPair(ctx *parser.PairContext)  {
	l.push(ctx.STRING().GetText())
}

func (l *JsonListener)ExitValue(ctx *parser.ValueContext)  {
	
}

func (l *JsonListener)ExitNUMBER(ctx *parser.NUMBERContext)  {
	switch ctx.GetNu().GetTokenType() {
	case parser.JsonParserInt:
		l.push("int")
	case parser.JsonParserFloat64:
		l.push("float64")
	}
}

func (l *JsonListener)ExitSTRING(ctx *parser.STRINGContext)  {

}

func (l *JsonListener)ExitBOOL(ctx *parser.BOOLContext)  {
	fmt.Println(ctx.GetText())
}

func (l *JsonListener)ExitNULL(ctx *parser.NULLContext)  {

}
