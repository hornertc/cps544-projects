[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/2MOrTyot)
# Assignment: Linux, Virtual Machines, and Go

In this assignment you will set out on a scavenger hunt trying to find very specific answers to very specific questions.

![Magnifying glass](https://upload.wikimedia.org/wikipedia/commons/thumb/e/ea/Magnifying_glass_with_focus_on_paper.png/640px-Magnifying_glass_with_focus_on_paper.png)

## Learning Objectives

- Capable of completing simple tasks with the UNIX shell
- Increased knowledge of basic Git commands
- Combine knowledge of Git and shell commands to accomplish simple tasks

## Tasks

Each task is worth two points.  Place your answer for the task in a file called `tasks/x.txt` where `x` is the task number.  The file should contain the answer followed by a new line.  You can check your work by running the script `grade.sh`.  All your work must be done from the command line to receive any credit for this assignment (i.e., not from a web browser).  In a file called `Work.md` please document the command(s) you used to accomplish each task.  Please format the document nicely in Markdown.  You do not have to write a script to accomplish the task but you do need to use the shell to find the answer.  Some of these tasks require that you combine a few shell commands to accomplish the task.

1. In our course content repository (at `git@github.com:ktarplee-courses/cps544-content.git`) there is a commit with a message "added the GO review sheet".  What is the commit digest of this commit?

2. In the Kubernetes repository (at `git@github.com:kubernetes/kubernetes.git`) (written entirely in Go by the way) there is a commit `5a6acf85fab839a99fb60063c0c97ac353b78c68`.  Who are the code OWNERS of the `pkg` directory?  Use the file `pkg/OWNERS` as your answer.

3. In Kubernetes, what is the commit digest of the `v1.25.0` tag?

4. In Kubernetes, how many tags have a major version of 0 (e.g., begin with `v0.`)?

5. In Kubernetes, how many merge commits where made to `v1.25.0` in its entire history?  Hint:  There are two some(squashed) commits that have commit messages with the word "Merge:" that you need to ignore since they are not true Git merge commits.

## Hints

These instructions and hints are intentionally vague so you get the experience of hopefully running into an issue that you need to problem solve your way through.

## Submission

Please commit and push your changes to your assignment's project on GitHub.  The GitHub Actions workflow will show your grade for the assignment but you can loose points for not properly documenting your steps in `Work.md`
