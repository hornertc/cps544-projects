![Git Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Git-logo.svg/512px-Git-logo.svg.png)

---

### Overview

- Git is a distributed version control system or source control management (SCM)
  - As opposed to a centralized version control system (e.g., CVS, SVN)
- Developed by Linus Torvalds (creator of Linux)
  - BitKeeper (distributed SCM) revoked its free license for the Linux kernel development team
- Initial release in 2005

---

### Basic Concepts

- **Repository** a local directory with a `.git` directory
- **Commit** an immutable change to the source code (has a parent commit, comment message, author, and a patch)
  - Identified by their SHA1 digests

----

- **Branch** a mutable human readable name that points to a commit
  - All branches are local
  - Some branches track/clone a branch in a remote repository
- **Tag** an immutable human readable name (often a version number) that points to a commit

Notes:

- Often hear I should have created a "branch" for this feature
  - Everything listed by `git branch -a` is a branch
  - `main` and `master` are branches (not special to Git but special to GitHub)
  - Can only work on a branch
  - So developers are always creating a **local** branch if they want to or not
- Developers often do not internalize that every branch is local in Git

---

### Basic Commands

<!-- ![Commands](https://upload.wikimedia.org/wikipedia/commons/d/d8/Git_operations.svg) -->
<img src="https://upload.wikimedia.org/wikipedia/commons/d/d8/Git_operations.svg" alt="Git operations" height="500">

Notes:

- `git fetch` and `git push` are remote only
- `git pull` is a fetch + merge
- Others commands are local only
- Remote is a different Git repository (accessed via SSH, HTTPS, git, or file protocols)
- Clone is a local Git repository
- Branches are names that point to the tip (HEAD) of a sequence of commits
- Working files are the files you actually see in the directory
  - Some of which might be staged to be committed

---

### Workflow

Simplest Git workflow

```sh
git clone https://github.com/...
# edit source code files
git add file1 file2
git commit -m"added unicode support"
git push # to share with others
```

----

### Typical non-GitHub Workflow

```sh
git clone https://github.com/...

git switch -c feature-unicode
# edit source code files
git add file1 file2
git commit -m"added unicode support"
git push # to share with others
```

Once `feature-unicode` is ready (team members agree on its status) then merge it into main.

```sh
git switch main
git pull # double check that "main" is up to date
git merge feature-unicode
git push
```

----

### Other Workflows

- Extra credit assignment uses the GitHub workflow
  - Common in OSS projects and industry
- [Trunk based development](https://trunkbaseddevelopment.com/) has many more workflows and in-depth comparisons between them

---

### Merging

- Git tracks content, not changes per file and line number
  - Developer X moves `f()` definition from file `a.sh` to `b.sh`
  - Developer Y fixed a bug in `f()` on a branch A (where `f()` is defined in `a.sh`)
  - Developer X pulls in developer Y's fix to `f()`
    - Git knows the content (i.e. `f()`) was in `a.sh` and was moved to `b.sh`
    - Git applies developer Y's change to the content in `b.sh`

----

### Conflicts

- Occurs when developer X and Y (on different branches) modify the same lines of code (chunk)
- Git will fail to merge those changes
- Forces the developer to take action to resolve merge conflicts
- Merge conflicts are a good thing 😲
  - Alternative is silent source code overwriting and corruption

---

### Git Tools

- `gitk --all` graphically displays the history (i.e., commit tree) for **all** branches including branches tracking remote branches
- `gitg` also graphically displays the history
- `git gui` is a GUI for making commits interactive.  Useful for reviewing changes and selecting parts of files to stage for commit.

Note:

- VSCode also has GUIs that support most of the features of the above tools

---

### Advanced Features

- `git cherrypick`
- `git rebase` - rewrite history (only in Git not in real life)
  - Change author emails or commit messages
  - Squash commits
  - Moves commits over to a different branch
- `git gc` - garbage collection
- `git fsck` - integrity checking
- signed commits and tags
- submodules and subtree

---

### Warnings

- Cloning downloads **all** commits by default
  - Want repository size to be small
- **Do not commit binary files**
  - Everytime the file is changed (even one bit) an entire copy of the file is added to the git history
- **Do not commit build products**
- Git LFS is a useful way to manage large files
  - Stores only the file's digest and size in the git history
  - File itself is stored on the server (github.com) and download only when requested
