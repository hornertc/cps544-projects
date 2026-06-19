#!/bin/sh

# set -x

echo "$0 is finding programs containing \"$1\""
ps aux | grep "$1"

echo "$*"
echo "$@"
