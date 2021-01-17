grammar json;

STRING
    : '"' [a-zA-Z_0-9 ]* '"';
Number  //Number是错误写法 语法只能小写，大写无法识别
    :'-'? ([0-9]+'.')? [0-9]+
    ;

WS
    :[\t\r\n ]+ ->skip;

value
    :(STRING|Number|'null'|'false'|'true') #SingleValue
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