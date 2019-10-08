#!/bin/bash

set -e

go run findlinks.go https://golang.org
cd golang.org
python -m SimpleHTTPServer 3000

