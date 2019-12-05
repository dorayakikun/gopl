package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/get", db.get)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	fmt.Println("usage:")
	fmt.Println("http://localhost:8000/list")
	fmt.Println("http://localhost:8000/price?item=:item")
	fmt.Println("http://localhost:8000/create?item=:item&price=:price")
	fmt.Println("http://localhost:8000/get?item=:item")
	fmt.Println("http://localhost:8000/update?item=:item&price=:price")
	fmt.Println("http://localhost:8000/delete?item=:item")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "missing item: %q\n", item)
		return
	}

	if price, ok := db[item]; ok {
		newPrice := req.URL.Query().Get("price")
		if newPrice == "" {
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(w, "missing price: %q\n", newPrice)
		} else {
			p, err := strconv.ParseFloat(newPrice, 32)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest) // 400
				fmt.Fprintf(w, "not number price: %q\n", newPrice)
			} else {
				mutex := &sync.Mutex{}
				mutex.Lock()
				db[item] = dollars(p)
				mutex.Unlock()
				fmt.Fprintf(w, "update %q %s -> %s\n", item, price, dollars(p))
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "missing item: %q\n", item)
		return
	}
	if _, ok := db[item]; ok {
		mutex := &sync.Mutex{}
		mutex.Lock()
		delete(db, item)
		mutex.Unlock()
		fmt.Fprintf(w, "delete %q \n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "missing item: %q\n", item)
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%q already exists\n", item)
		return
	}
	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "missing price: %q\n", price)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "not number price: %q\n", price)
		return
	}
	mutex := &sync.Mutex{}
	mutex.Lock()
	db[item] = dollars(p)
	mutex.Unlock()
	fmt.Fprintf(w, "create %s: %s \n", item, dollars(p))
}

func (db database) get(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "missing item: %q\n", item)
		return
	}
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
