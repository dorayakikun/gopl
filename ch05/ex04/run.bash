#!/bin/bash
set -e
curl https://golang.org | go run findlinks.go -
