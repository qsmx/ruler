%{

package ruler

%}

%union {
    value   string

    ruler   *rulerNode
}

%token  INT, FLOAT, TRUE, FALSE, STRING, IDENT, ARRAYS
%token  EQU, NEQ, LSS, LEQ, GTR, GEQ
%token  AND, OR, IN, NOTIN
%token  FUNC, ARGS, FILTER

%left   OR
%left   AND
%left   EQU NEQ
%left   LSS LEQ GTR GEQ
%left   '-' '+'
%left   '*' '/' '%'
%left   NEG     /* negation--unary minus */
%right  '^'     /* exponentiation */

%type   <value>     INT, FLOAT, STRING, IDENT, ARRAYS
%type   <ruler>     Int, Float, String, Ident, Arrays
%type   <ruler>     Expr, ExprExt
%type   <ruler>     ExprBool, BoolCell, Bool
%type   <ruler>     ExprCalc, CalcCell
%type   <ruler>     ExprFilter, FilterName
%type   <ruler>     ExprFunc, Args

%%

Input:    /* empty */ { }
    | Expr            { if outNode != nil { *outNode = $1 } }

Expr: ExprBool
    | ExprExt

ExprBool: Bool
    | ExprFunc
    | ExprFilter
    | BoolCell NEQ BoolCell     { $$ = nodeFromNode(NEQ, $1, $3) }
    | BoolCell EQU BoolCell     { $$ = nodeFromNode(EQU, $1, $3) }
    | ExprCalc LSS ExprCalc     { $$ = nodeFromNode(LSS, $1, $3) }
    | ExprCalc LEQ ExprCalc     { $$ = nodeFromNode(LEQ, $1, $3) }
    | ExprCalc GTR ExprCalc     { $$ = nodeFromNode(GTR, $1, $3) }
    | ExprCalc GEQ ExprCalc     { $$ = nodeFromNode(GEQ, $1, $3) }
    | ExprBool AND ExprBool     { $$ = nodeFromNode(AND, $1, $3) }
    | ExprBool OR ExprBool      { $$ = nodeFromNode(OR , $1, $3) }
    | '(' ExprBool ')'          { $$ = $2 }

BoolCell: Bool
    | ExprCalc

ExprCalc: CalcCell
    | ExprCalc '+' ExprCalc     { $$ = nodeFromNode('+', $1, $3) }
    | ExprCalc '-' ExprCalc     { $$ = nodeFromNode('-', $1, $3) }
    | ExprCalc '*' ExprCalc     { $$ = nodeFromNode('*', $1, $3) }
    | ExprCalc '/' ExprCalc     { $$ = nodeFromNode('/', $1, $3) }
    | ExprCalc '%' ExprCalc     { $$ = nodeFromNode('%', $1, $3) }
    | '-' ExprCalc %prec NEG    { $$ = nodeFromNode(NEG, $2) }
    | '(' ExprCalc ')'          { $$ = $2 }

CalcCell: ExprExt
    | ExprFunc

ExprExt: Ident
    | Arrays
    | Int
    | Float
    | String

ExprFunc: IDENT '(' ')'         { $$ = nodeFromString(FUNC, $1) }
    | IDENT '(' Args ')'        { $$ = nodeFromString(FUNC, $1); $$.ReplaceNode($3) }

Args: BoolCell                  { $$ = nodeFromNode(ARGS, $1) }
    | Args ',' BoolCell         { $$ = $$.AppendNode($3) }

ExprFilter: FilterName '[' ExprBool ']'  { $$ = nodeFromNode(FILTER, $1); $$.AppendNode($3) }

FilterName: Ident
    | Arrays

Ident: IDENT            { $$ = nodeFromIdent(IDENT, $1) }
Arrays: ARRAYS          { $$ = nodeFromIdent(ARRAYS, $1) }
Int: INT                { $$ = nodeFromInt(ConvertInt($1)) }
Float: FLOAT            { $$ = nodeFromFloat(ConvertFloat($1)) }
String: STRING          { $$ = nodeFromString(STRING, $1) }
Bool: TRUE              { $$ = nodeFromBool(true)  }
    | FALSE             { $$ = nodeFromBool(false) }

%%
