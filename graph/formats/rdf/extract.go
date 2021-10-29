
//line extract.rl:1
// Go code generated by go generate gonum.org/v1/gonum/graph/formats/rdf; DO NOT EDIT.

// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rdf

import (
	"fmt"
	"unicode"
)


//line extract.go:18
const extract_start int = 1
const extract_first_final int = 31
const extract_error int = 0

const extract_en_value int = 1


//line extract.rl:28


func extract(data []rune) (text, qual string, kind Kind, err error) {
	var (
		cs, p int
		pe    = len(data)
		eof   = pe

		iri     = -1
		blank   = -1
		literal = -1
		lang    = -1

		iriText     string
		blankText   string
		literalText string
		langText    string
	)

	
//line extract.go:47
	{
	cs = extract_start
	}

//line extract.rl:48

	
//line extract.go:55
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 31:
		goto st_case_31
	case 4:
		goto st_case_4
	case 32:
		goto st_case_32
	case 5:
		goto st_case_5
	case 33:
		goto st_case_33
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 34:
		goto st_case_34
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 35:
		goto st_case_35
	case 30:
		goto st_case_30
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 34:
			goto st2
		case 60:
			goto st8
		case 95:
			goto st28
		}
		goto tr0
tr0:
//line extract_actions.rl:74

		if p < len(data) {
			if r := data[p]; r < unicode.MaxASCII {
				return "", "", Invalid, fmt.Errorf("%w: unexpected rune %q at %d", ErrInvalidTerm, data[p], p)
			} else {
				return "", "", Invalid, fmt.Errorf("%w: unexpected rune %q (\\u%04[2]x) at %d", ErrInvalidTerm, data[p], p)
			}
		}
		return "", "", Invalid, ErrIncompleteTerm
	
	goto st0
//line extract.go:158
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 34:
			goto tr5
		case 92:
			goto tr6
		}
		switch {
		case data[p] < 11:
			if 0 <= data[p] && data[p] <= 9 {
				goto tr4
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 1114111 {
				goto tr4
			}
		default:
			goto tr4
		}
		goto tr0
tr4:
//line extract_actions.rl:34

		literal = p
	
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line extract.go:198
		switch data[p] {
		case 34:
			goto tr8
		case 92:
			goto st19
		}
		switch {
		case data[p] < 11:
			if 0 <= data[p] && data[p] <= 9 {
				goto st3
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 1114111 {
				goto st3
			}
		default:
			goto st3
		}
		goto tr0
tr5:
//line extract_actions.rl:34

		literal = p
	
//line extract_actions.rl:38

		if literal < 0 {
			panic("unexpected parser state: literal start not set")
		}
		literalText = unEscape(data[literal:p])
		kind = Literal
	
	goto st31
tr8:
//line extract_actions.rl:38

		if literal < 0 {
			panic("unexpected parser state: literal start not set")
		}
		literalText = unEscape(data[literal:p])
		kind = Literal
	
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line extract.go:247
		switch data[p] {
		case 64:
			goto tr39
		case 94:
			goto st6
		}
		goto st0
tr39:
//line extract_actions.rl:46

		lang = p
	
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line extract.go:266
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto tr0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 45 {
			goto st5
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		case data[p] >= 65:
			goto st32
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st33
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		default:
			goto st33
		}
		goto tr0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 45 {
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st33
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		default:
			goto st33
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 94 {
			goto st7
		}
		goto tr0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 60 {
			goto st8
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 62:
			goto tr14
		case 92:
			goto tr15
		case 95:
			goto tr13
		case 126:
			goto tr13
		}
		switch {
		case data[p] < 61:
			if 33 <= data[p] && data[p] <= 59 {
				goto tr13
			}
		case data[p] > 93:
			switch {
			case data[p] > 122:
				if 128 <= data[p] && data[p] <= 1114111 {
					goto tr13
				}
			case data[p] >= 97:
				goto tr13
			}
		default:
			goto tr13
		}
		goto tr0
tr13:
//line extract_actions.rl:8

		iri = p
	
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line extract.go:394
		switch data[p] {
		case 62:
			goto tr17
		case 92:
			goto st10
		case 95:
			goto st9
		case 126:
			goto st9
		}
		switch {
		case data[p] < 61:
			if 33 <= data[p] && data[p] <= 59 {
				goto st9
			}
		case data[p] > 93:
			switch {
			case data[p] > 122:
				if 128 <= data[p] && data[p] <= 1114111 {
					goto st9
				}
			case data[p] >= 97:
				goto st9
			}
		default:
			goto st9
		}
		goto tr0
tr14:
//line extract_actions.rl:8

		iri = p
	
//line extract_actions.rl:12

		if iri < 0 {
			panic("unexpected parser state: iri start not set")
		}
		iriText = unEscape(data[iri:p])
		if kind == Invalid {
			kind = IRI
		}
	
	goto st34
tr17:
//line extract_actions.rl:12

		if iri < 0 {
			panic("unexpected parser state: iri start not set")
		}
		iriText = unEscape(data[iri:p])
		if kind == Invalid {
			kind = IRI
		}
	
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line extract.go:456
		goto st0
tr15:
//line extract_actions.rl:8

		iri = p
	
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line extract.go:469
		switch data[p] {
		case 85:
			goto st11
		case 117:
			goto st15
		}
		goto tr0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st12
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st12
			}
		default:
			goto st12
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st14
			}
		default:
			goto st14
		}
		goto tr0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st15
			}
		default:
			goto st15
		}
		goto tr0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st17
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st17
			}
		default:
			goto st17
		}
		goto tr0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st18
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st18
			}
		default:
			goto st18
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st9
			}
		default:
			goto st9
		}
		goto tr0
