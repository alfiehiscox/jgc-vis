/*
package parser holds the parser for serializing the logs
from the java output to a format we can use in analysis.
*/
package parser

type token int

// Tokens
const (
	open_square  token = iota // [
	close_square              // ]
	open_parem                // (
	close_parem               // )
	arrow                     // ->
	equals                    // =
	colon                     // :
	comma                     // ,
	eof                       // Null

	literal

	ws // 1+ whitespace

	// KEY WORDS
	full_gc // 'Full GC'
	gc      // 'GC'
	seconds // 'Secs'
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLeter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
