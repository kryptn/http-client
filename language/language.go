package language

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// helpers
func ifs(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 14, col: 1, offset: 134},
			expr: &actionExpr{
				pos: position{line: 14, col: 13, offset: 146},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 14, col: 13, offset: 146},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 14, col: 13, offset: 146},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 14, col: 15, offset: 148},
							label: "block",
							expr: &ruleRefExpr{
								pos:  position{line: 14, col: 21, offset: 154},
								name: "Block",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 14, col: 27, offset: 160},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 14, col: 29, offset: 162},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Block",
			pos:  position{line: 18, col: 1, offset: 201},
			expr: &actionExpr{
				pos: position{line: 18, col: 10, offset: 210},
				run: (*parser).callonBlock1,
				expr: &seqExpr{
					pos: position{line: 18, col: 10, offset: 210},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 18, col: 10, offset: 210},
							label: "blockType",
							expr: &ruleRefExpr{
								pos:  position{line: 18, col: 20, offset: 220},
								name: "Word",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 18, col: 25, offset: 225},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 18, col: 27, offset: 227},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 18, col: 31, offset: 231},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 18, col: 33, offset: 233},
							label: "items",
							expr: &zeroOrMoreExpr{
								pos: position{line: 18, col: 39, offset: 239},
								expr: &choiceExpr{
									pos: position{line: 18, col: 40, offset: 240},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 18, col: 40, offset: 240},
											name: "KeyValuePair",
										},
										&ruleRefExpr{
											pos:  position{line: 18, col: 55, offset: 255},
											name: "Block",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 18, col: 63, offset: 263},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 18, col: 65, offset: 265},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "FunctionInvocation",
			pos:  position{line: 41, col: 1, offset: 826},
			expr: &actionExpr{
				pos: position{line: 41, col: 23, offset: 848},
				run: (*parser).callonFunctionInvocation1,
				expr: &seqExpr{
					pos: position{line: 41, col: 23, offset: 848},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 41, col: 23, offset: 848},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 41, col: 27, offset: 852},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 41, col: 29, offset: 854},
							label: "fnCall",
							expr: &oneOrMoreExpr{
								pos: position{line: 41, col: 36, offset: 861},
								expr: &ruleRefExpr{
									pos:  position{line: 41, col: 36, offset: 861},
									name: "FunctionCall",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 41, col: 50, offset: 875},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "FunctionCall",
			pos:  position{line: 50, col: 1, offset: 1074},
			expr: &actionExpr{
				pos: position{line: 50, col: 17, offset: 1090},
				run: (*parser).callonFunctionCall1,
				expr: &seqExpr{
					pos: position{line: 50, col: 17, offset: 1090},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 50, col: 17, offset: 1090},
							label: "fnName",
							expr: &ruleRefExpr{
								pos:  position{line: 50, col: 24, offset: 1097},
								name: "Name",
							},
						},
						&labeledExpr{
							pos:   position{line: 50, col: 29, offset: 1102},
							label: "argsi",
							expr: &zeroOrMoreExpr{
								pos: position{line: 50, col: 35, offset: 1108},
								expr: &ruleRefExpr{
									pos:  position{line: 50, col: 35, offset: 1108},
									name: "Argument",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 50, col: 45, offset: 1118},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 50, col: 47, offset: 1120},
							expr: &seqExpr{
								pos: position{line: 50, col: 48, offset: 1121},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 50, col: 48, offset: 1121},
										val:        "|",
										ignoreCase: false,
									},
									&ruleRefExpr{
										pos:  position{line: 50, col: 52, offset: 1125},
										name: "_",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Argument",
			pos:  position{line: 60, col: 1, offset: 1376},
			expr: &actionExpr{
				pos: position{line: 60, col: 13, offset: 1388},
				run: (*parser).callonArgument1,
				expr: &seqExpr{
					pos: position{line: 60, col: 13, offset: 1388},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 60, col: 13, offset: 1388},
							expr: &litMatcher{
								pos:        position{line: 60, col: 13, offset: 1388},
								val:        " ",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 60, col: 18, offset: 1393},
							label: "v",
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 20, offset: 1395},
								name: "ArgumentValue",
							},
						},
					},
				},
			},
		},
		{
			name: "KeyValuePair",
			pos:  position{line: 68, col: 1, offset: 1525},
			expr: &actionExpr{
				pos: position{line: 68, col: 17, offset: 1541},
				run: (*parser).callonKeyValuePair1,
				expr: &seqExpr{
					pos: position{line: 68, col: 17, offset: 1541},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 68, col: 17, offset: 1541},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 68, col: 19, offset: 1543},
							label: "k",
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 21, offset: 1545},
								name: "Key",
							},
						},
						&litMatcher{
							pos:        position{line: 68, col: 25, offset: 1549},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 68, col: 29, offset: 1553},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 68, col: 31, offset: 1555},
							label: "v",
							expr: &oneOrMoreExpr{
								pos: position{line: 68, col: 33, offset: 1557},
								expr: &ruleRefExpr{
									pos:  position{line: 68, col: 33, offset: 1557},
									name: "Value",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 68, col: 40, offset: 1564},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 78, col: 1, offset: 1786},
			expr: &actionExpr{
				pos: position{line: 78, col: 9, offset: 1794},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 78, col: 9, offset: 1794},
					expr: &charClassMatcher{
						pos:        position{line: 78, col: 9, offset: 1794},
						val:        "[0-9a-z_-]i",
						chars:      []rune{'_', '-'},
						ranges:     []rune{'0', '9', 'a', 'z'},
						ignoreCase: true,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "Name",
			pos:  position{line: 82, col: 1, offset: 1843},
			expr: &actionExpr{
				pos: position{line: 82, col: 9, offset: 1851},
				run: (*parser).callonName1,
				expr: &seqExpr{
					pos: position{line: 82, col: 9, offset: 1851},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 82, col: 9, offset: 1851},
							val:        "[a-z]i",
							ranges:     []rune{'a', 'z'},
							ignoreCase: true,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 82, col: 16, offset: 1858},
							expr: &charClassMatcher{
								pos:        position{line: 82, col: 16, offset: 1858},
								val:        "[0-9a-z_-]i",
								chars:      []rune{'_', '-'},
								ranges:     []rune{'0', '9', 'a', 'z'},
								ignoreCase: true,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Words",
			pos:  position{line: 86, col: 1, offset: 1907},
			expr: &oneOrMoreExpr{
				pos: position{line: 86, col: 10, offset: 1916},
				expr: &choiceExpr{
					pos: position{line: 86, col: 11, offset: 1917},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 86, col: 11, offset: 1917},
							name: "Word",
						},
						&oneOrMoreExpr{
							pos: position{line: 86, col: 18, offset: 1924},
							expr: &charClassMatcher{
								pos:        position{line: 86, col: 18, offset: 1924},
								val:        "[ \\t]",
								chars:      []rune{' ', '\t'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Key",
			pos:  position{line: 88, col: 1, offset: 1934},
			expr: &actionExpr{
				pos: position{line: 88, col: 8, offset: 1941},
				run: (*parser).callonKey1,
				expr: &labeledExpr{
					pos:   position{line: 88, col: 8, offset: 1941},
					label: "w",
					expr: &ruleRefExpr{
						pos:  position{line: 88, col: 10, offset: 1943},
						name: "Word",
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 94, col: 1, offset: 2031},
			expr: &choiceExpr{
				pos: position{line: 94, col: 10, offset: 2040},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 94, col: 10, offset: 2040},
						run: (*parser).callonValue2,
						expr: &labeledExpr{
							pos:   position{line: 94, col: 10, offset: 2040},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 13, offset: 2043},
								name: "FunctionInvocation",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 97, col: 5, offset: 2173},
						name: "MultiLineValue",
					},
					&ruleRefExpr{
						pos:  position{line: 97, col: 22, offset: 2190},
						name: "QuotedValue",
					},
					&actionExpr{
						pos: position{line: 97, col: 36, offset: 2204},
						run: (*parser).callonValue7,
						expr: &ruleRefExpr{
							pos:  position{line: 97, col: 36, offset: 2204},
							name: "Words",
						},
					},
				},
			},
		},
		{
			name: "QuotedValue",
			pos:  position{line: 102, col: 1, offset: 2290},
			expr: &actionExpr{
				pos: position{line: 102, col: 16, offset: 2305},
				run: (*parser).callonQuotedValue1,
				expr: &seqExpr{
					pos: position{line: 102, col: 16, offset: 2305},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 16, offset: 2305},
							val:        "\"",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 102, col: 21, offset: 2310},
							label: "contents",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 30, offset: 2319},
								name: "QuotedContents",
							},
						},
						&litMatcher{
							pos:        position{line: 102, col: 45, offset: 2334},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedContents",
			pos:  position{line: 106, col: 1, offset: 2377},
			expr: &actionExpr{
				pos: position{line: 106, col: 19, offset: 2395},
				run: (*parser).callonQuotedContents1,
				expr: &seqExpr{
					pos: position{line: 106, col: 19, offset: 2395},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 106, col: 19, offset: 2395},
							expr: &seqExpr{
								pos: position{line: 106, col: 20, offset: 2396},
								exprs: []interface{}{
									&anyMatcher{
										line: 106, col: 20, offset: 2396,
									},
									&notExpr{
										pos: position{line: 106, col: 22, offset: 2398},
										expr: &litMatcher{
											pos:        position{line: 106, col: 23, offset: 2399},
											val:        "\"",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&anyMatcher{
							line: 106, col: 29, offset: 2405,
						},
					},
				},
			},
		},
		{
			name: "MultiLineValue",
			pos:  position{line: 110, col: 1, offset: 2460},
			expr: &actionExpr{
				pos: position{line: 110, col: 19, offset: 2478},
				run: (*parser).callonMultiLineValue1,
				expr: &seqExpr{
					pos: position{line: 110, col: 19, offset: 2478},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 110, col: 19, offset: 2478},
							name: "tplq",
						},
						&labeledExpr{
							pos:   position{line: 110, col: 24, offset: 2483},
							label: "contents",
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 33, offset: 2492},
								name: "MultiLineContents",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 110, col: 51, offset: 2510},
							name: "tplq",
						},
					},
				},
			},
		},
		{
			name: "MultiLineContents",
			pos:  position{line: 115, col: 1, offset: 2594},
			expr: &actionExpr{
				pos: position{line: 115, col: 22, offset: 2615},
				run: (*parser).callonMultiLineContents1,
				expr: &seqExpr{
					pos: position{line: 115, col: 22, offset: 2615},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 115, col: 22, offset: 2615},
							expr: &seqExpr{
								pos: position{line: 115, col: 23, offset: 2616},
								exprs: []interface{}{
									&anyMatcher{
										line: 115, col: 23, offset: 2616,
									},
									&notExpr{
										pos: position{line: 115, col: 25, offset: 2618},
										expr: &ruleRefExpr{
											pos:  position{line: 115, col: 26, offset: 2619},
											name: "tplq",
										},
									},
								},
							},
						},
						&anyMatcher{
							line: 115, col: 33, offset: 2626,
						},
					},
				},
			},
		},
		{
			name: "ArgumentValue",
			pos:  position{line: 120, col: 1, offset: 2714},
			expr: &choiceExpr{
				pos: position{line: 120, col: 18, offset: 2731},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 120, col: 18, offset: 2731},
						run: (*parser).callonArgumentValue2,
						expr: &labeledExpr{
							pos:   position{line: 120, col: 18, offset: 2731},
							label: "fn",
							expr: &ruleRefExpr{
								pos:  position{line: 120, col: 21, offset: 2734},
								name: "FunctionInvocation",
							},
						},
					},
					&actionExpr{
						pos: position{line: 123, col: 5, offset: 2872},
						run: (*parser).callonArgumentValue5,
						expr: &ruleRefExpr{
							pos:  position{line: 123, col: 5, offset: 2872},
							name: "Word",
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 128, col: 1, offset: 2965},
			expr: &notExpr{
				pos: position{line: 128, col: 8, offset: 2972},
				expr: &anyMatcher{
					line: 128, col: 9, offset: 2973,
				},
			},
		},
		{
			name: "sp",
			pos:  position{line: 130, col: 1, offset: 2976},
			expr: &actionExpr{
				pos: position{line: 130, col: 7, offset: 2982},
				run: (*parser).callonsp1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 130, col: 7, offset: 2982},
					expr: &charClassMatcher{
						pos:        position{line: 130, col: 7, offset: 2982},
						val:        "[ \\t]",
						chars:      []rune{' ', '\t'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "nl",
			pos:  position{line: 134, col: 1, offset: 3014},
			expr: &actionExpr{
				pos: position{line: 134, col: 7, offset: 3020},
				run: (*parser).callonnl1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 134, col: 7, offset: 3020},
					expr: &charClassMatcher{
						pos:        position{line: 134, col: 7, offset: 3020},
						val:        "[\\r\\n]",
						chars:      []rune{'\r', '\n'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "tplq",
			pos:  position{line: 138, col: 1, offset: 3053},
			expr: &actionExpr{
				pos: position{line: 138, col: 9, offset: 3061},
				run: (*parser).callontplq1,
				expr: &litMatcher{
					pos:        position{line: 138, col: 9, offset: 3061},
					val:        "\"\"\"",
					ignoreCase: false,
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 143, col: 1, offset: 3114},
			expr: &actionExpr{
				pos: position{line: 143, col: 19, offset: 3132},
				run: (*parser).callon_1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 143, col: 19, offset: 3132},
					expr: &charClassMatcher{
						pos:        position{line: 143, col: 19, offset: 3132},
						val:        "[ \\t\\r\\n]",
						chars:      []rune{' ', '\t', '\r', '\n'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
	},
}

func (c *current) onDocument1(block interface{}) (interface{}, error) {
	return block.(Block), nil
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["block"])
}

func (c *current) onBlock1(blockType, items interface{}) (interface{}, error) {

	keyValuePairs := []KeyValuePair{}
	subBlocks := []Block{}

	for _, item := range ifs(items) {
		switch v := item.(type) {
		case KeyValuePair:
			keyValuePairs = append(keyValuePairs, v)
		case Block:
			subBlocks = append(subBlocks, v)
		default:
			return nil, errors.New("expected Block or KeyValuePair") // todo: does this need more?
		}
	}

	return Block{
		BlockType:     blockType.(string),
		KeyValuePairs: keyValuePairs,
		SubBlocks:     subBlocks,
	}, nil
}

func (p *parser) callonBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlock1(stack["blockType"], stack["items"])
}

func (c *current) onFunctionInvocation1(fnCall interface{}) (interface{}, error) {
	fns := []FunctionInvocation{}
	dumps("FunctionInvocation", c, fnCall)
	for _, fn := range ifs(fnCall) {
		fns = append(fns, fn.(FunctionInvocation))
	}
	return fns, nil
}

func (p *parser) callonFunctionInvocation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFunctionInvocation1(stack["fnCall"])
}

func (c *current) onFunctionCall1(fnName, argsi interface{}) (interface{}, error) {
	dumps("FunctionCall", c, c.text, fnName, argsi)
	args := []Argument{}
	for _, arg := range ifs(argsi) {
		args = append(args, arg.(Argument))

	}
	return FunctionInvocation{FuncName: fnName.(string), Arguments: args}, nil
}

func (p *parser) callonFunctionCall1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFunctionCall1(stack["fnName"], stack["argsi"])
}

func (c *current) onArgument1(v interface{}) (interface{}, error) {
	return Argument{
		StrValue:  v.(Value).StrValue,
		FuncValue: v.(Value).FuncValue,
	}, nil
}

func (p *parser) callonArgument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgument1(stack["v"])
}

func (c *current) onKeyValuePair1(k, v interface{}) (interface{}, error) {
	dumps("KeyValuePair", c, k, v)
	values := []Value{}
	for _, value := range v.([]interface{}) {
		values = append(values, value.(Value))
	}

	return KeyValuePair{Key: k.(Key), Values: values}, nil
}

func (p *parser) callonKeyValuePair1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onKeyValuePair1(stack["k"], stack["v"])
}

