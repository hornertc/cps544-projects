# Assignment 3: Scavenger Hunt

## CPS 544-01, Dr. Kyle Tarplee

## Tommy Horner, thorner1

## **Task 1**

**Command:** $ git log --stat --grep="added the GO"

- Displays the information of the commit that contains the words "added the GO"

**Output:**

commit bd5c7fc5e75d8f4bccea6efef1c67133402449f1

Author: Kyle M. Tarplee <tarpleek1@udayton.edu>

Date:   Sat Jul 29 06:31:48 2023 -0700

 added the GO review sheet

 README.md | 6 +++---

 go-review-sheet.md | 65 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

 2 files changed, 68 insertions(+), 3 deletions(-)

**Digest:** bd5c7fc5e75d8f4bccea6efef1c67133402449f1

## **Task 2**

**Command:** $ git checkout 5a6acf85fab839a99fb60063c0c97ac353b78c68

- Shows the state of the repository at the time of this commit

**Output:** Note: switching to '5a6acf85fab839a99fb60063c0c97ac353b78c68'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by switching back to a branch.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -c with the switch command. Example:

  git switch -c new-branch-name

Or undo this operation with:

  git switch -

Turn off this advice by setting config variable advice.detachedHead to false

HEAD is now at 5a6acf85fab Align lifecycle handlers and probes

**Command:** $ cat pkg/OWNERS

- Shows the contents of the pkg/OWNERS file

**Owners:** pkg/OWNERS

## **Task 3**

**Command:** $ git rev-list -1 v1.25.0

- Gets the commit hash for the v1.25.0 tag

**Output:** a866cbe2e5bbaa01cfd5e969aa3e033f3282a8a2

**Digest:** a866cbe2e5bbaa01cfd5e969aa3e033f3282a8a2

## **Task 4**

**Command:** $ git tag -l 'v0.*' | wc -l

- Gets all the tags that begin with 'v0.' and then counts the number of lines

**Output:** 64

**Tag Count (v0.):** 64

## **Task 5**

**Command:** $ git rev-list -n 1 v1.25.0

- Finds the commit that points to v1.25.0

**Output:** a866cbe2e5bbaa01cfd5e969aa3e033f3282a8a2

**Command:** $ git log --merges --invert-grep --grep="Merge:" a866cbe2e5bbaa01cfd5e969aa3e033f3282a8a2^ | grep -c '^commit'

- Gets all of the  merge commits for v1.25.0 that do not contain the word "Merge:" then counts the number of lines that start with the word 'commit'

**Output:** 46332

**Merge Commit Count (v1.25.0):** 46332
