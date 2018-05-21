%{

package ruler

import (
    "fmt"
)

%}

%union{
    value           interface{}

    value_int       int64
    value_bool      bool
    value_string    string
    value_float     float64
    value_slice     []interface{}
}

%token  INT, FLOAT, TRUE, FALSE, STRING, IDENT, ARRAYS, FILTER
%token  EQU, NEQ, LSS, LEQ, GTR, GEQ
%token  AND, OR, IN, NOTIN

%left   OR
%left   AND
%left   EQU NEQ
%left   LSS LEQ GTR GEQ
%left   '-' '+'
%left   '*' '/'
%left   NEG     /* negation--unary minus */
%right  '^'     /* exponentiation */

%type   <value_string>  INT, FLOAT, STRING, IDENT, ARRAYS, FILTER
%type   <value_string>  Expr, Filter
%type   <value>         FilterData
%type   <value_bool>   Condition

%% /* The grammar follows.  */

Input:    /* empty */ { }
    | Expr        		{ fmt.Println($1) }

Expr: INT           { fmt.Println("INT:", $1) }
    | FLOAT         { fmt.Println("FLOAT:", $1) }
    | STRING        { fmt.Println("STRING:", $1) }
    | IDENT         { fmt.Println("IDENT:", $1) }
    | ARRAYS        { fmt.Println("ARRAYS:", $1) }
    | Filter        { fmt.Println("FILTER:", $1) }
;

Filter: FilterData '[' Condition ']' {
        fmt.Printf("FilterData: %+v\n", $1)
        for _, v := range []int{1, 2, 3} {
            fmt.Println( "[",  $3 , "]", v)
        }
        $$ = "1"
    }

Condition: IDENT NEQ INT         {
        fmt.Println("sub:", ($1 + "!=" + $3))
        $$ = dp.GetInt($1) != ConvertInt($3)
    }

FilterData: IDENT           { $$ = dp.GetAttr($1); dp.Next($1); fmt.Println("data:", $1); }
    | ARRAYS                { $$ = $1; fmt.Println("Arrs:", $1); }

%%

func Parse(src string, data map[string]interface{}) bool {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			return
		}
	}()
    var res bool
    rulerParse(newLexer(src), NewDataPackage(data), &res)
    fmt.Println("res:", res)
    return res
}
