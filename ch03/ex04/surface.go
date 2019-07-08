package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var (
	width, height = 600., 320.
	cells         = 100
	xyrange       = 30.0
	xyscale       = float64(width) / 2 / float64(xyrange)
	zscale        = float64(height) * 0.4
	angle         = math.Pi / 6
	color         = "gray"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryWidth := r.URL.Query().Get("width")
		if queryWidth != "" {
			newWidth, err := strconv.ParseFloat(queryWidth, 10)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bad Request"))
			}
			width = newWidth
			xyscale = float64(width) / 2 / float64(xyrange)
		}

		queryHeight := r.URL.Query().Get("height")
		if queryHeight != "" {
			newHeight, err := strconv.ParseFloat(queryHeight, 10)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bad Request"))
			}
			height = newHeight
			zscale = float64(height) * 0.4
		}

		queryColor := r.URL.Query().Get("color")
		if queryColor != "" {
			color = queryColor
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(w http.ResponseWriter) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: "+color+"; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			f := func() {
				for _, point := range []float64{ax, ay, bx, by, cx, cy, dx, dy} {
					if math.IsInf(point, 0) {
						return
					}
				}
				builder.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g, %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy))
			}
			f()
		}
	}
	builder.WriteString("</svg>")
	_, err := fmt.Fprint(w, builder.String())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
}

func corner(i, j int) (sx float64, sy float64) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	z := f(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
