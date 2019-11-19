// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
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

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	//data, err = MarshalIndent(strangelove)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("MarshalIdent() = %s\n", data)

	// float
	data, err = Marshal(123.45)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshal() = %s\n", data)

	// complex
	data, err = Marshal(complex(1, 2))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshal() = %s\n", data)

	// bool
	data, err = Marshal(true)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshal() = %s\n", data)

	data, err = Marshal(false)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshal() = %s\n", data)

	// chan
	c := make(chan map[string]string)
	data, err = Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshal() = %s\n", data)

	// chan
	f := func(a int, b int, message string) string { return "" }
	data, err = Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Marshal() = %s\n", data)
}
