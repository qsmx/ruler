// Code generated by goyacc - DO NOT EDIT.

package ruler

import __yyfmt__ "fmt"

type yySymType struct {
	yys   int
	value string

	ruler *rulerNode
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57369
	yyEofCode = 57344
	AND       = 57359
	ARGS      = 57364
	ARRAYS    = 57352
	EQU       = 57353
	FALSE     = 57349
	FILTER    = 57365
	FLOAT     = 57347
	FUNC      = 57363
	GEQ       = 57358
	GTR       = 57357
	IDENT     = 57351
	IN        = 57361
	INT       = 57346
	LEQ       = 57356
	LISTA     = 57366
	LISTO     = 57367
	LSS       = 57355
	NEG       = 57368
	NEQ       = 57354
	NOTIN     = 57362
	OR        = 57360
	STRING    = 57350
	TRUE      = 57348
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -55
)

var (
	yyPrec = map[int]int{
		OR:  0,
		AND: 1,
		EQU: 2,
		NEQ: 2,
		LSS: 3,
		LEQ: 3,
		GTR: 3,
		GEQ: 3,
		'-': 4,
		'+': 4,
		'*': 5,
		'/': 5,
		'%': 5,
		NEG: 6,
		'^': 7,
	}

	yyXLAT = map[int]int{
		45:    0,  // '-' (59x)
		41:    1,  // ')' (54x)
		57344: 2,  // $end (53x)
		57359: 3,  // AND (46x)
		57360: 4,  // OR (46x)
		93:    5,  // ']' (38x)
		37:    6,  // '%' (34x)
		42:    7,  // '*' (34x)
		43:    8,  // '+' (34x)
		47:    9,  // '/' (34x)
		57353: 10, // EQU (32x)
		57354: 11, // NEQ (32x)
		44:    12, // ',' (29x)
		57358: 13, // GEQ (28x)
		57357: 14, // GTR (28x)
		57356: 15, // LEQ (28x)
		57355: 16, // LSS (28x)
		40:    17, // '(' (26x)
		57371: 18, // Arrays (25x)
		57352: 19, // ARRAYS (25x)
		57374: 20, // CalcCell (25x)
		57377: 21, // ExprCalc (25x)
		57378: 22, // ExprExt (25x)
		57380: 23, // ExprFunc (25x)
		57347: 24, // FLOAT (25x)
		57384: 25, // Float (25x)
		57351: 26, // IDENT (25x)
		57385: 27, // Ident (25x)
		57387: 28, // Int (25x)
		57346: 29, // INT (25x)
		57350: 30, // STRING (25x)
		57388: 31, // String (25x)
		57372: 32, // Bool (14x)
		57373: 33, // BoolCell (14x)
		57349: 34, // FALSE (14x)
		57348: 35, // TRUE (14x)
		57376: 36, // ExprBool (10x)
		91:    37, // '[' (5x)
		57382: 38, // FilterCell (4x)
		57383: 39, // FilterName (4x)
		57379: 40, // ExprFilter (2x)
		57370: 41, // Args (1x)
		57375: 42, // Expr (1x)
		57381: 43, // ExprList (1x)
		57386: 44, // Input (1x)
		57369: 45, // $default (0x)
		94:    46, // '^' (0x)
		57364: 47, // ARGS (0x)
		57345: 48, // error (0x)
		57365: 49, // FILTER (0x)
		57363: 50, // FUNC (0x)
		57361: 51, // IN (0x)
		57366: 52, // LISTA (0x)
		57367: 53, // LISTO (0x)
		57368: 54, // NEG (0x)
		57362: 55, // NOTIN (0x)
	}

	yySymNames = []string{
		"'-'",
		"')'",
		"$end",
		"AND",
		"OR",
		"']'",
		"'%'",
		"'*'",
		"'+'",
		"'/'",
		"EQU",
		"NEQ",
		"','",
		"GEQ",
		"GTR",
		"LEQ",
		"LSS",
		"'('",
		"Arrays",
		"ARRAYS",
		"CalcCell",
		"ExprCalc",
		"ExprExt",
		"ExprFunc",
		"FLOAT",
		"Float",
		"IDENT",
		"Ident",
		"Int",
		"INT",
		"STRING",
		"String",
		"Bool",
		"BoolCell",
		"FALSE",
		"TRUE",
		"ExprBool",
		"'['",
		"FilterCell",
		"FilterName",
		"ExprFilter",
		"Args",
		"Expr",
		"ExprList",
		"Input",
		"$default",
		"'^'",
		"ARGS",
		"error",
		"FILTER",
		"FUNC",
		"IN",
		"LISTA",
		"LISTO",
		"NEG",
		"NOTIN",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {44, 0},
		2:  {44, 1},
		3:  {42, 1},
		4:  {42, 1},
		5:  {42, 1},
		6:  {43, 1},
		7:  {43, 1},
		8:  {36, 1},
		9:  {36, 1},
		10: {36, 3},
		11: {36, 3},
		12: {36, 3},
		13: {36, 3},
		14: {36, 3},
		15: {36, 3},
		16: {36, 3},
		17: {36, 3},
		18: {36, 3},
		19: {33, 1},
		20: {33, 1},
		21: {21, 1},
		22: {21, 3},
		23: {21, 3},
		24: {21, 3},
		25: {21, 3},
		26: {21, 3},
		27: {21, 2},
		28: {21, 3},
		29: {20, 1},
		30: {20, 1},
		31: {22, 1},
		32: {22, 1},
		33: {22, 1},
		34: {22, 1},
		35: {22, 1},
		36: {23, 3},
		37: {23, 4},
		38: {41, 1},
		39: {41, 3},
		40: {40, 3},
		41: {40, 3},
		42: {40, 3},
		43: {40, 3},
		44: {40, 3},
		45: {38, 4},
		46: {39, 1},
		47: {39, 1},
		48: {27, 1},
		49: {18, 1},
		50: {28, 1},
		51: {25, 1},
		52: {31, 1},
		53: {32, 1},
		54: {32, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [90][]uint16{
		// 0
		{69, 2: 54, 17: 67, 71, 77, 68, 66, 59, 64, 79, 73, 75, 70, 72, 78, 80, 74, 63, 65, 82, 81, 58, 38: 61, 76, 62, 42: 57, 60, 56},
		{2: 55},
		{2: 53},
		{2: 52, 137, 138},
		{26, 2: 51, 6: 26, 26, 26, 26, 26, 26, 13: 26, 26, 26, 26},
		// 5
		{2: 50},
		{2: 49, 133, 134},
		{2: 48},
		{1: 47, 47, 47, 47, 47, 10: 36, 36},
		{25, 46, 46, 46, 46, 46, 25, 25, 25, 25, 25, 25, 13: 25, 25, 25, 25},
		// 10
		{10: 142, 141},
		{101, 6: 104, 102, 100, 103, 35, 35, 13: 99, 98, 97, 96},
		{69, 17: 67, 71, 77, 68, 89, 86, 64, 79, 73, 75, 70, 72, 78, 80, 74, 63, 65, 82, 81, 129, 38: 130, 76, 131},
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{69, 17: 107, 88, 77, 68, 128, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		// 15
		{24, 24, 24, 6: 24, 24, 24, 24, 24, 24, 13: 24, 24, 24, 24, 37: 9},
		{23, 23, 23, 6: 23, 23, 23, 23, 23, 23, 13: 23, 23, 23, 23, 37: 8},
		{22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 20
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 119, 37: 7},
		{37: 83},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 37: 6},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		// 25
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{1: 2, 2, 2, 2, 2, 10: 2, 2, 2},
		{1: 1, 1, 1, 1, 1, 10: 1, 1, 1},
		{69, 17: 85, 88, 77, 68, 66, 86, 64, 79, 73, 75, 87, 72, 78, 80, 74, 63, 65, 82, 81, 84},
		{3: 91, 92, 118},
		// 30
		{69, 17: 85, 88, 77, 68, 89, 86, 64, 79, 73, 75, 87, 72, 78, 80, 74, 63, 65, 82, 81, 90},
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{101, 105, 6: 104, 102, 100, 103, 35, 35, 13: 99, 98, 97, 96},
		// 35
		{1: 93, 3: 91, 92},
		{69, 17: 85, 88, 77, 68, 66, 86, 64, 79, 73, 75, 87, 72, 78, 80, 74, 63, 65, 82, 81, 95},
		{69, 17: 85, 88, 77, 68, 66, 86, 64, 79, 73, 75, 87, 72, 78, 80, 74, 63, 65, 82, 81, 94},
		{1: 37, 37, 37, 37, 37},
		{1: 38, 38, 91, 38, 38},
		// 40
		{1: 39, 39, 39, 39, 39},
		{69, 17: 107, 88, 77, 68, 117, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 116, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 115, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 114, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		// 45
		{69, 17: 107, 88, 77, 68, 113, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 112, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 111, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 110, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{69, 17: 107, 88, 77, 68, 106, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		// 50
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		{69, 17: 107, 88, 77, 68, 109, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74},
		{25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		{101, 105, 6: 104, 102, 100, 103},
		// 55
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		{32, 32, 32, 32, 32, 32, 104, 102, 32, 103, 32, 32, 32, 32, 32, 32, 32},
		{33, 33, 33, 33, 33, 33, 104, 102, 33, 103, 33, 33, 33, 33, 33, 33, 33},
		{101, 40, 40, 40, 40, 40, 104, 102, 100, 103},
		// 60
		{101, 41, 41, 41, 41, 41, 104, 102, 100, 103},
		{101, 42, 42, 42, 42, 42, 104, 102, 100, 103},
		{101, 43, 43, 43, 43, 43, 104, 102, 100, 103},
		{1: 10, 10, 10, 10},
		{69, 122, 17: 107, 88, 77, 68, 121, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74, 120, 124, 82, 81, 41: 123},
		// 65
		{1: 36, 36, 36, 36, 36, 12: 36},
		{101, 35, 35, 35, 35, 35, 104, 102, 100, 103, 12: 35},
		{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{1: 125, 12: 126},
		{1: 17, 12: 17},
		// 70
		{18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{69, 17: 107, 88, 77, 68, 121, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74, 120, 127, 82, 81},
		{1: 16, 12: 16},
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{1: 93, 3: 137, 138},
		// 75
		{3: 133, 134},
		{1: 132},
		{1: 11, 11},
		{69, 17: 85, 88, 77, 68, 66, 86, 64, 79, 73, 75, 87, 72, 78, 80, 74, 63, 65, 82, 81, 136},
		{69, 17: 85, 88, 77, 68, 66, 86, 64, 79, 73, 75, 87, 72, 78, 80, 74, 63, 65, 82, 81, 135},
		// 80
		{1: 14, 14, 91, 92},
		{1: 15, 15, 91, 92},
		{69, 17: 85, 71, 77, 68, 66, 86, 64, 79, 73, 75, 70, 72, 78, 80, 74, 63, 65, 82, 81, 95, 38: 140, 76},
		{69, 17: 85, 71, 77, 68, 66, 86, 64, 79, 73, 75, 70, 72, 78, 80, 74, 63, 65, 82, 81, 94, 38: 139, 76},
		{1: 12, 12},
		// 85
		{1: 13, 13},
		{69, 17: 107, 88, 77, 68, 121, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74, 120, 144, 82, 81},
		{69, 17: 107, 88, 77, 68, 121, 86, 108, 79, 73, 75, 87, 72, 78, 80, 74, 120, 143, 82, 81},
		{1: 44, 44, 44, 44, 44},
		{1: 45, 45, 45, 45, 45},
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

func yyParse(yylex yyLexer, outNode **rulerNode) int {
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
			if outNode != nil {
				*outNode = yyS[yypt-0].ruler
			}
		}
	case 10:
		{
			yyVAL.ruler = nodeFromNode(NEQ, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 11:
		{
			yyVAL.ruler = nodeFromNode(EQU, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 12:
		{
			yyVAL.ruler = nodeFromNode(LSS, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 13:
		{
			yyVAL.ruler = nodeFromNode(LEQ, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 14:
		{
			yyVAL.ruler = nodeFromNode(GTR, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 15:
		{
			yyVAL.ruler = nodeFromNode(GEQ, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 16:
		{
			yyVAL.ruler = nodeFromNode(AND, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 17:
		{
			yyVAL.ruler = nodeFromNode(OR, yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 18:
		{
			yyVAL.ruler = yyS[yypt-1].ruler
		}
	case 22:
		{
			yyVAL.ruler = nodeFromNode('+', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 23:
		{
			yyVAL.ruler = nodeFromNode('-', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 24:
		{
			yyVAL.ruler = nodeFromNode('*', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 25:
		{
			yyVAL.ruler = nodeFromNode('/', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 26:
		{
			yyVAL.ruler = nodeFromNode('%', yyS[yypt-2].ruler, yyS[yypt-0].ruler)
		}
	case 27:
		{
			yyVAL.ruler = nodeFromNode(NEG, yyS[yypt-0].ruler)
		}
	case 28:
		{
			yyVAL.ruler = yyS[yypt-1].ruler
		}
	case 36:
		{
			yyVAL.ruler = nodeFromString(FUNC, yyS[yypt-2].value)
		}
	case 37:
		{
			yyVAL.ruler = nodeFromString(FUNC, yyS[yypt-3].value)
			yyVAL.ruler.ReplaceNode(yyS[yypt-1].ruler)
		}
	case 38:
		{
			yyVAL.ruler = nodeFromNode(ARGS, yyS[yypt-0].ruler)
		}
	case 39:
		{
			yyVAL.ruler = yyVAL.ruler.AppendNode(yyS[yypt-0].ruler)
		}
	case 40:
		{
			yyVAL.ruler = nodeFromNode(LISTA, yyS[yypt-0].ruler)
			yyVAL.ruler.AppendNode(yyS[yypt-2].ruler)
		}
	case 41:
		{
			yyVAL.ruler = nodeFromNode(LISTO, yyS[yypt-0].ruler)
			yyVAL.ruler.AppendNode(yyS[yypt-2].ruler)
		}
	case 42:
		{
			yyVAL.ruler = nodeFromNode(LISTA, yyS[yypt-2].ruler)
			yyVAL.ruler.AppendNode(yyS[yypt-0].ruler)
		}
	case 43:
		{
			yyVAL.ruler = nodeFromNode(LISTO, yyS[yypt-2].ruler)
			yyVAL.ruler.AppendNode(yyS[yypt-0].ruler)
		}
	case 44:
		{
			yyVAL.ruler = yyS[yypt-1].ruler
		}
	case 45:
		{
			yyVAL.ruler = nodeFromNode(FILTER, yyS[yypt-3].ruler)
			yyVAL.ruler.AppendNode(yyS[yypt-1].ruler)
		}
	case 48:
		{
			yyVAL.ruler = nodeFromIdent(IDENT, yyS[yypt-0].value)
		}
	case 49:
		{
			yyVAL.ruler = nodeFromIdent(ARRAYS, yyS[yypt-0].value)
		}
	case 50:
		{
			yyVAL.ruler = nodeFromInt(ConvertInt(yyS[yypt-0].value))
		}
	case 51:
		{
			yyVAL.ruler = nodeFromFloat(ConvertFloat(yyS[yypt-0].value))
		}
	case 52:
		{
			yyVAL.ruler = nodeFromString(STRING, yyS[yypt-0].value)
		}
	case 53:
		{
			yyVAL.ruler = nodeFromBool(true)
		}
	case 54:
		{
			yyVAL.ruler = nodeFromBool(false)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}