package params

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	type query struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
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