func (c *current) onWord1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWord1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWord1()
}

func (c *current) onName1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onName1()
}

func (c *current) onKey1(w interface{}) (interface{}, error) {
	dumps("Key", c, w)
	return Key(w.(string)), nil
}

func (p *parser) callonKey1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onKey1(stack["w"])
}

func (c *current) onValue2(fn interface{}) (interface{}, error) {
	dumps("Value FunctionInvocation", c, fn)
	return Value{FuncValue: fn.([]FunctionInvocation)}, nil
}

func (p *parser) callonValue2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue2(stack["fn"])
}

func (c *current) onValue7() (interface{}, error) {
	dumps("Value Word", c)
	return Value{StrValue: string(c.text)}, nil
}

func (p *parser) callonValue7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue7()
}

func (c *current) onQuotedValue1(contents interface{}) (interface{}, error) {
	return contents.(Value), nil
}

func (p *parser) callonQuotedValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuotedValue1(stack["contents"])
}

func (c *current) onQuotedContents1() (interface{}, error) {
	return Value{StrValue: string(c.text)}, nil
}

func (p *parser) callonQuotedContents1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuotedContents1()
}

func (c *current) onMultiLineValue1(contents interface{}) (interface{}, error) {
	dumps("MultiLineValue", c, contents)
	return contents.(Value), nil
}

