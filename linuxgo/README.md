[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/1LAOf4oD)

# Assignment: Linux, Virtual Machines, and Go

In this assignment you will complete a short list of tasks in preparation for the rest of the course.

![Linux](https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/225px-Tux.svg.png) ![VirtualBox](https://upload.wikimedia.org/wikipedia/commons/thumb/d/d5/Virtualbox_logo.png/180px-Virtualbox_logo.png) ![Go](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/322px-Go_Logo_Blue.svg.png)

## Learning Objectives

- Able to create a Linux environment for development
- Familiarity with Linux
- Able to run a Go program

## Background

Please read over the documentation for Linux, virtual machines, and Go in our course [resources page](https://github.com/ktarplee-courses/cps544-content/blob/main/resources.md).

## Tasks

1. Create a [Google Cloud Shell Editor](https://console.cloud.google.com/cloudshelleditor) by logging in with your "@udayton.edu" account.  
    1. Run `lsb_release -a`.  Do you recognize anything in the output? Create a file named `cloudshell/release` in the repository with the contents being the output of `lsb_release -c`.  
    1. Run `uname -a`.  What component of Linux does this command inform you about?  Run `uname` and put the output in `cloudshell/uname`.
    1. What is the total system memory?  Put this number in [mebibytes](https://simple.wikipedia.org/wiki/Mebibyte) in the file `cloudshell/memory`.
  
1. Create a Linux virtual machine and ensure you have the guest additions installed and that you have root permissions (ideally sudo).
    1. Run `lsb_release -a`.  Is the output different than Cloud Shell.  Create a file named `vm/release` in the repository with the contents being the output of `lsb_release -c`.  
    1. Run `uname -a`.  Which system is running a newer Kernel?  Run `uname` and put the output in `vm/uname`.

1. Write a Go program that outputs hello in Bulgarian.  (i.e., "Здравейте").
    1. Run `go mod init greeting`
    1. Put the program in the `cmd/greeting` folder

1. Ensure that all the GitLab Actions pass.

## Hints

These instructions and hints are intentionally vague so you get the experience of hopefully running into an issue that you need to problem solve your way through.

### Install the VirtualBox Guest Additions

Guest Additions is a very useful addition for VirtualBox guest operating systems.  It allows better pointer integration and window/monitor resizing.

1. Become root in the guest with `su -`
1. In VirtualBox select "Devices -> Install Guest Additions CD Image".  This will mount the CD in the guest.  
1. Run `apt update && apt install build-essential` as root.  There might be other dependencies needed (like the linux headers depending on your Linux distribution and version).
1. Run the `autorun.sh` script from the CD image that is mounted.  You should be able to double-click on the script to run it.

### Add Yourself to sudo Group

It is useful to have sudo setup to be able to easily administer your VM.

1. Run `adduser vboxuser sudo` (replace vboxuser with your username use to create the VM).
1. Logout of the VM and log back in.
1. `id` should show that you are in the "sudo" group.  
1. Try `sudo apt update` and see if `sudo` works for you works.

### Format of Committed Files (Answers)

For `cloudshell/release`, etc. please include the trailing newline (not a carriage return followed by a newline).

## Submission

Please commit and push your changes to your assignment's project on GitHub.
