package params

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	type query struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max" validate:"max=90"`
		Exact      bool     `http:"x"`
	}

	q := query{
		Labels:     []string{"どらやき", "ゆべし", "ようかん", "赤福"},
		MaxResults: 100,
		Exact:      true,
	}

	s, err := Pack(&q)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("url: %s", s)

	req := httptest.NewRequest(http.MethodGet, "https://example.com", nil)
	req.URL.RawQuery = s

	var q2 = query{}
	Unpack(req, &q2)

	if !reflect.DeepEqual(q, q2) {
		t.Fatal("unmatched")
	}
}

func TestPack2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	type query struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max" validate:"min=90"`
		Exact      bool     `http:"x"`
	}

	q := query{
		Labels:     []string{"どらやき", "ゆべし", "ようかん", "赤福"},
		MaxResults: 1,
		Exact:      true,
	}

	s, err := Pack(&q)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("url: %s", s)

	req := httptest.NewRequest(http.MethodGet, "https://example.com", nil)
	req.URL.RawQuery = s

	var q2 = query{}
	Unpack(req, &q2)

	if !reflect.DeepEqual(q, q2) {
		t.Fatal("unmatched")
	}
}