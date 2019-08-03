#!/bin/bash
set -e
go run surface.go > out.svg
open -a "Google Chrome" out.svg