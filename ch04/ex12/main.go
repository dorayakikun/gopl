package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Comic struct {
	Transcript string `json:"Transcript"`
	Img        string `json:"Img"`
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: go run main.go <KEYWORD>\n")
		os.Exit(2)
	}

	keyword := os.Args[1]

	num := 1
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("get wd failed %s", err.Error())
		os.Exit(1)
	}

	_, err = os.Stat(fmt.Sprintf("%s/indexes", wd))
	if err != nil {
		if err := os.Mkdir(fmt.Sprintf("%s/indexes", wd), 0777); err != nil {
			fmt.Printf("create dir failed %s", err.Error())
			os.Exit(1)
		}
	}

	comics := make(map[int]Comic)

	fmt.Println("create indexes...")
	for {
		// TODO 以下をgo routine化したい
		// 404は `status code` が404になるので...
		for num == 404 {
			num++
			continue;
		}

		fname := fmt.Sprintf("%s/indexes/%d", wd, num)

		_, err := os.Stat(fname)
		// すでに取得済みならskip
		if err == nil {
			buf, err := ioutil.ReadFile(fname)
			if err != nil {
				fmt.Printf("read file failed %s\n", err.Error())
				os.Exit(1)
			}
			c := Comic{}
			if err := json.Unmarshal(buf, &c); err != nil {
				fmt.Printf("unmarshal failed %s\n", err.Error())
				os.Exit(1)
			}
			comics[num] = c
			num++
			continue;
		}

		resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", num))
		if err != nil {
			fmt.Printf("fetch failed %s", err.Error())
			os.Exit(1)
		}

		// 最新刊まで到達したらbreak
		if resp.StatusCode == 404 {
			break
		}

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("read failed %s", err.Error())
			os.Exit(1)
		}
		if err = ioutil.WriteFile(fname, buf, 0644); err != nil {
			fmt.Printf("create file failed %s\n", err.Error())
			os.Exit(1)
		}

		c := Comic{}
		if err := json.Unmarshal(buf, &c); err != nil {
			fmt.Printf("unmarshal failed %s\n", err.Error())
			os.Exit(1)
		}
		comics[num] = c
		num++
	}
	fmt.Println("created indexes")

	fmt.Println("search comics...")
	fmt.Println()

	var keys []int
	for k := range comics {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k := range keys {
		c := comics[k]
		if strings.Contains(c.Transcript, keyword) {
			fmt.Printf("img:\n%s\n\ntranscript:\n%s\n", c.Img, c.Transcript)
			os.Exit(0)
		}
	}
	fmt.Println("comic not found")
}