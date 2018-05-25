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
%token  FUNC, ARGS, FILTER, LISTA, LISTO

%left   OR
%left   AND
%left   EQU NEQ LSS LEQ GTR GEQ
%left   '-' '+'
%left   '*' '/' '%'
%left   NEG     /* negation--unary minus */
%right  '^'     /* exponentiation */

%type   <value>     INT, FLOAT, STRING, IDENT, ARRAYS
%type   <ruler>     Int, Float, String, Ident, Arrays
%type   <ruler>     Expr, ExprExt
%type   <ruler>     ExprBool, BoolCell, Bool
%type   <ruler>     ExprCalc, CalcCell
%type   <ruler>     ExprList, ExprFilter, FilterCell, FilterName
%type   <ruler>     ExprFunc, Args, InArgs

%%

Input:    /* empty */ { }
    | Expr            { if outNode != nil { *outNode = $1 } }

Expr: ExprBool
    | ExprExt
    | ExprList

ExprList: FilterCell
    | ExprFilter

ExprBool: Bool
    | ExprFunc
    | Ident IN '(' InArgs ')'       { $$ = nodeFromString(IN, "i"); $$.AppendNode($1, $4) }
    | Ident NOTIN '(' InArgs ')'    { $$ = nodeFromString(IN, "n"); $$.AppendNode($1, $4) }
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

InArgs: ExprExt                 { $$ = nodeFromNode(ARGS, $1) }
    |InArgs ',' ExprExt         { $$ = $$.AppendNode($3) }

ExprExt: Ident
    | Arrays
    | Int
    | Float
    | String

ExprFunc: IDENT '(' ')'         { $$ = nodeFromString(FUNC, $1) }
    | IDENT '(' Args ')'        { $$ = nodeFromString(FUNC, $1); $$.ReplaceNode($3) }

Args: BoolCell                  { $$ = nodeFromNode(ARGS, $1) }
    | Args ',' BoolCell         { $$ = $$.AppendNode($3) }

ExprFilter: FilterCell AND ExprBool   { $$ = nodeFromNode(LISTA, $3); $$.AppendNode($1) }
    | FilterCell OR ExprBool    { $$ = nodeFromNode(LISTO, $3); $$.AppendNode($1) }
    | ExprBool AND FilterCell   { $$ = nodeFromNode(LISTA, $1); $$.AppendNode($3) }
    | ExprBool OR FilterCell    { $$ = nodeFromNode(LISTO, $1); $$.AppendNode($3) }
    | '(' ExprFilter ')'        { $$ = $2 }

FilterCell: FilterName '[' ExprBool ']'  { $$ = nodeFromNode(FILTER, $1); $$.AppendNode($3) }

FilterName: Ident
    | Arrays

Ident: IDENT            { $$ = nodeFromIdent(IDENT, $1) }
    | IN                { $$ = nodeFromIdent(IDENT, "in") }
    | NOTIN             { $$ = nodeFromIdent(IDENT, "notin") }
Arrays: ARRAYS          { $$ = nodeFromIdent(ARRAYS, $1) }
Int: INT                { $$ = nodeFromInt(ConvertInt($1)) }
Float: FLOAT            { $$ = nodeFromFloat(ConvertFloat($1)) }
String: STRING          { $$ = nodeFromString(STRING, $1) }
Bool: TRUE              { $$ = nodeFromBool(true)  }
    | FALSE             { $$ = nodeFromBool(false) }

%%
