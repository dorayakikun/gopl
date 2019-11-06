#!/usr/bin/env bash

set -e

go test -bench=.

cd mapintset
go test -bench=.