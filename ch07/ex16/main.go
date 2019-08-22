package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"z/eval"
)

func main() {

	http.HandleFunc("/", calculate)

	fmt.Println("usage:")
	fmt.Println(fmt.Sprintf("http://localhost:8000/?expr=%s&x=%d", url.QueryEscape("1+x"), 2))
	fmt.Println(fmt.Sprintf("http://localhost:8000/?expr=%s&n=%d", url.QueryEscape("min(1, 2, 3) + sqrt(n)"), 16))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calculate(w http.ResponseWriter, req *http.Request) {
	expr := req.URL.Query().Get("expr")
	if expr == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "missing query string \"expr\"")
		return
	}

	e, err := eval.Parse(expr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, err)
		return
	}

	vars := map[eval.Var]bool{}
	err = e.Check(vars)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, err)
		return
	}

	env := make(map[eval.Var]float64)
	for v := range vars {
		s := req.URL.Query().Get(string(v))
		if s == "" {
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(w, "missing query string %s\n", v)
			return
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(w, "%q is no number: %s\n", v, s)
			return
		}

		env[v] = f
	}

	fmt.Fprintf(w, "result: %g\n", e.Eval(env))
}