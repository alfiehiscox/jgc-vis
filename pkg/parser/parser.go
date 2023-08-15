package parser

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const TIMESTAMP_LAYOUT = "2006-01-02T15:04:05-0700"

// === Main Structures ===
type GCEvent struct {
	BeforeSize int // In KB
	AfterSize  int // In KB
	TotalSize  int // In KB
}

type GCLog struct {
	Timestamp time.Time // Only if DateTimeStamps are enabled
	Type      string    // GC or Full GC
	Reason    string
	MainEvent GCEvent
	Time      string // In Seconds
	GenEvents []struct {
		Type  string
		Event GCEvent
	}
}

// === Parser ===
type Parser struct {
	TokenPairs []TokenPair
	Tokens     []Token
	Pos        int
}

func NewParser(log string) *Parser {
	tkps := Tokenize(log)
	var tks []Token

	for _, tp := range tkps {
		tks = append(tks, tp.Token)
	}

	return &Parser{
		Tokens:     tks,
		TokenPairs: tkps,
		Pos:        0,
	}
}

func (p *Parser) Parse() (*GCLog, error) {
	var gcLog GCLog

	for !p.eof() {
		if p.peek() == OPEN_SQUARE {
			p.next() // '['

			if p.peek() == GC || p.peek() == FULL_GC {
				gcLog.Type = p.nextPair().Literal

				// Parse Reason
				p.next()                            // '('
				gcLog.Reason = p.nextPair().Literal // 'Reason'
				p.next()                            // ')'
			} else if p.peek() == LABEL && p.dPeek() == COLON {
				genEvent := struct {
					Type  string
					Event GCEvent
				}{}
				genEvent.Type = p.nextPair().Literal

				// TODO: Fix this
				if genEvent.Type == "Times" {
					continue
				}

				p.next() // Throw away the ':'
				event, err := p.parseEvent()
				if err != nil {
					return nil, err
				}
				genEvent.Event = *event
				gcLog.GenEvents = append(gcLog.GenEvents, genEvent)
			}
		} else if p.peek() == CLOSE_SQUARE && p.dPeek() == SIZE {
			p.next() // ']'
			event, err := p.parseEvent()
			if err != nil {
				return nil, err
			}
			gcLog.MainEvent = *event
		} else if p.peek() == COMMA && p.dPeek() == TIME {
			p.next() // ','
			gcLog.Time = p.nextPair().Literal
			p.next() // 'secs'
		} else if p.peek() == TIMESTAMP {
			t, err := time.Parse(TIMESTAMP_LAYOUT, p.nextPair().Literal)
			if err != nil {
				return nil, err
			}
			gcLog.Timestamp = t
		} else {
			p.next()
		}

		// TODO: The Times section
	}

	return &gcLog, nil
}

func (p *Parser) parseEvent() (*GCEvent, error) {
	var gcEvent GCEvent
	// Events are always in this 6 Token format:
	// SIZE->SIZE(SIZE)
	for i := 0; i < 6; i++ {
		switch i {
		case 0:
			s, err := formatSize(p.nextPair().Literal)
			if err != nil {
				return nil, err
			}
			gcEvent.BeforeSize = s
		case 2:
			s, err := formatSize(p.nextPair().Literal)
			if err != nil {
				return nil, err
			}
			gcEvent.AfterSize = s
		case 4:
			s, err := formatSize(p.nextPair().Literal)
			if err != nil {
				return nil, err
			}
			gcEvent.TotalSize = s
		default:
			p.next() // Throw away other tokens
		}
	}

	return &gcEvent, nil
}

func formatSize(size string) (int, error) {
	s := size[:len(size)-1] // Pop off the K
	return strconv.Atoi(s)
}

func (p *Parser) next() Token {
	t := p.Tokens[p.Pos]
	p.Pos++
	return t
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Pos]
}

func (p *Parser) dPeek() Token {
	return p.Tokens[p.Pos+1]
}

func (p *Parser) eof() bool {
	return p.Pos == len(p.Tokens)-1
}

func (p *Parser) nextPair() TokenPair {
	t := p.TokenPairs[p.Pos]
	p.Pos++
	return t
}

func (p *Parser) peekPair() TokenPair {
	return p.TokenPairs[p.Pos]
}

func FetchLogs(file string) ([]GCLog, error) {
	var logs []GCLog

	bs, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	fs := string(bs)
	ss := strings.Split(fs, "\n")

	for _, l := range ss {
		parser := NewParser(l)
		log, err := parser.Parse()
		if err != nil {
			return nil, err
		}
		logs = append(logs, *log)
	}

	return logs, nil
}
