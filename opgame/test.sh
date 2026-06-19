#!/bin/bash

set -euo pipefail

mkdir -p results

echo ==== Test single
mkdir -p results/single
./opgame < testdata/single/basic_in.txt > results/single/basic_out.txt
diff results/single/basic_out.txt testdata/single/basic_out.txt

echo ==== Test multiple
mkdir -p results/all
./opgame -all < testdata/all/basic_in.txt > results/all/basic_out.txt
diff results/all/basic_out.txt testdata/all/basic_out.txt

echo ==== Test failure nan
mkdir -p results/failure
./opgame < testdata/failure/nan_in.txt > results/failure/nan_out.txt 2> results/failure/nan_err.txt && exit 100
cat results/failure/nan_err.txt
echo Ensure that stderr is non-empty
test -s results/failure/nan_err.txt

echo ==== Test failure short
mkdir -p results/failure
./opgame < testdata/failure/short_in.txt > results/failure/short_out.txt 2> results/failure/short_err.txt && exit 100
cat results/failure/short_err.txt
echo Ensure that stderr is non-empty
test -s results/failure/short_err.txt
