#!/bin/bash
go run newton.go -mode 0 > out.png
go run newton.go -mode 1 > out1.png
go run newton.go -mode 2 > out2.png
go run newton.go r-mode 3 > out3.png