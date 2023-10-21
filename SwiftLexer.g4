lexer grammar SwiftLexer;

// Tokens
//tipos
INT : 'Int';
FLOAT : 'Float';
BOOL : 'Bool';
STR : 'String';

//palabras reservadas
TRU: 'true';
FAL: 'false';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
FOR: 'for';
IN: 'in';
VAR: 'var';
LET: 'let';
NIL: 'nil';
BREAK: 'break';
CONTINUE: 'continue';
APPEND: 'append';
REMOVELAST: 'removeLast';
REMOVE: 'remove';
AT: 'at';
ISEMPTY: 'isEmpty';
COUNT: 'count';
ARRAY: 'array';
RETURN: 'return';
FUNC: 'func';
STRUCT: 'struct';
GUARD: 'guard';

//primitivos
NUMBER: [0-9]+ ('.' [0-9]+)?;
STRING: '"'~["]*'"';
CHAR: '\''~[']'\'';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

//Simbolos
DIFE: '!=';
IG_IG: '==';
NOT: '!';
OR: '||';
AND: '&&';
IGUAL: '=';
MAYIG: '>=';
MENIG: '<=';
MAYOR: '>';
MENOR: '<';
MULT: '*';
DIV: '/';
SUM: '+';
RES: '-';
MOD: '%';
PAR_IZQ: '(';
PAR_DER: ')';
LLAVE_IZQ: '{';
LLAVE_DER: '}';
DOSPUNTOS: ':';
COR_IZQ: '[';
COR_DER: ']';
COMA: ',';
CIERRAPREGUNTA: '?';
PUNTOCOMA: ';';
PUNTO: '.';
FLECHA: '->';

//Saltar
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;