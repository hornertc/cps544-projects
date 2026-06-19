## Basic UNIX and Shell

---

### UNIX/Linux Architecture

- Kernel (Linux or BSD)
- Everything is a file (e.g., socket, pipe, serial port, hard drive, shared memory)
  
- `/proc` (e.g., `/proc/cpuinfo`, `/proc/12345/environ`)
- `/dev/sda`, `/dev/tty0`, `/dev/urandom`, `/dev/stdout`, `/dev/shm`
- Virtual Filesystem (VFS) - single directory hierarchy for all mounted filesystems

Notes:

- Multi-user system
- Single-user mode (for troubleshooting)
- Plan9 (developed by Rob Pike, Ken Thompson)
  - Considered "better UNIX" but abandoned
  - "everything is a file" to the logical conclusion
    - includes process state
  - Pike, et. al.
    > [t]he foundations of the system are built on two ideas: a per-process name space and a simple message-oriented file system protocol.

---

### Filesystem Paths

- File paths form a tree
- The root is at `/`
  - Directories are joined by `/`, e.g., `/home/bob/file.txt`
- `..` is the parent directory (up the tree)
- `.` is the current directory
- Relative paths do not start with a `/`
  - `mydir/file.txt`
  - `../mydir/myfile.dat`

----

### Important Paths

- `/boot` pointed to by the MBR (master boot record), contains the bootloader and kernel
- `/usr` (UNIX System Resources)
- `/usr/local` non-system resources that are available to all users
- `/home` home directories (user local resources)
- `/etc` system level configuration

----

### Important Subpaths

- `/usr/bin`, `/usr/local/bin` is for executables (binaries)
- `/usr/lib`, `/usr/local/lib` is for libraries
- `/usr/include`, `/usr/local/include` is for C/C++ header files to "include"
- `/usr/sbin` is for admin executables (super user binaries)

Notes:

- More at Filesystem Hierarchy Standard [FHS](https://en.wikipedia.org/wiki/Filesystem_Hierarchy_Standard)
- System package manager (e.g., `apt/dpkg`, `yum/rpm`, `pacman`) installs to `/usr`.  Do not manually put file there!
  - Use `/usr/local` for multi-user resources
  - Home directory for single-user resources

----

### Home Directory

- `~` expands to the absolute path of the current user's home directory
- `~/.bashrc`, `~/.profile`, `~/.bash_profile` are for configuring the bash shell
- `~/.ssh` is for storing your "encrypted" SSH key pairs and client configuration
- `~/.config`, a.k.a., `XDG_CONFIG_HOME` is the **new** location for configuration files
- `~/.local/bin` user's personal executables
- `~/.gnupg` is for GNU Privacy Guard, GnuPG, GPG key pairs, and config

---

### Getting Help

- Command help (e.g., `ssh-keygen --help`)
- Manual pages (e.g., `man ssh-keygen`) for more detailed help
- `whatis` - search **all** man pages efficiently (but only the one-line description)
- `apropos` - search **all** man pages efficiently (full descriptions)
- `info` - navigate documentation
- Web search engine (e.g., Google, Bing)

Notes:

- `whatis chmod`
- `man chmod`
- `man 2 chmod`
- `apropos permissions`

----

### Man Page Sections

1. Executable programs or shell commands
2. System calls (functions provided by the kernel)
3. Library calls (functions within program libraries)
4. Special files (usually found in /dev)
5. File formats and conventions, e.g. /etc/passwd
6. Games
7. Miscellaneous (including  macro  packages  and  conven‐
           tions), e.g. man(7), groff(7), man-pages(7)
8. System administration commands (usually only for root)
9. Kernel routines [Non-standard]

Notes:

- `man man`

### Getting More Help

- Read the source code
  - Likely open source
  - Might even be on your filesystem
  - Open/Closed issues
  - Stack Overflow or similar web sites

---

### User and Group Management

- `/etc/passwd` has all users, `/etc/group` has all groups
- `/etc/shadow` stores hashed passwords
- Root is UID/GID 0
  - Obtained via `sudo` or `su`
- Regular users often start at UID/GID 1000
- `adduser myfriend` adds a new user
- `adduser username newgroup` adds a user to a group

----

### Filesystem Permissions

- UID/GID int is stored on disk (not user and group strings)
  - `chown` to change owner or group
- Permissions stored as an int (bitmask) on disk
  - `chmod` to change permissions
    - user, group, other (ugo)
    - read, write, execute (rwx)
  - octal or symbolic mode
  - special bits (setuid, setgid, sticky)

---

### Environment Variables

- Search paths
  - `PATH` - for binaries
  - `MANPATH` - for man pages
    <!-- - `CDPATH` - for directory alternate names -->
- `PWD` - current working directory
- `HOME` - home directory
- `SHELL` - the current path to the shell executable

---

### Other commands

- `stat` - file and permission information
- `file` - inspect the contents of the file to determine it's type
  - Magic numbers vs. file extensions
- `sudo`, `apt`, `su`
- `xdg-open`
