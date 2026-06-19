#!/bin/bash
set -eo pipefail

echo Testing table
go run ./cmd/windchill -table | grep 30.6 | grep 13.0
