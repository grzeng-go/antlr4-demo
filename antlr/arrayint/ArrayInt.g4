grammar ArrayInt;
init: '{' value (',' value)* '}';
value: init
       | INT
       ;
INT: [0-9]+;
WS: [ \t\n\r]+ -> skip;