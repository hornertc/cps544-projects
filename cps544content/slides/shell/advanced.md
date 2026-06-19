# Advanced Shell

---

## Clean Shell

```sh
env -i HOME=$(mktemp -d) bash --noprofile --norc
cd
pwd
printenv
```

---

## Here String (Bash-Only)

- `<<<` can be used to feed a string into STDIN of a command

```bash
program <<< "Some string $myvar"

# equivalent to the "here-document" approach
program <<DELIM
Some string $myvar
DELIM
```

Notes:

- `wc <<<"hello world"`

---

## Math

- `$((x+45))` evaluates to the values of `x` plus 45
- only integers supported
- addition, subtraction, multiplication, division

```bash
x=45
test $((x-20)) -lt 10
echo $?
```

outputs 1 (false)

----

## More Math Options

- `expr` can be used to slightly more flexible math
- best to use `python` or similar for more complex math
  - `python -c 'import math; print(math.sin(0.3))'`

Notes:

- if using `uv` then it is `uv run python ...`

---

## Redirection

- `program` is equivalent to `program >&1`
  - write to file descriptor number 1 (STDOUT)
- `program >filename` is equivalent to `program 1>filename`
  - 1 is STDOUT and is the default
- `program 2>filename` writes STDERR to filename

Notes:

- `curl -v https://example.com 2>log.txt` writes STDERR to `log.txt`
  - STDOUT goes to the terminal

----

## Redirect STDERR to STDOUT

- `curl -o data.txt httpss://example.com >log.txt`
  - `log.txt` is empty
- `curl -o data.txt httpss://example.com >log.txt 2>&1`
  - `log.txt` now has the error
- `2>&1` redirects STDERR(2) to STDOUT(1)

----

## Beyond 0, 1, and 2

```bash
cat - > morefiles.sh <<EOF
#!/bin/sh
echo "First line"
echo "Second" >&2
echo "Third" >&3
echo "Forth" >&4
EOF
chmod +x morefiles.sh
./morefiles.sh 1>file1 2>file2 3>file3 4>&1
```

Notes:

- `3>file3` sets fd 3 to write to `file3` (closes after command finishes)
- POSIX docs are [here](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_07_02)

----

## Command Output/Input as File Argument (Bash-Only)

- `<(program)` returns a pipe file that when read streams STDOUT of `program`
- `>(program)` returns a pipe file that when written to streams to STDIN of `program`

```bash
wc morefiles.sh <(echo "Hello world") <(echo -n "What")
curl -o >(wc) https://example.com
curl https://example.com | wc
```

---

## Functions

```sh
inc_A() {
  # Increment A by 1
  A=$((A + 1)) 
}
A=1 
while [ "$A" -le 10 ] 
do 
  echo $A
  inc_A
done
```

----

## Function Local Variables

- `local myvar=somvalue` (`local myvar`) is a local variable with local scope
  - function body only

```sh
myfunc() {
  local x=45
  echo x is $x
  x=1
  readonly x
}
myfunc
echo $x
readonly
```

----

## Shifting Arguments

- shifts the arguments to the left by one
  - `1=$2`
  - `2=$3`, ...
- either in a script or function
- useful to shift argument `$1` out of existence after being processed
- often process arguments in a loop
  - always process the `$1` argument (no need for an index variable)

```bash
#handle $1
shift
# handle the next argument, which is $1
shift
# handle the next argument, which is still $1
```

----

## Type of Symbols

- `type` tells you the type of the symbol name passed as an argument

```bash
type rsync     # rsync is /usr/bin/rsync
type echo      # echo is a shell builtin
type now       # type: now: not found
alias now=date # create an alias
type now       # now is aliased to `date'
type grep # grep is aliased to `grep --color=auto'
```

---

## Sourcing Files

`library.sh`

```sh
somevar=something
```

```sh
source library.sh
echo $somevar
```

----

## Check if a Program is Available

- `command` similar to `which` but builtin
- `command -v some-program-name`
  - check exit code with `$?`
- `command -v some-program-name >/dev/null 2>&1 || echo 'We need "some-program-name" but it is not available'`
- Better than calling `which`
  - `command is builtin not an external program , thus faster
  - some `which` implementations do not properly set the exit code

---

## Performance

- Each command forks a new process (expensive)
- Avoid forking by letting the shell do mundane tasks (string manipulation)
- Prefer `chmod +x $*` to

```sh
for i in $*
do
  chmod +x $i
done
```

- Also prefer builtin commands instead of external commands

---

## Pipefail (Bash-Only)

```bash
cat missing-file | grep -c bolts
echo $? # returns 0 (but expected 1)
```

```bash
set -o pipefail
cat missing-file | grep -c bolts
echo $? # returns 1 as expected
```

- `bash` scripts often use `set -euo pipefail`

## Bash Completion (Bash-Only)

- Enabled when tab is pressed
- Each shell uses different shell completion implementations
- Arguments to a program are only known to the program
  - But shell completion allows for interactive command line completion

```bash
# at installation time
kubectl bash completion > ~/.bash_completion.d/kubectl
helm bash completion > ~/.bash_completion.d/helm
# in .profile - sourced at shell creation
for COMPLETION in ~/.bash_completion.d/*
do
  [[ -r "${COMPLETION}" ]] && source "${COMPLETION}"
done
```

Notes:

- Demonstrate shell completion with `kubectl` and `rsync`

---

## Future Items (time permitting)

TODO

---

## Previous Commands (Bash-Only)

- `history` or `history | grep curl`
  - `history` is a builtin function
- `!1452`
- `!!` - run last command
  - same as `!-1`

---

## Traps

e.g., create a temporary directory

---

## SetUid

---

## SetGid

---

## Sudoers

- Show how you can become users other than root

---

## Concurrency in Scripts

Show & and `wait` and maybe `trap`.
