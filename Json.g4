grammar Json;

String : '"' ([a-zA-Z0-9\\" '])* '"';
Number : [0-9]+;
WS :[\t\r\n ]+  ->skip;
True:'true';
False:'false';
Null:'null';

bool:True|False;
value :
    Number
      |String
      |array
      |object
      |bool
      |Null
        ;

pair:
    String ':' value;

object :
    '{' '}'
    |'{' pair (','pair)* '}'
;

array:
    '[' ']'
    |'[' value (','value)*']'
    ;

init :
    array
    |object;

