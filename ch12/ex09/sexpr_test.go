// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"text/scanner"
)

func TestEncoder_Encode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	err := enc.Encode(strangelove)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	var m Movie
	err = Unmarshal(bytes.NewReader(buf.Bytes()), &m)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(strangelove, m) {
		t.Fatal("restore data failed")
	}
}

func TestDecoder_Decode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	err := enc.Encode(strangelove)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	dec := NewDecoder(bytes.NewReader(buf.Bytes()))
	var m Movie
	err = dec.Decode(&m)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if !reflect.DeepEqual(strangelove, m) {
		t.Fatal("restore data failed")
	}
}

func TestDecoder2_Token(t *testing.T) {
	in := `((Title "") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 0) )`
	tokens := []Token{
		StartList{},
		StartList{},
		Symbol{"Title"},
		String{},
		EndList{},
		StartList{},
		Symbol{"Subtitle"},
		String{"How I Learned to Stop Worrying and Love the Bomb"},
		EndList{},
		StartList{},
		Symbol{"Year"},
		Int{0},
		EndList{},
		EndList{},
	}

	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader([]byte(in)))
	lex.next()
	d := &Decoder2{ lex }

	for _, want := range tokens {
		token, err := d.Token()
		if err != nil {
			t.Errorf("%v\n", err)
		}

		if !reflect.DeepEqual(token, want) {
			t.Errorf("actual: %v want: %v\n", token, want)
		}
	}

	_, err := d.Token()

	if err != io.EOF {
		t.Fatalf("unexpected error: %v", err)
	}
}