
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
