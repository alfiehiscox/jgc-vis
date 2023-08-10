package parser

import (
	"bytes"
	"errors"
	"fmt"
)

// === Token Constants ===
type TokenPair struct {
	Token   Token
	Literal string
}

type Token int

const (
	OPEN_SQUARE Token = iota
	CLOSE_SQUARE
	OPEN_PAREN
	CLOSE_PAREN
	COMMA
	COLON
	EQUAL

	LABEL
	SIZE
	TIME
	ARROW

	TIMESTAMP

	EOF

	GC
	FULL_GC

	NUL
)

func Tokenize(log string) []TokenPair {
	var tps []TokenPair

	// Provides a callback via a closure to pop off TokenPairs by 'scan' function
	pop := func() {
		if len(tps) > 0 {
			tps = tps[:len(tps)-1]
		}
	}

	s := newStream(log)
	for !s.eof() {
		t, l := s.scan(pop)
		tps = append(tps, TokenPair{t, l})
	}

	return tps
}

// ==== Stream ====
type stream struct {
	input []rune

	pos  int
	line int
	col  int
}

func newStream(input string) *stream {
	return &stream{
		input: []rune(input),
		pos:   0,
		line:  0,
		col:   0,
	}
}

func (s *stream) peek() rune {
	return s.input[s.pos]
}

func (s *stream) next() rune {
	r := s.input[s.pos]
	s.pos++
	if r == '\n' {
		s.line++
		s.col = 0
	} else {
		s.col++
	}
	return r
}

func (s *stream) eof() bool {
	return s.pos == len(s.input)
}

func (s *stream) croak(msg string) error {
	return errors.New(fmt.Sprintf("%v (%d:%d)", msg, s.line, s.col))
}

func (s *stream) dPeek() rune {
	return s.input[s.pos+1]
}

func (s *stream) retreat(by int) rune {
	if s.pos-by < 0 {
		s.pos = 0
		return s.input[s.pos]
	}
	s.pos = (s.pos - by)
	return s.input[s.pos]
}

func (s *stream) scan(pop func()) (Token, string) {
	if s.eof() {
		return EOF, ""
	}

	r := s.peek()

	// Multi Characters
	if isLetter(r) {
		return s.scanText()
	} else if isWhitespace(r) {
		s.scanWhiteSpace()
		return s.scan(pop)
	} else if isNumber(r) {
		return s.scanNumbers()
	} else if isHyphen(r) {
		if s.dPeek() == '>' {
			var aBuf bytes.Buffer
			for i := 0; i < 2; i++ {
				aBuf.WriteRune(s.next())
			}
			return ARROW, aBuf.String()
		} else if isNumber(s.dPeek()) {
			// PARSE TIMESTAMP
			// TIMESTAMP LAYOUT IS ALWAYS: 2023-08-10T11:09:31.795+0000
			pop()        // Get rid of 'yyyy' that has already been added as token
			s.retreat(4) // Back to beginning
			return s.scanTimestamp()
		}
	}

	// Single Characters
	switch r {
	case '[':
		return OPEN_SQUARE, string(s.next())
	case ']':
		return CLOSE_SQUARE, string(s.next())
	case '(':
		return OPEN_PAREN, string(s.next())
	case ')':
		return CLOSE_PAREN, string(s.next())
	case ',':
		return COMMA, string(s.next())
	case ':':
		return COLON, string(s.next())
	case '=':
		return EQUAL, string(s.next())
	}

	return NUL, ""
}

func (s *stream) scanWhiteSpace() {
	for {
		if s.eof() || !isWhitespace(s.peek()) {
			break
		} else {
			s.next()
		}
	}
}

func (s *stream) scanTimestamp() (Token, string) {
	var buf bytes.Buffer

	// Increment by 28 character as format is constant for now: 2023-08-10T11:09:31.795+0000
	for i := 0; i < 28; i++ {
		buf.WriteRune(s.next())
	}

	return TIMESTAMP, buf.String()
}

func (s *stream) scanText() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.next())

	// We allow for two words with a space in between
	for {
		// fmt.Println(buf.String())
		if s.eof() {
			break
		} else if s.peek() == ' ' {
			if isLetter(s.dPeek()) {
				buf.WriteRune(s.next()) // The ' '
				for {
					if !s.eof() && isLetter(s.peek()) {
						buf.WriteRune(s.next())
					} else {
						break
					}
				}
			} else {
				break
			}
		} else if s.peek() == '.' {
			buf.WriteRune(s.next())
		} else if !isLetter(s.peek()) {
			break
		} else {
			buf.WriteRune(s.next())
		}
	}

	switch buf.String() {
	case "GC":
		return GC, buf.String()
	case "Full GC":
		return FULL_GC, buf.String()
	}

	return LABEL, buf.String()
}

// Only ever SIZE or TIME
func (s *stream) scanNumbers() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.next())

	for {
		if s.eof() {
			break
		} else if s.peek() == 'K' {
			buf.WriteRune(s.next())
			return SIZE, buf.String()
		} else if s.peek() != '.' && !isNumber(s.peek()) {
			break
		} else {
			buf.WriteRune(s.next())
		}
	}

	return TIME, buf.String()
}

// === Utility Functions ===
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isHyphen(ch rune) bool {
	return ch == '-'
}
