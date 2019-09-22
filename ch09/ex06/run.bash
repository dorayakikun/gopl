#!/bin/bash

set -e

GOMAXPROCS=1 go test -bench .
GOMAXPROCS=2 go test -bench .
GOMAXPROCS=3 go test -bench .
GOMAXPROCS=4 go test -bench .
GOMAXPROCS=5 go test -bench .
GOMAXPROCS=6 go test -bench .