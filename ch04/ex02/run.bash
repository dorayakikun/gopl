#!/bin/bash
export GO111MODULE=on
echo "x" | go run main.go
echo "x" | go run main.go -mode 1
echo "x" | go run main.go -mode 2