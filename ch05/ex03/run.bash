#!/bin/bash
set -e
curl https://golang.org | go run findtexts.go -
