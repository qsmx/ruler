// Code generated by golex. DO NOT EDIT.

package ruler

type yylexer struct {
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
	throwException(E_RULER_SYNTAX, "ruler[%s] buf[%s] pos[%d]", y.src, string(y.buf), y.pos)
}

func (y *yylexer) Lex(lval *yySymType) int {
	c := y.current
	if y.empty {
		c, y.empty = y.getc(), false
	}

yystate0:

	y.buf = y.buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = y.getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '&':
		goto yystate9
	case c == '.':
		goto yystate11
	case c == '0':
		goto yystate18
	case c == '<':
		goto yystate27
	case c == '=':
		goto yystate29
	case c == '>':
		goto yystate31
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c == 'f':
		goto yystate36
	case c == 'i':
		goto yystate41
	case c == 'n':
		goto yystate43
	case c == 't':
		goto yystate48
	case c == '|':
		goto yystate52
	case c >= '1' && c <= '9':
		goto yystate24
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c == 'g' || c == 'h' || c >= 'j' && c <= 'm' || c >= 'o' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate33
	}

yystate2:
	c = y.getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate4
	}

yystate4:
	c = y.getc()
	goto yyrule3

yystate5:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate5
	}

yystate6:
	c = y.getc()
	goto yyrule16

yystate7:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate8
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate5
	}

yystate8:
	c = y.getc()
	switch {
	default:
		goto yyrule16
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate5
	}

yystate9:
	c = y.getc()
	switch {
	default:
		goto yyrule10
	case c == '&':
		goto yystate10
	}

yystate10:
	c = y.getc()
	goto yyrule8

yystate11:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '0':
		goto yystate12
	case c >= '1' && c <= '9':
		goto yystate15
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate14
	}

yystate12:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate13
	case c == 'X' || c == 'x':
		goto yystate16
	case c >= '0' && c <= '9':
		goto yystate15
	case c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate14
	}

yystate13:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate14
	}

yystate14:
	c = y.getc()
	switch {
	default:
		goto yyrule20
	case c == '.':
		goto yystate13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate14
	}

yystate15:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate13
	case c >= '0' && c <= '9':
		goto yystate15
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate14
	}

yystate16:
	c = y.getc()
	switch {
	default:
		goto yyrule20
	case c == '.':
		goto yystate13
	case c == '_':
		goto yystate14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate17:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == '.':
		goto yystate13
	case c == '_':
		goto yystate14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate17
	}

yystate18:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate19
	case c == 'X' || c == 'x':
		goto yystate25
	case c >= '0' && c <= '9':
		goto yystate24
	}

yystate19:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '0':
		goto yystate20
	case c >= '1' && c <= '9':
		goto yystate21
	}

yystate20:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c == 'X' || c == 'x':
		goto yystate22
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate21:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate22:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate23
	}

yystate23:
	c = y.getc()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate23
	}

yystate24:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate24
	}

yystate25:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate26
	}

yystate26:
	c = y.getc()
	switch {
	default:
		goto yyrule17
	case c == '.':
		goto yystate19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate26
	}

yystate27:
	c = y.getc()
	switch {
	default:
		goto yyrule7
	case c == '=':
		goto yystate28
	}

yystate28:
	c = y.getc()
	goto yyrule6

yystate29:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate30
	}

yystate30:
	c = y.getc()
	goto yyrule2

yystate31:
	c = y.getc()
	switch {
	default:
		goto yyrule5
	case c == '=':
		goto yystate32
	}

yystate32:
	c = y.getc()
	goto yyrule4

yystate33:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate34:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate35
	}

yystate35:
	c = y.getc()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate35
	}

yystate36:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'a':
		goto yystate37
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate33
	}

yystate37:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'l':
		goto yystate38
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate33
	}

yystate38:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 's':
		goto yystate39
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate33
	}

yystate39:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'e':
		goto yystate40
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate33
	}

yystate40:
	c = y.getc()
	switch {
	default:
		goto yyrule13
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate41:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'n':
		goto yystate42
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate33
	}

yystate42:
	c = y.getc()
	switch {
	default:
		goto yyrule14
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate43:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'o':
		goto yystate44
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate33
	}

yystate44:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 't':
		goto yystate45
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate33
	}

yystate45:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'i':
		goto yystate46
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate33
	}

yystate46:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'n':
		goto yystate47
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate33
	}

yystate47:
	c = y.getc()
	switch {
	default:
		goto yyrule15
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate48:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'r':
		goto yystate49
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate33
	}

yystate49:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'u':
		goto yystate50
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate33
	}

yystate50:
	c = y.getc()
	switch {
	default:
		goto yyrule19
	case c == '.':
		goto yystate13
	case c == 'e':
		goto yystate51
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate33
	}

yystate51:
	c = y.getc()
	switch {
	default:
		goto yyrule12
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate52:
	c = y.getc()
	switch {
	default:
		goto yyrule11
	case c == '.':
		goto yystate13
	case c == '|':
		goto yystate53
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate53:
	c = y.getc()
	switch {
	default:
		goto yyrule9
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate35
	}

yyrule1: // [ \t\r]+

	goto yystate0
yyrule2: // "=="
	{
		return EQU
	}
yyrule3: // "!="
	{
		return NEQ
	}
yyrule4: // ">="
	{
		return GEQ
	}
yyrule5: // ">"
	{
		return GTR
	}
yyrule6: // "<="
	{
		return LEQ
	}
yyrule7: // "<"
	{
		return LSS
	}
yyrule8: // "&&"
	{
		return AND
	}
yyrule9: // "||"
	{
		return OR
	}
yyrule10: // "&"
	{
		return BITAND
	}
yyrule11: // "|"
	{
		return BITOR
	}
yyrule12: // true
	{
		return TRUE
	}
yyrule13: // false
	{
		return FALSE
	}
yyrule14: // in
	{
		lval.value = "in"
		return IN
		goto yystate0
	}
yyrule15: // notin
	{
		lval.value = "notin"
		return NOTIN
		goto yystate0
	}
yyrule16: // {string}
	{
		lval.value = string(y.buf[1 : len(y.buf)-1])
		return STRING
		goto yystate0
	}
yyrule17: // {int}
	{
		lval.value = string(y.buf)
		return INT
		goto yystate0
	}
yyrule18: // {float}
	{
		lval.value = string(y.buf)
		return FLOAT
		goto yystate0
	}
yyrule19: // {identifier}
	{
		lval.value = string(y.buf)
		return IDENT
		goto yystate0
	}
yyrule20: // {arrays}
	{
		lval.value = string(y.buf)
		return ARRAY
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	if len(y.buf) != 0 {
		y.Error("")
	}
	return int(c)
}

func newLexer(src string) (y *yylexer) {
	y = &yylexer{
		src:     src,
		size:    len(src),
		pos:     1,
		current: src[0],
	}
	return
}
