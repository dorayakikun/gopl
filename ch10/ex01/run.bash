#!/bin/bash

set -e

go build gopl.io/ch3/mandelbrot
go build -o imageconv main.go

./mandelbrot | ./imageconv >mandelbrot.jpg
