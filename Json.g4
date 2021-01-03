grammar Json;

STRING
    : '"' [a-zA-Z_0-9 ]* '"';
Int : [0-9]+;
Float64 : [0-9]+ '.' [0-9]+;

number  //Number是错误写法 语法只能小写，大写无法识别
    :Int
    |Float64
    ;

WS
    :[\t\r\n ]+ ->skip;

value
    :nu=(Int|Float64) #NUMBER
    |STRING #STRING
    |('true'|'false') #BOOL
    |'null' #NULL
    |array#ARRAY
    |object#OBJECT
    ;

pair:
    STRING ':' value;

object
    :'{'  pair ( ',' pair )* '}'
    |'{' '}'
    ;

array
    :'[' value ( ',' value )* ']'
    |'[' ']'
    ;

json
    :object
    |array
    ;