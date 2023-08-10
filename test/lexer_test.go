package main

import (
	"testing"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
)

func TestJava8TokeniserWithEventSyntax(t *testing.T) {
	log := "8704K->992K(9728K)"

	expected := []parser.TokenPair{
		{Token: parser.SIZE, Literal: "8704K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "992K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "9728K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
	}

	tokenPairs := parser.Tokenize(log)

	if len(expected) != len(tokenPairs) {
		t.Fatalf("Length does not match: Expected len=%d, Got len=%d", len(expected), len(tokenPairs))
	}

	for i, v := range tokenPairs {
		if v.Token != expected[i].Token {
			t.Errorf("Token does not match: Expected token=%v, Got token=%v", expected[i].Token, v.Token)
		}

		if v.Literal != expected[i].Literal {
			t.Errorf("Literal does not match: Expected literal=%v, Got literal=%v", expected[i].Literal, v.Literal)
		}
	}

}

func TestJava8TokeniserWithMinorGCLog(t *testing.T) {
	log := "[GC (Allocation Failure) [PSYoungGen: 8704K->992K(9728K)] 8704K->3776K(31744K), 0.0073015 secs] [Times: user=0.02 sys=0.00, real=0.01 secs]"

	expected := []parser.TokenPair{
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.GC, Literal: "GC"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.LABEL, Literal: "Allocation Failure"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "PSYoungGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "8704K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "992K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "9728K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.SIZE, Literal: "8704K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "3776K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "31744K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.TIME, Literal: "0.0073015"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "Times"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.LABEL, Literal: "user"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.02"},
		{Token: parser.LABEL, Literal: "sys"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.00"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.LABEL, Literal: "real"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.01"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
	}

	tokenPairs := parser.Tokenize(log)

	if len(expected) != len(tokenPairs) {
		t.Fatalf("Length does not match: Expected len=%d, Got len=%d", len(expected), len(tokenPairs))
	}

	for i, v := range tokenPairs {
		if v.Token != expected[i].Token {
			t.Errorf("Token does not match: Expected token=%v, Got token=%v", expected[i].Token, v.Token)
		}

		if v.Literal != expected[i].Literal {
			t.Errorf("Literal does not match: Expected literal=%v, Got literal=%v", expected[i].Literal, v.Literal)
		}
	}
}

func TestJava8TokeniserWithMajorGCLog(t *testing.T) {
	log := "[Full GC (Ergonomics) [PSYoungGen: 57344K->0K(113664K)] [ParOldGen: 337435K->261246K(339968K)] 394779K->261246K(453632K), [Metaspace: 2866K->2866K(1056768K)], 0.3415608 secs] [Times: user=1.26 sys=0.00, real=0.35 secs]"

	expected := []parser.TokenPair{
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.FULL_GC, Literal: "Full GC"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.LABEL, Literal: "Ergonomics"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "PSYoungGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "57344K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "0K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "113664K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "ParOldGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "337435K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "261246K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "339968K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.SIZE, Literal: "394779K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "261246K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "453632K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "Metaspace"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "2866K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "2866K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "1056768K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.TIME, Literal: "0.3415608"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "Times"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.LABEL, Literal: "user"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "1.26"},
		{Token: parser.LABEL, Literal: "sys"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.00"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.LABEL, Literal: "real"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.35"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
	}

	tokenPairs := parser.Tokenize(log)

	if len(expected) != len(tokenPairs) {
		t.Fatalf("Length does not match: Expected len=%d, Got len=%d", len(expected), len(tokenPairs))
	}

	for i, v := range tokenPairs {
		if v.Token != expected[i].Token {
			t.Errorf("Token does not match: Expected token=%v, Got token=%v", expected[i].Token, v.Token)
		}

		if v.Literal != expected[i].Literal {
			t.Errorf("Literal does not match: Expected literal=%v, Got literal=%v", expected[i].Literal, v.Literal)
		}
	}
}

func TestJava8TokeniserWithSystemGC(t *testing.T) {
	log := "[GC (System.gc()) [PSYoungGen: 0K->0K(0K)]]"

	expected := []parser.TokenPair{
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.GC, Literal: "GC"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.LABEL, Literal: "System.gc"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "PSYoungGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "0K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "0K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "0K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
	}

	tokenPairs := parser.Tokenize(log)

	if len(expected) != len(tokenPairs) {
		t.Errorf("Expected: %v", expected)
		t.Errorf("Actual  : %v", tokenPairs)
		t.Fatalf("Length does not match: Expected len=%d, Got len=%d", len(expected), len(tokenPairs))
	}

	for i, v := range tokenPairs {
		if v.Token != expected[i].Token {
			t.Errorf("Token does not match: Expected token=%v, Got token=%v", expected[i].Token, v.Token)
		}

		if v.Literal != expected[i].Literal {
			t.Errorf("Literal does not match: Expected literal=%v, Got literal=%v", expected[i].Literal, v.Literal)
		}
	}
}

func TestJava8TokeniserMinorGCWithTimestamp(t *testing.T) {
	log := "2023-08-10T11:09:31.795+0000: [GC (Allocation Failure) [PSYoungGen: 73162K->46048K(90112K)] 215235K->215465K(301056K), 0.0482088 secs] [Times: user=0.17 sys=0.02, real=0.05 secs]"

	expected := []parser.TokenPair{
		{Token: parser.TIMESTAMP, Literal: "2023-08-10T11:09:31.795+0000"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.GC, Literal: "GC"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.LABEL, Literal: "Allocation Failure"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "PSYoungGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "73162K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "46048K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "90112K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.SIZE, Literal: "215235K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "215465K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "301056K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.TIME, Literal: "0.0482088"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "Times"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.LABEL, Literal: "user"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.17"},
		{Token: parser.LABEL, Literal: "sys"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.02"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.LABEL, Literal: "real"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.05"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
	}

	tokenPairs := parser.Tokenize(log)

	if len(expected) != len(tokenPairs) {
		t.Errorf("Expected: %v", expected)
		t.Errorf("Actual  : %v", tokenPairs)
		t.Fatalf("Length does not match: Expected len=%d, Got len=%d", len(expected), len(tokenPairs))
	}

	for i, v := range tokenPairs {
		if v.Token != expected[i].Token {
			t.Errorf("Token does not match: Expected token=%v, Got token=%v", expected[i].Token, v.Token)
		}

		if v.Literal != expected[i].Literal {
			t.Errorf("Literal does not match: Expected literal=%v, Got literal=%v", expected[i].Literal, v.Literal)
		}
	}
}

func TestJava8TokeniserMajorGCWithTimestamp(t *testing.T) {
	log := "2023-08-10T11:09:31.795+0000: [Full GC (Ergonomics) [PSYoungGen: 57344K->0K(113664K)] [ParOldGen: 337435K->261246K(339968K)] 394779K->261246K(453632K), [Metaspace: 2866K->2866K(1056768K)], 0.3415608 secs] [Times: user=1.26 sys=0.00, real=0.35 secs]"

	expected := []parser.TokenPair{
		{Token: parser.TIMESTAMP, Literal: "2023-08-10T11:09:31.795+0000"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.FULL_GC, Literal: "Full GC"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.LABEL, Literal: "Ergonomics"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "PSYoungGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "57344K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "0K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "113664K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "ParOldGen"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "337435K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "261246K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "339968K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.SIZE, Literal: "394779K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "261246K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "453632K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "Metaspace"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.SIZE, Literal: "2866K"},
		{Token: parser.ARROW, Literal: "->"},
		{Token: parser.SIZE, Literal: "2866K"},
		{Token: parser.OPEN_PAREN, Literal: "("},
		{Token: parser.SIZE, Literal: "1056768K"},
		{Token: parser.CLOSE_PAREN, Literal: ")"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.TIME, Literal: "0.3415608"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
		{Token: parser.OPEN_SQUARE, Literal: "["},
		{Token: parser.LABEL, Literal: "Times"},
		{Token: parser.COLON, Literal: ":"},
		{Token: parser.LABEL, Literal: "user"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "1.26"},
		{Token: parser.LABEL, Literal: "sys"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.00"},
		{Token: parser.COMMA, Literal: ","},
		{Token: parser.LABEL, Literal: "real"},
		{Token: parser.EQUAL, Literal: "="},
		{Token: parser.TIME, Literal: "0.35"},
		{Token: parser.LABEL, Literal: "secs"},
		{Token: parser.CLOSE_SQUARE, Literal: "]"},
	}

	tokenPairs := parser.Tokenize(log)

	if len(expected) != len(tokenPairs) {
		t.Fatalf("Length does not match: Expected len=%d, Got len=%d", len(expected), len(tokenPairs))
	}

	for i, v := range tokenPairs {
		if v.Token != expected[i].Token {
			t.Errorf("Token does not match: Expected token=%v, Got token=%v", expected[i].Token, v.Token)
		}

		if v.Literal != expected[i].Literal {
			t.Errorf("Literal does not match: Expected literal=%v, Got literal=%v", expected[i].Literal, v.Literal)
		}
	}

}
