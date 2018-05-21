func validateParse(yylex validateLexer, dp *DataPackage, res *bool) int {
	const yyError = 22

	yyEx, _ := yylex.(validateLexerEx)
	var yyn int
	var yylval validateSymType
	var yyVAL validateSymType
	yyS := make([]validateSymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if validateDebug >= 2 {
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
		nyys := make([]validateSymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = validatelex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = validateXLAT[yychar]; !ok {
			yyxchar = len(validateSymNames) // > tab width
		}
	}
	if validateDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := validateParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += validateTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if validateDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if validateDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if validateDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", validateSymName(yychar), yystate)
			}
			msg, ok := validateXErrors[validateXError{yystate, yyxchar}]
			if !ok {
				msg, ok = validateXErrors[validateXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = validateXErrors[validateXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = validateXErrors[validateXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := validateTokenLiteralStrings[yychar]
				if ls == "" {
					ls = validateSymName(yychar)
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
				row := validateParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + validateTabOfs
					if yyn > 0 { // hit
						if validateDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if validateDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if validateDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if validateDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", validateSymName(yychar))
			}
			if yychar == validateEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := validateReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]validateSymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(validateParseTab[yyS[yyp].yys][x]) + validateTabOfs
	/* reduction by production r */
	if validateDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, validateSymNames[x], yystate)
	}

	switch r {
	case 2:
		{
			fmt.Println(yyS[yypt-0].value_string)
		}
	case 3:
		{
			fmt.Println("INT:", yyS[yypt-0].value_string)
		}
	case 4:
		{
			fmt.Println("FLOAT:", yyS[yypt-0].value_string)
		}
	case 5:
		{
			fmt.Println("STRING:", yyS[yypt-0].value_string)
		}
	case 6:
		{
			fmt.Println("IDENT:", yyS[yypt-0].value_string)
		}
	case 7:
		{
			fmt.Println("ARRAYS:", yyS[yypt-0].value_string)
		}
	case 8:
		{
			fmt.Println("FILTER:", yyS[yypt-0].value_string)
		}
	case 9:
		{
			fmt.Printf("FilterData: %+v\n", yyS[yypt-3].value)
			for _, v := range []int{1, 2, 3} {
				fmt.Println("[", yyS[yypt-1].value_bool, "]", v)
			}
			yyVAL.value_string = "1"
		}
	case 10:
		{
			fmt.Println("sub:", (yyS[yypt-2].value_string + "!=" + yyS[yypt-0].value_string))
			yyVAL.value_bool = dp.GetInt(yyS[yypt-2].value_string) != ConvertInt(yyS[yypt-0].value_string)
		}
	case 11:
		{
			yyVAL.value = dp.GetAttr(yyS[yypt-0].value_string)
			dp.Next(yyS[yypt-0].value_string)
			fmt.Println("data:", yyS[yypt-0].value_string)
		}
	case 12:
		{
			yyVAL.value = yyS[yypt-0].value_string
			fmt.Println("Arrs:", yyS[yypt-0].value_string)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
