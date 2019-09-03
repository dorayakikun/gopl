#!/bin/bash

set -e

echo "go run main.go"
echo "7 8"
echo "2 2"
echo "4 5"
echo "########"
echo "#......#"
echo "#.######"
echo "#..#...#"
echo "#..##..#"
echo "##.....#"
echo "########"

go run main.go <<EOS
7 8
2 2
4 5
########
#......#
#.######
#..#...#
#..##..#
##.....#
########
EOS