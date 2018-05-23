%{
package ruler

type yylexer struct{
    src     string
    buf     []byte
    empty   bool
    current byte
    size    int
    pos     int
}

func (y *yylexer) getc() byte {
    if y.current != 0 {
        y.buf = append(y.buf, y.current)
    }
    y.current = 0
    if y.pos < y.size {
        y.current = y.src[y.pos]
        y.pos++
    }

    return y.current
}

func (y yylexer) Error(string) {
    throwException(ERR_STATUS_ABORT, "ruler[%s] buf[%s] pos[%d]", y.src, string(y.buf), y.pos)
}

func (y *yylexer) Lex(lval *yySymType) int {
    c := y.current
    if y.empty {
        c, y.empty = y.getc(), false
    }
%}

%yyc c
%yyn c = y.getc()

arrays          {identifier}*(\.({letter}|{digit})+)+
identifier      \|?{letter}({letter}|{digit})*
ident           {letter}({letter}|{digit})*
int             {digit}+
float           {int}\.{int}|\.{int}
string          \"{stringcell}\"|\"{stringcell}(\\\"{stringcell})*\"

digit           [0-9]
letter          [_a-zA-Z]
stringcell      [^"]*

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

{string}        lval.value = string(y.buf[1:len(y.buf)-1]); return STRING
{int}           lval.value = string(y.buf); return INT
{float}         lval.value = string(y.buf); return FLOAT
{identifier}    lval.value = string(y.buf); return IDENT
{arrays}        lval.value = string(y.buf); return ARRAYS

%%
    y.empty = true
    if len(y.buf) != 0 {
        y.Error("")
    }
    return int(c)
}

func newLexer(src string) (y *yylexer) {
    y = &yylexer{
        src: src,
        size: len(src),
        pos: 1,
        current: src[0],
    }
    return
}