tr6:
//line extract_actions.rl:34

		literal = p
	
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line extract.go:632
		switch data[p] {
		case 34:
			goto st3
		case 39:
			goto st3
		case 85:
			goto st20
		case 92:
			goto st3
		case 98:
			goto st3
		case 102:
			goto st3
		case 110:
			goto st3
		case 114:
			goto st3
		case 116:
			goto st3
		case 117:
			goto st24
		}
		goto tr0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st21
			}
		default:
			goto st21
		}
		goto tr0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st22
			}
		default:
			goto st22
		}
		goto tr0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st24
			}
		default:
			goto st24
		}
		goto tr0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st25
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st25
			}
		default:
			goto st25
		}
		goto tr0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st27
			}
		default:
			goto st27
		}
		goto tr0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st3
			}
		default:
			goto st3
		}
		goto tr0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 58 {
			goto st29
		}
		goto tr0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 95 {
			goto tr36
		}
		switch {
		case data[p] < 895:
			switch {
			case data[p] < 192:
				switch {
				case data[p] < 65:
					if 48 <= data[p] && data[p] <= 58 {
						goto tr36
					}
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr36
					}
				default:
					goto tr36
				}
			case data[p] > 214:
				switch {
				case data[p] < 248:
					if 216 <= data[p] && data[p] <= 246 {
						goto tr36
					}
				case data[p] > 767:
					if 880 <= data[p] && data[p] <= 893 {
						goto tr36
					}
				default:
					goto tr36
				}
			default:
				goto tr36
			}
		case data[p] > 8191:
			switch {
			case data[p] < 12289:
				switch {
				case data[p] < 8304:
					if 8204 <= data[p] && data[p] <= 8205 {
						goto tr36
					}
				case data[p] > 8591:
					if 11264 <= data[p] && data[p] <= 12271 {
						goto tr36
					}
				default:
					goto tr36
				}
			case data[p] > 55295:
				switch {
				case data[p] < 65008:
					if 63744 <= data[p] && data[p] <= 64975 {
						goto tr36
					}
				case data[p] > 65533:
					if 65536 <= data[p] && data[p] <= 983039 {
						goto tr36
					}
				default:
					goto tr36
				}
			default:
				goto tr36
			}
		default:
			goto tr36
		}
		goto tr0
