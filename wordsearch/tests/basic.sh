#!/bin/bash
set -euo pipefail

echo "Test 1"
go run ./cmd/wordsearch de puzzle1.txt > test.out
diff test.out tests/out1.txt
echo PASS

echo "Test 2"
go run ./cmd/wordsearch œke puzzle1.txt > test.out
diff test.out tests/out2.txt
echo PASS

echo "Test 3"
go run ./cmd/wordsearch cx puzzle2.txt > test.out
diff test.out tests/out3.txt
echo PASS

echo "Test 4"
go run ./cmd/wordsearch i puzzle2.txt > test.out
diff test.out tests/out4.txt
echo PASS

echo "Test 5"
go run ./cmd/wordsearch not-there puzzle1.txt > test.out
diff test.out tests/out5.txt
echo PASS