func (p *parser) callonMultiLineValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMultiLineValue1(stack["contents"])
}

func (c *current) onMultiLineContents1() (interface{}, error) {
	return Value{StrValue: string(c.text)}, nil
}

func (p *parser) callonMultiLineContents1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMultiLineContents1()
}

func (c *current) onArgumentValue2(fn interface{}) (interface{}, error) {
	dumps("ArgumentValue FunctionInvocation", c, fn)
	return Value{FuncValue: fn.([]FunctionInvocation)}, nil
}

func (p *parser) callonArgumentValue2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgumentValue2(stack["fn"])
}

func (c *current) onArgumentValue5() (interface{}, error) {
	dumps("ArgumentValue Word", c)
	return Value{StrValue: string(c.text)}, nil
}

func (p *parser) callonArgumentValue5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgumentValue5()
}

func (c *current) onsp1() (interface{}, error) {
	return nil, nil
}

func (p *parser) callonsp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsp1()
}

func (c *current) onnl1() (interface{}, error) {
	return nil, nil
}

func (p *parser) callonnl1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onnl1()
}

func (c *current) ontplq1() (interface{}, error) {
	dumps("tplq", c)
	return c, nil
}

func (p *parser) callontplq1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.ontplq1()
}

func (c *current) on_1() (interface{}, error) {
	return nil, nil
}

func (p *parser) callon_1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.on_1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
