package sqlparser

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"unicode"
)

type tokenType int

const (
	tokenNumber      tokenType = iota
	tokenString                // string
	tokenIdent                 // identifier
	tokenPeriod                // period symbol .
	tokenEquals                // equals symbol =
	tokenGreaterThan           // greater than symbol >
	tokenLessThan              // less than symbol <
	tokenPlus                  // plus symbol +
	tokenMinus                 // minus symbol
	tokenAsterisk              // multiplication symbol *
	tokenSlash                 // division symbol /
	tokenCaret                 // exponentiation symbol ^
	tokenPercent               // The modulo symbol %
	tokenExclamation           // The factorial or not symbol !
	tokenQuestion              // query parameter marker ?
	tokenOpenParen             // opening parenthesis (
	tokenCloseParen            // closing parenthesis )
	tokenComma                 // comma ,
	tokenNone                  // None
)

var tokens = []string{
	tokenNumber:      "tokenNumber",
	tokenString:      "tokenString",
	tokenIdent:       "tokenIdent",
	tokenPeriod:      ".",
	tokenEquals:      "=",
	tokenGreaterThan: ">",
	tokenLessThan:    "<",
	tokenPlus:        "+",
	tokenMinus:       "-",
	tokenAsterisk:    "*",
	tokenSlash:       "/",
	tokenCaret:       "^",
	tokenPercent:     "%",
	tokenExclamation: "!",
	tokenQuestion:    "?",
	tokenOpenParen:   "(",
	tokenCloseParen:  ")",
	tokenComma:       ",",
	tokenNone:        "",
}

func (tokType tokenType) String() string {
	s := ""
	if 0 <= tokType && tokType < tokenType(len(tokens)) {
		s = tokens[tokType]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tokType)) + ")"
	}

	return s
}

var eof = rune(0)

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isAlphabetic(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isAlphaNumeric(ch rune) bool {
	return isAlphabetic(ch) || isDigit(ch)
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}

func (s *Scanner) peek() rune {
	ch, err := s.r.Peek(1)
	if err != nil {
		return eof
	}

	return rune(ch[0])
}

func (s *Scanner) Scan() (tok tokenType, lit string) {
	s.consumeWhiteSpace()

	ch := s.peek()

	if isDigit(ch) {
		return s.scanNumber()
	}

	if isAlphabetic(ch) {
		return s.scanIDent()
	}

	switch ch {
	case '\'':
		return s.scanString()
	case eof:
		return 10, ""
	}

	return s.scanSymbol()
}

func (s *Scanner) consumeWhiteSpace() {
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhiteSpace(ch) {
			s.unread()
			break
		}
	}
}

func (s *Scanner) scanString() (tok tokenType, lit string) {
	var buf bytes.Buffer
	s.read()

	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch != '\'' {
			buf.WriteRune(ch)
		} else {
			break
		}

	}

	return tokenString, buf.String()
}

func (s *Scanner) scanNumber() (tok tokenType, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isDigit(ch) {
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	return tokenNumber, buf.String()
}

//
func (s *Scanner) scanSymbol() (tok tokenType, lit string) {
	ch := string(s.read())
	switch ch {
	case tokenPeriod.String():
		return tokenPeriod, ch
	case tokenEquals.String():
		return tokenEquals, ch
	case tokenGreaterThan.String():
		return tokenGreaterThan, ch
	case tokenLessThan.String():
		return tokenLessThan, ch
	case tokenMinus.String():
		return tokenMinus, ch
	case tokenMinus.String():
		return tokenAsterisk, ch
	case tokenSlash.String():
		return tokenSlash, ch
	case tokenPlus.String():
		return tokenPlus, ch
	case tokenCaret.String():
		return tokenCaret, ch
	case tokenPercent.String():
		return tokenPercent, ch
	case tokenExclamation.String():
		return tokenExclamation, ch
	case tokenQuestion.String():
		return tokenQuestion, ch
	case tokenOpenParen.String():
		return tokenOpenParen, ch
	case tokenCloseParen.String():
		return tokenCloseParen, ch
	case tokenComma.String():
		return tokenComma, ch
	}

	return tokenNone, ""
}

func (s *Scanner) scanIDent() (tok tokenType, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isAlphaNumeric(ch) || ch == '_' {
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	return tokenIdent, buf.String()
}