tr36:
//line extract_actions.rl:22

		blank = p
	
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line extract.go:895
		switch data[p] {
		case 45:
			goto st35
		case 46:
			goto st30
		case 95:
			goto st35
		case 183:
			goto st35
		}
		switch {
		case data[p] < 8204:
			switch {
			case data[p] < 192:
				switch {
				case data[p] < 65:
					if 48 <= data[p] && data[p] <= 58 {
						goto st35
					}
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto st35
					}
				default:
					goto st35
				}
			case data[p] > 214:
				switch {
				case data[p] < 248:
					if 216 <= data[p] && data[p] <= 246 {
						goto st35
					}
				case data[p] > 893:
					if 895 <= data[p] && data[p] <= 8191 {
						goto st35
					}
				default:
					goto st35
				}
			default:
				goto st35
			}
		case data[p] > 8205:
			switch {
			case data[p] < 12289:
				switch {
				case data[p] < 8304:
					if 8255 <= data[p] && data[p] <= 8256 {
						goto st35
					}
				case data[p] > 8591:
					if 11264 <= data[p] && data[p] <= 12271 {
						goto st35
					}
				default:
					goto st35
				}
			case data[p] > 55295:
				switch {
				case data[p] < 65008:
					if 63744 <= data[p] && data[p] <= 64975 {
						goto st35
					}
				case data[p] > 65533:
					if 65536 <= data[p] && data[p] <= 983039 {
						goto st35
					}
				default:
					goto st35
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 45:
			goto st35
		case 46:
			goto st30
		case 95:
			goto st35
		case 183:
			goto st35
		}
		switch {
		case data[p] < 8204:
			switch {
			case data[p] < 192:
				switch {
				case data[p] < 65:
					if 48 <= data[p] && data[p] <= 58 {
						goto st35
					}
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto st35
					}
				default:
					goto st35
				}
			case data[p] > 214:
				switch {
				case data[p] < 248:
					if 216 <= data[p] && data[p] <= 246 {
						goto st35
					}
				case data[p] > 893:
					if 895 <= data[p] && data[p] <= 8191 {
						goto st35
					}
				default:
					goto st35
				}
			default:
				goto st35
			}
		case data[p] > 8205:
			switch {
			case data[p] < 12289:
				switch {
				case data[p] < 8304:
					if 8255 <= data[p] && data[p] <= 8256 {
						goto st35
					}
				case data[p] > 8591:
					if 11264 <= data[p] && data[p] <= 12271 {
						goto st35
					}
				default:
					goto st35
				}
			case data[p] > 55295:
				switch {
				case data[p] < 65008:
					if 63744 <= data[p] && data[p] <= 64975 {
						goto st35
					}
				case data[p] > 65533:
					if 65536 <= data[p] && data[p] <= 983039 {
						goto st35
					}
				default:
					goto st35
				}
			default:
				goto st35
			}
		default:
			goto st35
		}
		goto tr0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 31, 34:
//line extract_actions.rl:57

		switch kind {
		case IRI:
			return iriText, "", kind, nil
		case Blank:
			return blankText, "", kind, nil
		case Literal:
			qual = iriText
			if qual == "" {
				qual = langText
			}
			return literalText, qual, kind, nil
		default:
			return "", "", kind, ErrInvalidTerm
		}
	
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
//line extract_actions.rl:74

		if p < len(data) {
			if r := data[p]; r < unicode.MaxASCII {
				return "", "", Invalid, fmt.Errorf("%w: unexpected rune %q at %d", ErrInvalidTerm, data[p], p)
			} else {
				return "", "", Invalid, fmt.Errorf("%w: unexpected rune %q (\\u%04[2]x) at %d", ErrInvalidTerm, data[p], p)
			}
		}
		return "", "", Invalid, ErrIncompleteTerm
	
		case 35:
//line extract_actions.rl:26

		if blank < 0 {
			panic("unexpected parser state: blank start not set")
		}
		blankText = string(data[blank:p])
		kind = Blank
	
//line extract_actions.rl:57

		switch kind {
		case IRI:
			return iriText, "", kind, nil
		case Blank:
			return blankText, "", kind, nil
		case Literal:
			qual = iriText
			if qual == "" {
				qual = langText
			}
			return literalText, qual, kind, nil
		default:
			return "", "", kind, ErrInvalidTerm
		}
	
		case 32, 33:
//line extract_actions.rl:50

		if lang < 0 {
			panic("unexpected parser state: lang start not set")
		}
		langText = string(data[lang:p])
	
//line extract_actions.rl:57

		switch kind {
		case IRI:
			return iriText, "", kind, nil
		case Blank:
			return blankText, "", kind, nil
		case Literal:
			qual = iriText
			if qual == "" {
				qual = langText
			}
			return literalText, qual, kind, nil
		default:
			return "", "", kind, ErrInvalidTerm
		}
	
//line extract.go:1175
		}
	}

	_out: {}
	}

//line extract.rl:50

	return "", "", 0, ErrInvalidTerm
}
