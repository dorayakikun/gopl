#! /bin/bash

set -e

go build .
TZ=US/Eastern ./clock2 -port 8010 &
TZ=ASIA/Eastern ./clock2 -port 8020 &
TZ=Europe/London ./clock2 -port 8030 &
clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030