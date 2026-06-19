# More Shell

---

## Commands

- `date; echo "done"` outputs the date followed by "done"
  - Run sequentially
- Equivalent to

```sh
date
echo "end"
```

---

## Comments

- if shell word begins with  `#` the rest of the line is ignored

```sh
$ echo hello # there
hello
$ echo hello#there
hello#there
```

- second case there are only two shell words
  - `hello`
  - `hello#there`
- neither begins with a `#`

---

## Variables

- Can create a new variable `myvar=somevalue`
- No spaces around `=` sign
- RHS must be a word or needs to be quoted
  - `myvar="some value"`
- By convention, special variable names are all capital with underscores
- `now="The date is $(date)"`
- `readonly myvar=somevalue` (`readonly myvar`) declares a constant
  - `readonly` lists all readonly variables

Notes:

```sh
readonly x=hello
readonly
x=world # error
```

---

## Standard Streams

- Standard input (file descriptor 0)
  - regular program input, STDIN
- Standard output (file descriptor 1)
  - regular program output, STDOUT
- Standard error (file descriptor 2)
  - abnormal/error program output, STDERR

Notes:

- You must memorize the integers

----

## File Descriptors

- Integer assigned by the Kernel:
  - Standard streams
  - Files
  - Network sockets
- FD used in `open()`, `write()`, `close()` (see `man 2 write`)

----

## Redirection

- `date` outputs to standard output (stdout)
- `date > date.txt` writes the date to file `date.txt`
- `grep dog < animals` outputs only lines matching "dog" in the animals file
  - `grep dog animals` is equivalent
- `date >> starting-dates.txt` will append current date to the file `starting-dates.txt`

----

## Here Document

```sh
x=dog
cat - > myfile <<EOF
Favorite animal
is $x
EOF
```

- `EOF` can be **any** string that does not appear in file contents
- `EOF` must appear on a line by itself
- Replacement occurs as normal
- Feature of shell not `cat`

Notes:

- Show replacing `cat -` with `grep dog` ...

---

## Pipelines

- `who | wc -c` - count the characters output to STDOUT by who
- `date; who | wc`
  - Only `who` is piped through `wc` (not `date`)
  - Forms a single command
  - Runs after `date`
  - Precedence of `|` is higher than `;`

----

- `(date; who) | wc`
  - Output of `date` and `who` are concatenated
  - Precedence of `()` is higher than `|`

----

## Tee

- Data flow through a pipe can be tapped (saved to a file)
- `(date; who) | tee save.txt | wc`
  - `cat save.txt`
  - `wc` receives the data as if `tee` was not in the pipeline
  - `tee` only copies STDIN to file

---

## Long Running Commands

- Long running commands block your terminal
  - Unable to use the terminal for other tasks
- Option 1:
  - `sleep 100`
  - Ctrl-Z will send the SIGTSTP signal
  - `bg` backgrounds stopped command (starts it again but in the background)
  - `fg` bring to foreground (block on it)
- Option 2:
  - `sleep 100 &` runs the command in the background from the start

----

- `(sleep 5; date) & date`
  - Parentheses are needed (precedence of `&` is higher than `;`)
- `(sleep 300; echo Tea is ready) &` reminder mechanism

----

- `pr file | lpr &`
  - Same as `(pr file | lpr) &`
  - Pipelines are commands and `&` applies to commands
- `jobs` list background jobs

----

## Job Control Signals

- `SIGSTOP` stops a process
- `SIGTSTP` interactive stop a process
- `SIGCONT` continue a process
- `SIGCHLD` signals that a child has terminated or stopped

----

## Killing Rogue Processes (Progression)

1. `kill -s SIGINT 1234` kindly request process termination (same as Ctrl-C)

2. `kill -s SIGTERM 1234` strongly request process termination.  Can be "caught" by the receiving process (and ignored if desired)

3. `kill -s SIGKILL 1234` or `kill -9 1234` terminates process.  Cannot be caught!

Notes:

- Processes with tight controls on shutdown procedures will listen to signals (e.g., servers, databases)
- Some processes will gracefully shutdown on the first `SIGINT`
  - Ungracefully shutdown on second `SIGINT`

----

## Signal Handling

