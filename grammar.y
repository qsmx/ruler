%{

package ruler

%}

%union {
    value   string
    ruler   *node
}

%token  INT, INTHEX, FLOAT, TRUE, FALSE, STRING, IDENT, ARRAY
%token  EQU, NEQ, LSS, LEQ, GTR, GEQ
%token  AND, OR, IN, NOTIN, BITAND, BITOR
%token  FUNC, ARGS, FILTER, LISTA, LISTO

%left   OR
%left   AND
%left   BITOR
%left   BITAND
%left   EQU NEQ LSS LEQ GTR GEQ
%left   '-' '+'
%left   '*' '/' '%'
%left   NEG     /* negation--unary minus */
%right  '^'     /* exponentiation */

%type   <value>     INT, FLOAT, TRUE, FALSE, STRING, IDENT, ARRAY, IN, NOTIN
%type   <value>     IDENTS
%type   <ruler>     Int, Ident, Float, Bool, String
%type   <ruler>     CalcCell, BoolCell, FuncArgs
%type   <ruler>     Expr, ExprBool, ExprCalc, ExprFunc, ExprFilter

%%

Input: /* empty */      { }
    | Expr              { *outNode = $1 }

Expr: ExprBool
    | ExprFilter

ExprBool: Bool
    | ExprFunc
    | Ident IN '(' FuncArgs ')'     { $$ = nodeFromNode2(IN, "i", $1, $4) }
    | Ident NOTIN '(' FuncArgs ')'  { $$ = nodeFromNode2(IN, "n", $1, $4) }
    | ExprBool AND ExprBool     { $$ = nodeFromNode(AND, $1, $3) }
    | ExprBool OR  ExprBool     { $$ = nodeFromNode(OR , $1, $3) }
    | BoolCell EQU BoolCell     { $$ = nodeFromNode(EQU, $1, $3) }
    | BoolCell NEQ BoolCell     { $$ = nodeFromNode(NEQ, $1, $3) }
    | ExprCalc LSS ExprCalc     { $$ = nodeFromNode(LSS, $1, $3) }
    | ExprCalc LEQ ExprCalc     { $$ = nodeFromNode(LEQ, $1, $3) }
    | ExprCalc GTR ExprCalc     { $$ = nodeFromNode(GTR, $1, $3) }
    | ExprCalc GEQ ExprCalc     { $$ = nodeFromNode(GEQ, $1, $3) }
    | '(' ExprBool ')'          { $$ = $2 }

ExprFilter: '[' ExprBool ']'    { $$ = nodeFromNode(FILTER, nodeFromIdent(IDENT, DEFAULT_SLICE_LABEL), $2) }
    | Ident '[' ExprBool ']'    { $$ = nodeFromNode(FILTER, $1, $3) }
    | ExprBool AND ExprFilter   { $$ = nodeFromNode(LISTA, $1, $1) }
    | ExprBool OR  ExprFilter   { $$ = nodeFromNode(LISTO, $1, $1) }
    | ExprFilter AND ExprBool   { $$ = nodeFromNode(LISTA, $3, $1) }
    | ExprFilter OR  ExprBool   { $$ = nodeFromNode(LISTO, $3, $1) }
    | '(' ExprFilter ')'        { $$ = $2 }

ExprFunc: IDENTS '(' ')'        { $$ = nodeFromFunc(FUNC, $1) }
    | IDENTS '(' FuncArgs ')'   { $$ = nodeFromFunc(FUNC, $1, $3) }

FuncArgs: BoolCell              { $$ = nodeFromNode(ARGS, $1) }
    | FuncArgs ',' BoolCell     { $$.AppendNode($3) }

ExprCalc: CalcCell
    | ExprCalc '+' ExprCalc     { $$ = nodeFromNode('+', $1, $3) }
    | ExprCalc '-' ExprCalc     { $$ = nodeFromNode('-', $1, $3) }
    | ExprCalc '*' ExprCalc     { $$ = nodeFromNode('*', $1, $3) }
    | ExprCalc '/' ExprCalc     { $$ = nodeFromNode('/', $1, $3) }
    | ExprCalc '%' ExprCalc     { $$ = nodeFromNode('%', $1, $3) }
    | ExprCalc BITAND ExprCalc  { $$ = nodeFromNode(BITAND, $1, $3) }
    | ExprCalc BITOR ExprCalc   { $$ = nodeFromNode(BITOR, $1, $3) }
    | '-' ExprCalc %prec NEG    { $$ = nodeFromNode(NEG, $2) }
    | '(' ExprCalc ')'          { $$ = $2 }

BoolCell: Bool | ExprCalc
CalcCell: Int | Float | String | Ident | ExprFunc

Int: INT                { $$ = nodeFromInt("", ConvertInt($1)) }
Float: FLOAT            { $$ = nodeFromFloat("", ConvertFloat($1)) }
String: STRING          { $$ = nodeFromString("", $1) }
Bool: TRUE              { $$ = nodeFromBool("", true) }
    | FALSE             { $$ = nodeFromBool("", false) }
Ident: IDENTS           { $$ = nodeFromIdent(IDENT, $1) }
    | ARRAY             { $$ = nodeFromIdent(ARRAY, $1) }

IDENTS: IDENT
    | IN
    | NOTIN
%%
