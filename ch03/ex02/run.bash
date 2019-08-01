#!/bin/bash
export GO111MODULE=on
go run surface.go > out0.svg
go run surface.go -mode 1 > out1.svg
go run surface.go -mode 2 > out2.svg
go run surface.go -mode 3 > out3.svg