- `kill -SIGUSR1 1234` send user defined signal `SIGUSR1` to process 1234
- Dozens of signals in UNIX
- Used for errors, window resizing, job control, and many other notifications
- `man 7 signal` and [GNU libc](https://www.gnu.org/software/libc/manual/html_node/Standard-Signals.html)

Notes:

- Run `go run examples/signals/main.go`
- `kill -SIGUSR1 1234`
- Resize window

---

## Command Arguments

- Arguments are words
  - Separated by spaces and tabs
  - Presented as an array of strings to the program
- Programs often parse arguments beginning with `-` as options
- Special characters (e.g., `>`, `<`, `|`, `;`, `&`) are not arguments
  - Control how the shell handles the command
- `-` by itself often means Stdin

----

- `echo Hello >junk`
  - `>junk` is not an argument to `echo`
  - STDOUT of `echo` is written to `junk` by the shell
  - String `junk` is never seen by program `echo`
  - Need not be the last word
    - `>junk echo Hello` is equivalent

---

## Exit Codes

- Exit code 0 indicates success
- Non-zero exit code indicates failure
  - Value indicates type of error
  - STDERR often contains the error message(s)
- Scripts condition on exit code (not stderr) for error handling
  - Want distinct error code for each error class

---

## Metacharacters

- Redirection
  - `>` ⟶ Output redirection
  - `>>` ⟶ Output redirection (append)
  - `<` ⟶ Input redirection
  - `<<` ⟶Input redirection, "here document"

- File substitution
  - `*` ⟶ zero or more characters
  - `?` ⟶ one character
  - `[ ]` ⟶ any character between brackets

----

- Command execution
  - ``` `cmd` ``` ⟶ Command substitution (legacy)
  - `$(cmd)`⟶ Command substitution
  - `$p_1$ | $p_2$` ⟶ Pipe, Stdout of $p_1$ to Stdin of $p_2$
  - `$p_1$; $p_2$` ⟶ Command sequence, run $p_1$ then $p_2$
  - `$p_1$ || $p_2$` ⟶ OR, run $p_1$, if unsuccessful run $p_2$
  - `$p_1$ && $p_2$` ⟶ AND, run $p_1$ if successful run $p_2$

----

- `(cmd)` ⟶ Run commands in sub-shell
- `{cmd}` ⟶ Run commands in current shell, (rarely used)
- `&`   ⟶ Run command in the background
- `#`   ⟶ Comment
- `=` ⟶ assign value on RHS to variable on LHS

Notes:

- similarity between `(cmd)` and `$(cmd)`

----

- Quotes and Escaping
  - `' '` ⟶ take literally
  - `" "` ⟶ take literally after `$`, ``` ` ` ```, `\` interpreted
  - `$`   ⟶ Expand the value of a variable
  - `\`   ⟶ Prevent or escape interpretation of the next character

---

## Matching

- Shell recognizes `*` as special
- Tell shell to search directory for filenames in which any string of characters occurs in the position of the `*`
  - `echo *` is poor replacement for `ls`
- known as globbing
- also `ls file?.txt` and `ls file.[ch]`

---

## Quotes

- Prevent globbing of `echo X*Y`

```sh
echo "X*Y" # double quotes
echo 'X*Y' # single quotes
echo X\*Y  # escape \*
echo X'*'Y # quote part of string
```
  
- Prevent splitting on spaces (shell words)

```sh
echo  X   Y  # X Y
echo "X   Y" # X   Y
echo X\ \ \ Y # X   Y
```

- shell expands `$`, ` `` `, `\` inside double quotes

----

## Single Quotes

- Taken literally `echo 'Foo *\n Bar'`
- Does not allow any substitution

```sh
echo 'Hello $GREETING' # Hello $GREETING
```

- Prefer single quotes over double quotes

----

- Quotes of one kind protect quotes of another

```sh
echo "Don't do that!" # Don't do that
echo 'Hello "John"' # Hello "John"
```

----

## Multiline Strings

- Quotes allow multiline strings

```sh
$ echo 'hello
> world'
hello
world
```

- `>` is a secondary prompt
  - stored in PS2 variable

----

## Not Multiline Strings

```sh
$ echo abc\
> def\
ghi
abcdefghi
```

- trailing `\` causes the real newline to be ignored

---

## Shell Scripts (Basic)

Put this script in `pgrep.sh`

```sh
echo "Finding program containing $1"
ps aux | grep "$1"
```

- run it with `sh script.sh ssh`
  - `$1` is replaced with `ssh`
- Notice the quotes around `$1`
  - Prevents globbing and word splitting
- User must know that `pgrep.sh` is to be run with `sh` and not `python`, `zsh`, ...

Notes:

- `shellcheck` is valuable linter and VSCode extension

----

## Shell Scripts (Preferred)

Put this script in `pgrep.sh`

```sh
#!/bin/sh
echo "Finding program containing $1"
ps aux | grep "$1"
```

- `chmod +x pgrep.sh`
  - run it with `./pgrep.sh`

----

## Shebang

- `#!` is a [shebang](https://en.wikipedia.org/wiki/Shebang_(Unix))
- runs `/bin/sh pgrep.sh`
  - not the same as `/bin/sh < pgrep.sh` due to STDIN of `sh` no longer connected to terminal
- Interpreted by the kernel
  - see [binfmt](https://github.com/torvalds/linux/blob/v6.5/fs/binfmt_script.c#L41)
  - just like ELF formatted binaries (non-script executables)

----

- Interpreter `/bin/sh` can be replaced with any program that can be called like `interp filename`
- if full path is not know to interpreter

```sh
#!/usr/bin/env myinterpreter
my program source code
```

- Can include arguments to interpreter
  - If run like `sh script.sh` then shebang is ignored (commented out)

```sh
#!/bin/sh -ex
# -x echos all commands to STDERR
# -e exits the script on first error
your-code-here
```

---

## Adding Programs to Your PATH

- Can add `pgrep.sh` to your PATH
  - `cp pgrep.sh ~/.local/bin/mypgrep`
  - `.bashrc` probably has this already

```sh
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/bin" ] ; then
    PATH="$HOME/bin:$PATH" # colon separates directories
fi

# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi
```

- Allows calling the program like `mypgrep`

----

## Program Hash Table

- Shells first consult a hash table for program's location
  - Fullback to searching `PATH`
  - Can cause problems if executable changed paths or multiple matches

- builtin `hash` provide access to hash table
  - `hash` displays the hash table and hit counts
  - `hash -r` drops cache (like a new shell)

- `bash` has more options for `hash`

----

## Program Arguments

- Within a script program arguments are `$1`, `$2`, ... , `$9`, `${10}`, ...
  - Can use `${1}`
- `$0` is the program name
- `$*` is all arguments (space separated)
  - `$@` is similar

---

## If Statements

```sh
if [ -f /path/to/file ]
then
  echo "File exists"
else
  echo "File does not exist"
fi

# Equivalent one-liners
if [ -f /path/to/file ]; then echo "File exists"; else echo "File does not exist"; fi
[ -f /path/to/file ] && echo "File exists" || echo "File does not exist"
```

- `test -f /path/to/file` implicitly called by []

Notes:

- Also an else if clause called `elif`

---

## Loops

```sh
names='Stan Kyle Cartman'
for name in $names
do
  echo $name
done

# equivalent one-liner
for name in $names; do echo $name; done
```

- `$names` can be replaced with a glob or command

Notes:

- There are also `while` loops

---

## Exported Variables

- `export MYVAR` or `export MYVAR=MYVALUE`
- *export*ed variables are inherited by children
- non-*export*ed variables are only local to the current process

---

## Sudo Gotchas

```sh
curl https://myscript.sh | sh -
# but need the script run as root so might try
sudo curl https://myscript.sh | sh -
# sh still runs as non-root 🤦
cmd1 | cmd2 # cmd1 is sudo curl https://myscript.sh
```

- `sudo` is just a command
  - Not a prefix to everything after it in the line
- `sudo program arg1 arg2` run a child process as root
  - `program arg1 arg2`
- one solution: `curl https://myscript.sh | sudo sh`
- `echo "Contents of file" | sudo sh -c "cat > file-owned-by-root.txt"`

Notes:

- To test this works `echo "id" | sudo sh` to run `id` as root
