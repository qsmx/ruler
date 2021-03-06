// Code generated by goyacc - DO NOT EDIT.

package ruler

import __yyfmt__ "fmt"

type yySymType struct {
	yys   int
	value string
	ruler *node
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57372
	yyEofCode = 57344
	AND       = 57360
	ARGS      = 57367
	ARRAY     = 57353
	BITAND    = 57364
	BITOR     = 57365
	EQU       = 57354
	FALSE     = 57350
	FILTER    = 57368
	FLOAT     = 57348
	FUNC      = 57366
	GEQ       = 57359
	GTR       = 57358
	IDENT     = 57352
	IN        = 57362
	INT       = 57346
	INTHEX    = 57347
	LEQ       = 57357
	LISTA     = 57369
	LISTO     = 57370
	LSS       = 57356
	NEG       = 57371
	NEQ       = 57355
	NOTIN     = 57363
	OR        = 57361
	STRING    = 57351
	TRUE      = 57349
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -56
)

var (
	yyPrec = map[int]int{
		OR:     0,
		AND:    1,
		BITOR:  2,
		BITAND: 3,
		EQU:    4,
		NEQ:    4,
		LSS:    4,
		LEQ:    4,
		GTR:    4,
		GEQ:    4,
		'-':    5,
		'+':    5,
		'*':    6,
		'/':    6,
		'%':    6,
		NEG:    7,
		'^':    8,
	}

	yyXLAT = map[int]int{
		45:    0,  // '-' (66x)
		41:    1,  // ')' (64x)
		57360: 2,  // AND (58x)
		57361: 3,  // OR (58x)
		57344: 4,  // $end (56x)
		93:    5,  // ']' (44x)
		57362: 6,  // IN (37x)
		57363: 7,  // NOTIN (37x)
		37:    8,  // '%' (36x)
		40:    9,  // '(' (36x)
		42:    10, // '*' (36x)
		43:    11, // '+' (36x)
		47:    12, // '/' (36x)
		57364: 13, // BITAND (36x)
		57365: 14, // BITOR (36x)
		44:    15, // ',' (34x)
		57354: 16, // EQU (34x)
		57355: 17, // NEQ (34x)
		57353: 18, // ARRAY (30x)
		57375: 19, // CalcCell (30x)
		57378: 20, // ExprCalc (30x)
		57380: 21, // ExprFunc (30x)
		57348: 22, // FLOAT (30x)
		57381: 23, // Float (30x)
		57359: 24, // GEQ (30x)
		57358: 25, // GTR (30x)
		57352: 26, // IDENT (30x)
		57384: 27, // Ident (30x)
		57383: 28, // IDENTS (30x)
		57346: 29, // INT (30x)
		57386: 30, // Int (30x)
		57357: 31, // LEQ (30x)
		57356: 32, // LSS (30x)
		57351: 33, // STRING (30x)
		57387: 34, // String (30x)
		57373: 35, // Bool (17x)
		57374: 36, // BoolCell (17x)
		57350: 37, // FALSE (17x)
		57349: 38, // TRUE (17x)
		57377: 39, // ExprBool (11x)
		91:    40, // '[' (10x)
		57379: 41, // ExprFilter (4x)
		57382: 42, // FuncArgs (3x)
		57376: 43, // Expr (1x)
		57385: 44, // Input (1x)
		57372: 45, // $default (0x)
		94:    46, // '^' (0x)
		57367: 47, // ARGS (0x)
		57345: 48, // error (0x)
		57368: 49, // FILTER (0x)
		57366: 50, // FUNC (0x)
		57347: 51, // INTHEX (0x)
		57369: 52, // LISTA (0x)
		57370: 53, // LISTO (0x)
		57371: 54, // NEG (0x)
	}

	yySymNames = []string{
		"'-'",
		"')'",
		"AND",
		"OR",
		"$end",
		"']'",
		"IN",
		"NOTIN",
		"'%'",
		"'('",
		"'*'",
		"'+'",
		"'/'",
		"BITAND",
		"BITOR",
		"','",
		"EQU",
		"NEQ",
		"ARRAY",
		"CalcCell",
		"ExprCalc",
		"ExprFunc",
		"FLOAT",
		"Float",
		"GEQ",
		"GTR",
		"IDENT",
		"Ident",
		"IDENTS",
		"INT",
		"Int",
		"LEQ",
		"LSS",
		"STRING",
		"String",
		"Bool",
		"BoolCell",
		"FALSE",
		"TRUE",
		"ExprBool",
		"'['",
		"ExprFilter",
		"FuncArgs",
		"Expr",
		"Input",
		"$default",
		"'^'",
		"ARGS",
		"error",
		"FILTER",
		"FUNC",
		"INTHEX",
		"LISTA",
		"LISTO",
		"NEG",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {44, 0},
		2:  {44, 1},
		3:  {43, 1},
		4:  {43, 1},
		5:  {39, 1},
		6:  {39, 1},
		7:  {39, 5},
		8:  {39, 5},
		9:  {39, 3},
		10: {39, 3},
		11: {39, 3},
		12: {39, 3},
		13: {39, 3},
		14: {39, 3},
		15: {39, 3},
		16: {39, 3},
		17: {39, 3},
		18: {41, 3},
		19: {41, 4},
		20: {41, 3},
		21: {41, 3},
		22: {41, 3},
		23: {41, 3},
		24: {41, 3},
		25: {21, 3},
		26: {21, 4},
		27: {42, 1},
		28: {42, 3},
		29: {20, 1},
		30: {20, 3},
		31: {20, 3},
		32: {20, 3},
		33: {20, 3},
		34: {20, 3},
		35: {20, 3},
		36: {20, 3},
		37: {20, 2},
		38: {20, 3},
		39: {36, 1},
		40: {36, 1},
		41: {19, 1},
		42: {19, 1},
		43: {19, 1},
		44: {19, 1},
		45: {19, 1},
		46: {30, 1},
		47: {23, 1},
		48: {34, 1},
		49: {35, 1},
		50: {35, 1},
		51: {27, 1},
		52: {27, 1},
		53: {28, 1},
		54: {28, 1},
		55: {28, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [103][]uint16{
		// 0
		{70, 4: 55, 6: 81, 82, 9: 66, 18: 79, 69, 65, 62, 75, 72, 26: 80, 63, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 59, 67, 60, 43: 58, 57},
		{4: 56},
		{4: 54},
		{2: 146, 147, 53},
		{2: 141, 142, 52},
		// 5
		{1: 51, 51, 51, 51, 51, 16: 17, 17},
		{11, 50, 50, 50, 50, 50, 8: 11, 10: 11, 11, 11, 11, 11, 16: 11, 11, 24: 11, 11, 31: 11, 11},
		{12, 12, 6: 131, 132, 12, 10: 12, 12, 12, 12, 12, 16: 12, 12, 24: 12, 12, 31: 12, 12, 40: 156},
		{16: 152, 153},
		{89, 8: 92, 10: 90, 88, 91, 93, 94, 16: 16, 16, 24: 120, 119, 31: 118, 117},
		// 10
		{70, 6: 81, 82, 9: 66, 18: 79, 69, 116, 62, 75, 72, 26: 80, 63, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 139, 67, 140},
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 65, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 113},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 103, 5, 5, 5, 5, 5, 5, 5, 5, 24: 5, 5, 31: 5, 5, 40: 5},
		{27, 27, 27, 27, 27, 27, 8: 27, 10: 27, 27, 27, 27, 27, 27, 27, 27, 24: 27, 27, 31: 27, 27},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 83, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		// 15
		{15, 15, 15, 15, 15, 15, 8: 15, 10: 15, 15, 15, 15, 15, 15, 15, 15, 24: 15, 15, 31: 15, 15},
		{14, 14, 14, 14, 14, 14, 8: 14, 10: 14, 14, 14, 14, 14, 14, 14, 14, 24: 14, 14, 31: 14, 14},
		{13, 13, 13, 13, 13, 13, 8: 13, 10: 13, 13, 13, 13, 13, 13, 13, 13, 24: 13, 13, 31: 13, 13},
		{10, 10, 10, 10, 10, 10, 8: 10, 10: 10, 10, 10, 10, 10, 10, 10, 10, 24: 10, 10, 31: 10, 10},
		{9, 9, 9, 9, 9, 9, 8: 9, 10: 9, 9, 9, 9, 9, 9, 9, 9, 24: 9, 9, 31: 9, 9},
		// 20
		{8, 8, 8, 8, 8, 8, 8: 8, 10: 8, 8, 8, 8, 8, 8, 8, 8, 24: 8, 8, 31: 8, 8},
		{1: 7, 7, 7, 7, 7, 15: 7, 7, 7},
		{1: 6, 6, 6, 6, 6, 15: 6, 6, 6},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 10: 4, 4, 4, 4, 4, 4, 4, 4, 24: 4, 4, 31: 4, 4, 40: 4},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 24: 3, 3, 31: 3, 3, 40: 3},
		// 25
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 24: 2, 2, 31: 2, 2, 40: 2},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 24: 1, 1, 31: 1, 1, 40: 1},
		{19, 19, 19, 19, 19, 19, 8: 19, 10: 19, 19, 19, 19, 19, 19, 19, 19, 24: 19, 19, 31: 19, 19},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 87, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{12, 12, 12, 12, 12, 12, 8: 12, 10: 12, 12, 12, 12, 12, 12, 12, 12, 24: 12, 12, 31: 12, 12},
		// 30
		{11, 11, 11, 11, 11, 11, 8: 11, 10: 11, 11, 11, 11, 11, 11, 11, 11, 24: 11, 11, 31: 11, 11},
		{89, 95, 8: 92, 10: 90, 88, 91, 93, 94},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 102, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 101, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 100, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		// 35
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 99, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 98, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 97, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 96, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{18, 18, 18, 18, 18, 18, 8: 18, 10: 18, 18, 18, 18, 18, 18, 18, 18, 24: 18, 18, 31: 18, 18},
		// 40
		{89, 20, 20, 20, 20, 20, 8: 92, 10: 90, 88, 91, 93, 20, 20, 20, 20, 24: 20, 20, 31: 20, 20},
		{89, 21, 21, 21, 21, 21, 8: 92, 10: 90, 88, 91, 21, 21, 21, 21, 21, 24: 21, 21, 31: 21, 21},
		{22, 22, 22, 22, 22, 22, 8: 22, 10: 22, 22, 22, 22, 22, 22, 22, 22, 24: 22, 22, 31: 22, 22},
		{23, 23, 23, 23, 23, 23, 8: 23, 10: 23, 23, 23, 23, 23, 23, 23, 23, 24: 23, 23, 31: 23, 23},
		{24, 24, 24, 24, 24, 24, 8: 24, 10: 24, 24, 24, 24, 24, 24, 24, 24, 24: 24, 24, 31: 24, 24},
		// 45
		{25, 25, 25, 25, 25, 25, 8: 92, 10: 90, 25, 91, 25, 25, 25, 25, 25, 24: 25, 25, 31: 25, 25},
		{26, 26, 26, 26, 26, 26, 8: 92, 10: 90, 26, 91, 26, 26, 26, 26, 26, 24: 26, 26, 31: 26, 26},
		{70, 104, 6: 81, 82, 9: 84, 18: 79, 69, 107, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73, 108, 106, 78, 77, 42: 105},
		{31, 31, 31, 31, 31, 31, 8: 31, 10: 31, 31, 31, 31, 31, 31, 31, 31, 24: 31, 31, 31: 31, 31},
		{1: 109, 15: 110},
		// 50
		{1: 29, 15: 29},
		{89, 16, 16, 16, 16, 16, 8: 92, 10: 90, 88, 91, 93, 94, 16},
		{1: 17, 17, 17, 17, 17, 15: 17},
		{30, 30, 30, 30, 30, 30, 8: 30, 10: 30, 30, 30, 30, 30, 30, 30, 30, 24: 30, 30, 31: 30, 30},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 107, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73, 108, 111, 78, 77},
		// 55
		{1: 28, 15: 28},
		{12, 12, 6: 131, 132, 12, 10: 12, 12, 12, 12, 12, 16: 12, 12, 24: 12, 12, 31: 12, 12},
		{2: 125, 126, 5: 130},
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 116, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 115},
		{1: 127, 125, 126},
		// 60
		{89, 95, 8: 92, 10: 90, 88, 91, 93, 94, 16: 16, 16, 24: 120, 119, 31: 118, 117},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 124, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 123, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 122, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 121, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73},
		// 65
		{89, 40, 40, 40, 40, 40, 8: 92, 10: 90, 88, 91, 93, 94},
		{89, 41, 41, 41, 41, 41, 8: 92, 10: 90, 88, 91, 93, 94},
		{89, 42, 42, 42, 42, 42, 8: 92, 10: 90, 88, 91, 93, 94},
		{89, 43, 43, 43, 43, 43, 8: 92, 10: 90, 88, 91, 93, 94},
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 65, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 129},
		// 70
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 65, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 128},
		{1: 39, 39, 39, 39, 39},
		{1: 46, 125, 46, 46, 46},
		{1: 47, 47, 47, 47, 47},
		{1: 38, 38, 38, 38},
		// 75
		{9: 136},
		{9: 133},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 107, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73, 108, 106, 78, 77, 42: 134},
		{1: 135, 15: 110},
		{1: 48, 48, 48, 48, 48},
		// 80
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 107, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73, 108, 106, 78, 77, 42: 137},
		{1: 138, 15: 110},
		{1: 49, 49, 49, 49, 49},
		{1: 127, 146, 147},
		{1: 143, 141, 142},
		// 85
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 65, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 145},
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 65, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 144},
		{1: 32, 32, 32, 32},
		{1: 33, 125, 33, 33},
		{1: 34, 34, 34, 34},
		// 90
		{70, 6: 81, 82, 9: 66, 18: 79, 69, 65, 62, 75, 72, 26: 80, 63, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 150, 67, 151},
		{70, 6: 81, 82, 9: 66, 18: 79, 69, 65, 62, 75, 72, 26: 80, 63, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 148, 67, 149},
		{1: 46, 146, 46, 46},
		{1: 35, 141, 35, 35},
		{1: 47, 47, 47, 47},
		// 95
		{1: 36, 36, 36, 36},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 107, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73, 108, 155, 78, 77},
		{70, 6: 81, 82, 9: 84, 18: 79, 69, 107, 86, 75, 72, 26: 80, 85, 68, 74, 71, 33: 76, 73, 108, 154, 78, 77},
		{1: 44, 44, 44, 44, 44},
		{1: 45, 45, 45, 45, 45},
		// 100
		{70, 6: 81, 82, 9: 114, 18: 79, 69, 65, 62, 75, 72, 26: 80, 112, 68, 74, 71, 33: 76, 73, 61, 64, 78, 77, 157},
		{2: 125, 126, 5: 158},
		{1: 37, 37, 37, 37},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer, outNode **node) int {
	const yyError = 48

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 2:
		{
			*outNode = yyS[yypt-0].ruler
		}
	case 7:
		{
			yyVAL.ruler = nodeFromNode2(IN, "i", yyS[yypt-4].ruler, yyS[yypt-1].ruler)
		}
	case 8:
		{
			yyVAL.ruler = nodeFromNode2(IN, "n", yyS[yypt-4].ruler, yyS[yypt-1].ruler)
		}
	case 9:
		{
			yyVAL.ruler = nodeFromNode(AND, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 10:
		{
			yyVAL.ruler = nodeFromNode(OR, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 11:
		{
			yyVAL.ruler = nodeFromNode(EQU, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 12:
		{
			yyVAL.ruler = nodeFromNode(NEQ, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 13:
		{
			yyVAL.ruler = nodeFromNode(LSS, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 14:
		{
			yyVAL.ruler = nodeFromNode(LEQ, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 15:
		{
			yyVAL.ruler = nodeFromNode(GTR, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 16:
		{
			yyVAL.ruler = nodeFromNode(GEQ, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 17:
		{
			yyVAL.ruler = yyS[yypt-1].ruler
		}
	case 18:
		{
			yyVAL.ruler = nodeFromNode(FILTER, nodeFromIdent(IDENT, DEFAULT_SLICE_LABEL), yyS[yypt-1].ruler)
		}
	case 19:
		{
			yyVAL.ruler = nodeFromNode(FILTER, yyS[yypt-3].ruler, yyS[yypt-1].ruler)
		}
	case 20:
		{
			yyVAL.ruler = nodeFromNode(LISTA, yyS[yypt-2].ruler, yyS[yypt-2].ruler)
		}
	case 21:
		{
			yyVAL.ruler = nodeFromNode(LISTO, yyS[yypt-2].ruler, yyS[yypt-2].ruler)
		}
	case 22:
		{
			yyVAL.ruler = nodeFromNode(LISTA, yyS[yypt-0].ruler, yyS[yypt-2].ruler)
		}
	case 23:
		{
			yyVAL.ruler = nodeFromNode(LISTO, yyS[yypt-0].ruler, yyS[yypt-2].ruler)
		}
	case 24:
		{
			yyVAL.ruler = yyS[yypt-1].ruler
		}
	case 25:
		{
			yyVAL.ruler = nodeFromFunc(FUNC, yyS[yypt-2].value)
		}
	case 26:
		{
			yyVAL.ruler = nodeFromFunc(FUNC, yyS[yypt-3].value, yyS[yypt-1].ruler)
		}
	case 27:
		{
			yyVAL.ruler = nodeFromNode(ARGS, yyS[yypt-0].ruler)
		}
	case 28:
		{
			yyVAL.ruler.AppendNode(yyS[yypt-0].ruler)
		}
	case 30:
		{
			yyVAL.ruler = nodeFromNode('+', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 31:
		{
			yyVAL.ruler = nodeFromNode('-', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 32:
		{
			yyVAL.ruler = nodeFromNode('*', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 33:
		{
			yyVAL.ruler = nodeFromNode('/', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 34:
		{
			yyVAL.ruler = nodeFromNode('%', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 35:
		{
			yyVAL.ruler = nodeFromNode(BITAND, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 36:
		{
			yyVAL.ruler = nodeFromNode(BITOR, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 37:
		{
			yyVAL.ruler = nodeFromNode(NEG, yyS[yypt-0].ruler)
		}
	case 38:
		{
			yyVAL.ruler = yyS[yypt-1].ruler
		}
	case 46:
		{
			yyVAL.ruler = nodeFromInt("", ConvertInt(yyS[yypt-0].value))
		}
	case 47:
		{
			yyVAL.ruler = nodeFromFloat("", ConvertFloat(yyS[yypt-0].value))
		}
	case 48:
		{
			yyVAL.ruler = nodeFromString("", yyS[yypt-0].value)
		}
	case 49:
		{
			yyVAL.ruler = nodeFromBool("", true)
		}
	case 50:
		{
			yyVAL.ruler = nodeFromBool("", false)
		}
	case 51:
		{
			yyVAL.ruler = nodeFromIdent(IDENT, yyS[yypt-0].value)
		}
	case 52:
		{
			yyVAL.ruler = nodeFromIdent(ARRAY, yyS[yypt-0].value)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
