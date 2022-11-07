grammar Neon;

program
    : ((stat | func))* EOF
    ;

// complicated stuff
stat : 
    // function calls, assign, aka no curlies
    ((
        decl
        | assign
        | funccall
        | return
    )';')
    |
    // stuff with curlies
    (if | while)
    ;

if  : IF '(' expr ')' '{' stat* '}' elif* else?;
elif: ELIF '(' expr ')' '{' stat* '}';
else: ELSE '{' stat* '}';
while  : WHILE '(' expr ')' '{' stat* '}';
func : DEF type ID ('(' (funcarg (',' funcarg)*)? ')')? '{' stat* '}';
funccall : ID '(' (expr (',' expr)*)? ')'; // allow stuff like func()() later, where func() returns a function

funcarg :
     'number' ID #NumArg
     | 'string' ID #StrArg
     | 'bool' ID #BoolArg
     ;

decl : 'number' assign #NumVar
     | 'string' assign #StrVar
     | 'bool' assign #BoolVar
     ;
assign: var=ID '=' expr;
return: RETURN expr;

expr : expr op=MDM expr #MDM
     | expr op=ADD_SUB expr #AddSub
     | expr op=COMP expr #Comparison
     | '!' expr #NotExpr
     | funccall #FuncCall
     | INT #Int
     | BOOL #Bool
     | STRING #String
     | ID #Identifier
     | '(' expr ')' #NoClue
     ;

ADD_SUB : '+' | '-';
MDM : '*' | '/' | '%';
COMP : '==' | '<=' | '>=' | '<' | '>';
// grouping
type    : 'number' | 'string' | 'bool' | 'void';
BOOL    : 'true' | 'false';
// place all special identifiers here
DEF      : 'def';
RETURN   : 'return';
IF       : 'if';
ELIF     : 'elif';
ELSE     : 'else';
WHILE    : 'while';

// regular stuff
STRING : '"' (' '..'~')* '"';
ID     : ('a'..'z'|'A'..'Z')+;
INT    : '0'..'9'+;
WS     : [ \t\n\r]+ -> skip ;