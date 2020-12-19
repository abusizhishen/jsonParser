grammar Json;

DQM: '"'; //DoubleQuotationMarks
Key : DQM[a-zA-Z_]+DQM;
Str : DQM .*? DQM;
Int : [0-9]+;
WS :[\t\r\n ]+  ->skip;

value :Int
      |Key
      |Str
      |array
      |object
       ;

keyValue:Str ':' value;

object : '{}'
    |'{'keyValue'}'
    |'{'keyValue','keyValue+'}'
    ;

array:'[]'
    |'[' value ']'
    |'[' value+ ']'
    ;

init : array
    |object;

