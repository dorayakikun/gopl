// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"reflect"
	"testing"
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