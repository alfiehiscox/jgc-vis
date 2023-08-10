package main

import (
	"testing"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
)

func TestJava8TokeniserWithEventSyntax(t *testing.T) {
	log := "8704K->992K(9728K)"

	expected := []parser.TokenPair{
		{parser.SIZE, "8704K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "992K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "9728K"},
		{parser.CLOSE_PAREN, ")"},
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
		{parser.OPEN_SQUARE, "["},
		{parser.GC, "GC"},
		{parser.OPEN_PAREN, "("},
		{parser.LABEL, "Allocation Failure"},
		{parser.CLOSE_PAREN, ")"},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "PSYoungGen"},
		{parser.COLON, ":"},
		{parser.SIZE, "8704K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "992K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "9728K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.SIZE, "8704K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "3776K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "31744K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.COMMA, ","},
		{parser.TIME, "0.0073015"},
		{parser.LABEL, "secs"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "Times"},
		{parser.COLON, ":"},
		{parser.LABEL, "user"},
		{parser.EQUAL, "="},
		{parser.TIME, "0.02"},
		{parser.LABEL, "sys"},
		{parser.EQUAL, "="},
		{parser.TIME, "0.00"},
		{parser.COMMA, ","},
		{parser.LABEL, "real"},
		{parser.EQUAL, "="},
		{parser.TIME, "0.01"},
		{parser.LABEL, "secs"},
		{parser.CLOSE_SQUARE, "]"},
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
		{parser.OPEN_SQUARE, "["},
		{parser.FULL_GC, "Full GC"},
		{parser.OPEN_PAREN, "("},
		{parser.LABEL, "Ergonomics"},
		{parser.CLOSE_PAREN, ")"},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "PSYoungGen"},
		{parser.COLON, ":"},
		{parser.SIZE, "57344K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "0K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "113664K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "ParOldGen"},
		{parser.COLON, ":"},
		{parser.SIZE, "337435K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "261246K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "339968K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.SIZE, "394779K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "261246K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "453632K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.COMMA, ","},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "Metaspace"},
		{parser.COLON, ":"},
		{parser.SIZE, "2866K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "2866K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "1056768K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.COMMA, ","},
		{parser.TIME, "0.3415608"},
		{parser.LABEL, "secs"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "Times"},
		{parser.COLON, ":"},
		{parser.LABEL, "user"},
		{parser.EQUAL, "="},
		{parser.TIME, "1.26"},
		{parser.LABEL, "sys"},
		{parser.EQUAL, "="},
		{parser.TIME, "0.00"},
		{parser.COMMA, ","},
		{parser.LABEL, "real"},
		{parser.EQUAL, "="},
		{parser.TIME, "0.35"},
		{parser.LABEL, "secs"},
		{parser.CLOSE_SQUARE, "]"},
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
		{parser.OPEN_SQUARE, "["},
		{parser.GC, "GC"},
		{parser.OPEN_PAREN, "("},
		{parser.LABEL, "System.gc"},
		{parser.OPEN_PAREN, "("},
		{parser.CLOSE_PAREN, ")"},
		{parser.CLOSE_PAREN, ")"},
		{parser.OPEN_SQUARE, "["},
		{parser.LABEL, "PSYoungGen"},
		{parser.COLON, ":"},
		{parser.SIZE, "0K"},
		{parser.ARROW, "->"},
		{parser.SIZE, "0K"},
		{parser.OPEN_PAREN, "("},
		{parser.SIZE, "0K"},
		{parser.CLOSE_PAREN, ")"},
		{parser.CLOSE_SQUARE, "]"},
		{parser.CLOSE_SQUARE, "]"},
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
