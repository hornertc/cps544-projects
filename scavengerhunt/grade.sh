#!/bin/sh

sha256sum -c checksum > grade.txt
points=$(grep -c OK grade.txt)
points=$((points * 2)) # 2 points each
echo "Points: $points" >> grade.txt
cat grade.txt
