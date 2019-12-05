package poster

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

func FetchMovie(apikey string, keyword string) (*Movie, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&t=%s", apikey, url.QueryEscape(keyword)))

	defer resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "send request failed")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("missing movie")
	}

	var m Movie
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, errors.Wrap(err, "decode failed")
	}

	return &m, nil
}

func FetchPosterImage(url string, filename string) error {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return errors.Wrap(err, "create request failed(fetch poster image)")
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("missing poster")
	}

	f, err := os.Create(fmt.Sprintf("%s.png", filename))
	defer f.Close()
	if err != nil {
		return errors.Wrap(err, "create image file failed")
	}

	io.Copy(f, resp.Body)

	fmt.Printf("downloaded %q.png\n", filename)

	return nil
}
