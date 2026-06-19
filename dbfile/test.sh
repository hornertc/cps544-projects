#!/bin/bash

set -euo pipefail

export DBFILE_DATABASE=$PWD/testdb
rm -rf "$DBFILE_DATABASE"

rm -rf results
mkdir -p results

echo "==== no sudo"
grep "sudo " src/dbfile && exit 100

echo "==== no su"
grep "su " src/dbfile && exit 101

echo "==== create tables"
src/dbfile create people "First,Last,Address,Social Security Number,Age"
test -s "$DBFILE_DATABASE/people.csv"

src/dbfile create dog "Breed,Cost,Family"
test -s "$DBFILE_DATABASE/dog.csv"

NAME=list
src/dbfile list > results/$NAME
cat > results/$NAME.expected <<EOF
dog
people
EOF
diff results/$NAME results/$NAME.expected

src/dbfile create dog "Breed,Family" && exit 100

echo "==== add records"
src/dbfile add people "John,Smith,300 E. St,123-45-6789,21"
src/dbfile add people "Jan,Smith,300 E. St,123-45-6790,20"
src/dbfile add people "Joe,Biden,Whitehouse,321-45-6789,80"
test -s "$DBFILE_DATABASE/people.csv"

echo "==== Lookup Joe"
NAME=lookup_joe
src/dbfile lookup people First=Joe > results/$NAME
cat > results/$NAME.expected <<EOF
Joe,Biden,Whitehouse,321-45-6789,80
EOF
diff results/$NAME results/$NAME.expected

echo "==== Lookup Smith"
NAME=lookup_smith
src/dbfile lookup people Last=Smith | sort > results/$NAME
cat > results/$NAME.expected <<EOF
Jan,Smith,300 E. St,123-45-6790,20
John,Smith,300 E. St,123-45-6789,21
EOF
diff results/$NAME results/$NAME.expected

echo "==== Lookup Address"
NAME=lookup_address
src/dbfile lookup people Address="300 E. St" | sort > results/$NAME
cat > results/$NAME.expected <<EOF
Jan,Smith,300 E. St,123-45-6790,20
John,Smith,300 E. St,123-45-6789,21
EOF
diff results/$NAME results/$NAME.expected

echo "==== Delete Smith"
src/dbfile delete people Last=Smith
test -s "$DBFILE_DATABASE/people.csv"

echo "==== Lookup Biden (post delete)"
NAME="lookup_biden_post_delete"
src/dbfile lookup people Last="Biden" > results/$NAME
cat > results/$NAME.expected <<EOF
Joe,Biden,Whitehouse,321-45-6789,80
EOF
diff results/$NAME results/$NAME.expected

echo "==== Add Josephine"
src/dbfile add people "Josephine,Smith,300 E. St,999-45-6790,2"

echo "==== Update Biden"
src/dbfile update people Address="Whitehouse" "Joe,Biden,Whitehouse,321-45-6789,81"

echo "==== Lookup Biden (post update)"
NAME=lookup_biden_post_update
src/dbfile lookup people Last="Biden" > results/$NAME
cat > results/$NAME.expected <<EOF
Joe,Biden,Whitehouse,321-45-6789,81
EOF
diff results/$NAME results/$NAME.expected

echo "==== error handling"
src/dbfile add food Sandwitch,Good && exit 100

src/dbfile addd foo Sandwitch,Good && exit 100

echo "==== All tests pass!"
