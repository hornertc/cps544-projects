# Resources

Below is a list of resources for students.

## Linux

- [Google Cloud Shell Editor](https://console.cloud.google.com/cloudshelleditor) is a useful VSCode-Like IDE that includes Go, SSH, and Git preinstalled. Login with your `@udayton.edu` account.  It provides a containerized (i.e., limited) Linux (based on Debian) environment but is still very useful.  Root access on the container can be achieved via `sudo` which is useful for installing packages via `apt`.

### Virtual Machines

To create a Vm using any of the methods below you will need a ISO (a virtual CD/DVD with the operating system on it).  Ubuntu is the recommended Linux distribution and here is the [ISO download](https://ubuntu.com/download/desktop).

- VirtualBox [download](https://www.virtualbox.org/)
- [UTM](https://mac.getutm.app/) can be used on a Mac (including Apple M1 hardware) to create a VM but you need to make sure you download the [ARM64 image of Ubuntu](https://www.youtube.com/redirect?event=video_description&redir_token=QUFFLUhqbmxaT3doYkJVcldHbmlyY1lzaUl4QlpnQ1JaQXxBQ3Jtc0trZFFsT0NhWjMxTDdCeUlyTWxvb2lKWWg3TmdmN3RMdUlpME5idHZleThnUGFlNmN6d2tXTGwxWTE2VWp0ZVlTUktPa0JRVlRQSUFMdGFfaDRGYmhSRGpab3U5RVk3dEVuRmJVUXA4NE9lcl9tM1dLMA&q=https%3A%2F%2Fubuntu.com%2Fdownload%2Fserver%2Farm&v=O19mv1pe76M).  More details are in this [video](https://www.youtube.com/watch?v=O19mv1pe76M).
- [Multipass](https://multipass.run/) can also be used to quickly spin up a VM, but it only provides a shell and no GUI (which is sufficient for this class). This also works on Apple M1/M2 hardware.

- [UD VDI](https://drive.google.com/file/d/1QI-AfYwGtikXn3KlWlIy9AY2UewrZdvI/view) provides an alternative to creating a Linux VM but on UD's infrastructure.

To configure a Linux system for this course see the instructions [here](./Setup.md).

## Git

- [Git-SCM](https://git-scm.com/) is the official website for Git.
- Run `git help -g` to access various guides.  `git help everyday` is a great place to start reading, followed by `git help core-tutorial`.
- [Trunk based development](https://trunkbaseddevelopment.com/) is the most common way to manage development within Git and similar SCM systems.  "Trunk" is synonymous with "master" and "main".
- [Signing commits with SSH](https://docs.gitlab.com/ee/user/project/repository/ssh_signed_commits/)

## Shell

- [Explain Shell](https://explainshell.com/) decodes a shell command and explains what it does.

## Go Programming

- [Downloads and install](https://go.dev/)
- [Tour of Go](https://go.dev/tour)
- [API docs](https://pkg.go.dev/std)
- [Docker](https://www.docker.com) or [Podman](https://podman.io) can be used to create a limited Linux environment.  The container image `docker.io/library/golang` has the Go toolchain included.
- [Online Go Assembly Viewer](https://godbolt.org/)

### Configure golangci-lint

Go provides very powerful linters to help you improve code quality.  `golangci-lint` incorporates nearly all useful linters into one command.  You should consider using some of these linters to help find potential issues with your code.

Refer the [installation instructions](https://golangci-lint.run/usage/install/) or run `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` to install `golangci-lint`.  Make sure `$(go env GOPATH)/bin` is on your `PATH`.

Add this `.golangci.yaml` ([example](https://gist.github.com/ktarplee/42e83fdebe817430d5c93d1de99f5402)) file to the top level of the project and then run `golangci-lint run`.  Correct any issue it finds.

IDEs such as VSCode can be configured to use `golangci-lint` whenever a source code file is saved.  It will also be indicated in the source code (with yellow lines) where specifically the issue is.  Hovering over will let you read the error.

## Software Engineering

- [12 factor app](https://12factor.net/)

## Bash Scripting

- Best practices and a coding standard are documented [here](https://github.com/icy/bash-coding-style?tab=readme-ov-file)
