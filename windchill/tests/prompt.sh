#!/bin/bash
set -eo pipefail

echo "Testing prompt, 13 F, 10 mph, is 0.18 F"
cat << EOF | go run ./cmd/windchill | grep -F '0.18'
13
10
EOF

echo "Testing prompt, 15 F, 30 mph, is 0.18 F"
cat << EOF | go run ./cmd/windchill | grep -F -e '-5.49'
15
30
EOF
