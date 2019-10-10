#!/bin/bash

set -e

go build gopl.io/ch3/mandelbrot
go build -o image main.go

./mandelbrot | ./image >mandelbrot.jpg
