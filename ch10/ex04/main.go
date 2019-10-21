package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type result struct {
	Deps []string `json:"Deps"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("missing argument")
	}

	out, err := exec.Command("go", "list", "-json", os.Args[1]).Output()
	if err != nil {
		log.Fatal(err)
	}

	r := &result{}
	err = json.Unmarshal(out, r)
	if err != nil {
		log.Fatal(err)
	}

	deps := make(map[string]bool)
	for _, d := range r.Deps {
		out2, err := exec.Command("go", "list", "-json", d).Output()
		if err != nil {
			log.Fatal(err)
		}
		r2 := &result{}
		err = json.Unmarshal(out2, r2)
		if err != nil {
			log.Fatal(err)
		}

		for _, d2 := range r2.Deps {
			deps[d2] = true
		}
	}
	for d := range deps {
		fmt.Println(d)
	}
}
