%{
package ruler

import (
    "bufio"
    "fmt"
    "bytes"
)

type yylexer struct{
    src     *bufio.Reader
    buf     []byte
    empty   bool
    current byte
    pos     int
}

func (y *yylexer) getc() byte {
    if y.current != 0 {
        y.buf = append(y.buf, y.current)
    }
    y.current = 0
    if b, err := y.src.ReadByte(); err == nil {
        y.current = b
        y.pos++
    }
    return y.current
}

func (y yylexer) Error(e string) {
    fmt.Printf("pos: %v '%c'\n", y.pos, y.current)
}

func (y *yylexer) Lex(lval *rulerSymType) int {
    c := y.current
    if y.empty {
        c, y.empty = y.getc(), false
    }
%}

%yyc c
%yyn c = y.getc()

identifier      {letter}({letter}|{digit})*
int             {digit}+
float           {int}\.{int}|\.{int}
string          \"{stringcell}\"|\"{stringcell}(\\\"{stringcell})*\"

digit           [0-9]
letter          [_a-zA-Z]
stringcell      [^"]*
arrays          {identifier}*(\.({letter}|{digit})+)+

%%
    y.buf = y.buf[:0]

[ \t\r]+

"=="            return EQU
"!="            return NEQ
">="            return GEQ
">"             return GTR
"<="            return LEQ
"<"             return LSS
"&&"            return AND
"||"            return OR

true            return TRUE
false           return FALSE

in              return IN
notin           return NOTIN

{string}        lval.value_string = string(y.buf[1:len(y.buf)-1]); return STRING
{int}           lval.value_string = string(y.buf); return INT
{float}         lval.value_string = string(y.buf); return FLOAT
{identifier}    lval.value_string = string(y.buf); return IDENT
{arrays}        lval.value_string = string(y.buf); return ARRAYS

%%
    y.empty = true
    return int(c)
}

func newLexer(str string) (y *yylexer) {
    src := bufio.NewReader(bytes.NewBufferString(str))
    y = &yylexer{src: src}
    if b, err := src.ReadByte(); err == nil {
        y.current = b
    }
    return
}
