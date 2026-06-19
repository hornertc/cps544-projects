#!/bin/bash
set -euo pipefail

echo "Test 1 (stdin)"
go run ./cmd/wordsearch de - < puzzle1.txt > test.out
diff test.out tests/out1.txt
echo PASS
