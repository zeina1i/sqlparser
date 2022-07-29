package sqlparser

import (
	"strings"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		tok tokenType
		lit string
	}{
		{tok: tokenIdent, lit: "SELECT"},
		{tok: tokenIdent, lit: "artist"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "name"},
		{tok: tokenComma, lit: tokenComma.String()},
		{tok: tokenIdent, lit: "album"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "name"},
		{tok: tokenComma, lit: tokenComma.String()},
		{tok: tokenIdent, lit: "EXTRACT"},
		{tok: tokenOpenParen, lit: tokenOpenParen.String()},
		{tok: tokenIdent, lit: "YEAR"},
		{tok: tokenIdent, lit: "FROM"},
		{tok: tokenIdent, lit: "NOW"},
		{tok: tokenOpenParen, lit: tokenOpenParen.String()},
		{tok: tokenCloseParen, lit: tokenCloseParen.String()},
		{tok: tokenCloseParen, lit: tokenCloseParen.String()},
		{tok: tokenMinus, lit: tokenMinus.String()},
		{tok: tokenIdent, lit: "album"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "release_year"},
		{tok: tokenIdent, lit: "AS"},
		{tok: tokenIdent, lit: "age"},
		{tok: tokenIdent, lit: "FROM"},
		{tok: tokenIdent, lit: "artist"},
		{tok: tokenIdent, lit: "INNER"},
		{tok: tokenIdent, lit: "JOIN"},
		{tok: tokenIdent, lit: "album"},
		{tok: tokenIdent, lit: "ON"},
		{tok: tokenIdent, lit: "album"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "artist_id"},
		{tok: tokenEquals, lit: tokenEquals.String()},
		{tok: tokenIdent, lit: "artist"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "id"},
		{tok: tokenIdent, lit: "WHERE"},
		{tok: tokenIdent, lit: "album"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "genre"},
		{tok: tokenExclamation, lit: tokenExclamation.String()},
		{tok: tokenEquals, lit: tokenEquals.String()},
		{tok: tokenString, lit: "country"},
		{tok: tokenIdent, lit: "AND"},
		{tok: tokenIdent, lit: "album"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "release_year"},
		{tok: tokenGreaterThan, lit: tokenGreaterThan.String()},
		{tok: tokenEquals, lit: tokenEquals.String()},
		{tok: tokenNumber, lit: "1980"},
		{tok: tokenIdent, lit: "ORDER"},
		{tok: tokenIdent, lit: "BY"},
		{tok: tokenIdent, lit: "artist"},
		{tok: tokenPeriod, lit: tokenPeriod.String()},
		{tok: tokenIdent, lit: "name"},
		{tok: tokenIdent, lit: "ASC"},
		{tok: tokenComma, lit: tokenComma.String()},
		{tok: tokenIdent, lit: "age"},
		{tok: tokenIdent, lit: "DESC"},
	}

	query := `
			SELECT artist.name, album.name, EXTRACT(YEAR FROM NOW()) - album.release_year AS age
            FROM artist INNER JOIN album ON album.artist_id = artist.id
            WHERE album.genre != 'country' AND album.release_year >= 1980
            ORDER BY artist.name ASC, age DESC`

	s := NewScanner(strings.NewReader(query))
	for i, tt := range tests {
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. token mismatch: exp=%q got=%q <%q>", i, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. literal mismatch: exp=%q got=%q", i, tt.lit, lit)
		}
	}
}
