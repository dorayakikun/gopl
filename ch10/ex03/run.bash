#!/usr/bin/env bash

set -e

go build gopl.io/ch1/fetch
./fetch "http://gopl.io/ch1/helloworld?go-get=1"