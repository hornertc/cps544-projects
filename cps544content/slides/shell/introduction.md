## Introduction to the UNIX Shell

Open a terminal of your own and please follow along.

---

## Shell

- Shell prompt `$`
- Commands are used to start a new process:
  - `COMMAND OPTIONS ARGUMENT1 ARGUMENT2...`
  - `ls mydir`
  - `ls -l mydir myotherdir`
  - Arguments are "words" separated by spaces and tabs
  - Options or flags are optional modifiers
  
---

## Filesystem Commands

- `mv` moves/rename a file
  - overwrites the target <!-- .element: class="warning" -->
- `cp` copies a file
- `rm` removes a file
  - not undo-able <!-- .element: class="warning" -->
  - `rm -rf /` would recursively delete your entire system <!-- .element: class="warning" -->

Notes:

- `mv -i` and `rm -i`

---

## Security Caution

- Never run a command on your system without **fully** understanding what it will do
  - Includes commands copied from the internet
- Command you run have the **full permissions** of your user on that system

Notes:

- See [Explain Shell](https://explainshell.com/)

---

## Viewing File Contents

- `cat` prints the contents of the files named by its arguments
- `more` prints one page at a time (known as a "pager")
- `less` prints one page at a time but allows backward pagination
  - hence `less` is `more`
- `head` prints the start of a file
- `tail` prints the end of a file

---

### Searching Content in Files

- `wc` prints the number of character, words, and lines
- `grep` searches for a pattern in files
- `cmp` prints the location of the first difference (byte by byte comparison)
- `diff` shows the difference between files (a.k.a., patch)

---

## Filesystem

- Files are arranged in a tree structure
  - `/` is the top-most (root) directory
  - All other files and folders are under root directory
- The shell runs commands at a directory
  - `pwd` prints the current working directory
  - `cd` changes the current working directory
  - `ls` list content of current directory (or pass a path to view details of that path)

---

## Editors

- `nano`
  - commands are displayed at the bottom of window
  - `^o` means save.  Type Control (Ctrl) and the letter O at the same time
  - will be prompted for filename then hit enter to save
- `vi` and `vim`
  - online interactive [tutorial](https://www.openvim.com/)
