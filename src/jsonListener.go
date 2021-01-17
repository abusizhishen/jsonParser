package src

import (
	"fmt"
	"github.com/abusizhishen/jsonParser/parser"
	"strings"
)

type JsonListener struct {
	*parser.BasejsonListener
	stack []string
	level int
	perform Perform
}

type Perform int
const (
	normal = iota
	compress
	format
	space = "  "
	newLine = '\n'
)

type Type int
const (
	Map Type = iota
	Array
	Value
)

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

//func (l *JsonListener)EnterARRAY(ctx *parser.ARRAYContext)  {
//	l.pl(leftF)
//	l.level++
//}
//
//func (l *JsonListener)ExitARRAY(ctx *parser.ARRAYContext)  {
//	l.level--
//	l.pl(rightF)
//}
//
//func (l *JsonListener)EnterOBJECT(ctx *parser.OBJECTContext)  {
//	l.pl(leftH)
//	l.level++
//}
//func (l *JsonListener)ExitOBJECT(ctx *parser.OBJECTContext)  {
//	l.level--
//	l.pl(rightH)
//}

func (l *JsonListener)EnterArray(ctx *parser.ArrayContext)  {
	l.level++
}

func (l *JsonListener)ExitArray(ctx *parser.ArrayContext)  {
	count := (ctx.GetChildCount()-1)/2
	var pairs []string

	for i:=0;i<count;i++{
		value := l.pop()
		pairs = append(pairs, l.p(value))
	}

	s := strings.Builder{}


	s.WriteString(l.p("[\n"))
	s.WriteString(strings.Join(pairs, ",\n")+"\n")
	s.WriteString("]\n")

	l.push(s.String())
	l.level--
}

func (l *JsonListener)EnterObject(ctx *parser.ObjectContext)  {
	l.level++
}
func (l *JsonListener)ExitObject(ctx *parser.ObjectContext)  {

	count := (ctx.GetChildCount()-1)/2
	var pairs []string

	for i:=0;i<count;i++{
		value,key := l.pop(),l.pop()
		pairs = append(pairs, l.p(fmt.Sprintf("%s%s%s", key, ":", value)))
	}

	s := strings.Builder{}
	s.WriteString("{\n")
	s.WriteString(strings.Join(pairs, ",\n")+"\n")
	s.WriteString("}\n")
	l.push(s.String())

	l.level--
}

func (l *JsonListener)EnterPair(ctx *parser.PairContext)  {
	l.push(ctx.STRING().GetText())
}

func (l *JsonListener)ExitPair(ctx *parser.PairContext)  {}

func (l *JsonListener)ExitValue(ctx *parser.ValueContext)  {
	
}

func (l *JsonListener)EnterJson(ctx *parser.JsonContext)  {
}

func (l *JsonListener)ExitJson(ctx *parser.JsonContext)  {
	for _,s := range l.stack{
		fmt.Print(s)
	}
}

func (l *JsonListener)ExitSingleValue(ctx *parser.SingleValueContext)  {
	l.push(ctx.GetText())
}

func (l *JsonListener)p(text string) string{
	spaces := strings.Repeat(space,l.level)
	return fmt.Sprintf("%s%s",spaces, text)
}