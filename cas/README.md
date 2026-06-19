[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/4WU5cQeq)
# Content Addressable Storage Server

![Database](https://informationage-staging.s3.amazonaws.com/uploads/2022/10/AdobeStock_54409222-1568x1045.jpeg)

## Overview

Often web content is retrieved by its location.  For example a HTTP server might serve an `index.html` file.  The key used to find the file is the name `index.html`.  Content addressable storage (CAS) systems instead use the content itself as the key.  Specifically they use the digest of the content as the key.  In this assignment you will write an HTTP server that does simple CAS storage and retrieval.  The system will have some cryptographic agility since we will allow more than one digest (different algorithms) to be used to retrieve the data.

## Learning Objectives

- Write an HTTP server in Go
- Marshal JSON
- Use cryptographic digests to build a CAS system
- Use interfaces such as `io.Reader`, `io.Writer`, and `http.Handler`

## Requirements

- Support at least the following digests: "sha256", "sha384", "sha512", "sha512-224", "sha512-256"
- Since files can be large:
  - Do not store the contents of a file more than once on disk.  You must use [hard links](https://en.wikipedia.org/wiki/Hard_link) to give a file many "names".
  - Do not store the entire contents of a file in memory for any length of time.
- HTTP servers in Go create a goroutine for each request.  The server must be able to handle multiple simultaneous requests.  For this assignment you are **not allowed** to use any Go synchronization primitives.  This means channels and the "sync" package is forbidden.  However, some POSIX filesystems operations are atomic (e.g, move, delete, link) so you must use those carefully to implement the server.
- Support the following handlers:
  - `POST /blob` to upload a file.  The response must include `x-digest-<name of algorithm>` headers for each digest that the server instance is supporting.
  - `GET /blob/<name of digest algorithm>/<hex encoding of digest>` will retrieve the file by digest.
  - `GET /stats` returns a JSON dictionary with fields
    - `Count` is the integer number of unique files stored
    - `Mean` is the population mean (average) size of the files
    - `Stddev` is the population standard deviation of the size of the files

- Your code will be graded on completeness and form.  You must add appropriate comments to receive full credit.
- You must only edit `pkg/cas/cas.go`.

- Do not use any library other than the Go standard library.
- The source code must compile with the most recent version of the Go compiler.
- The program must not panic under any circumstances.
- All tests (in GitHub) must pass, otherwise you will receive no credit on this assignment.
- Make sure your code is "gofmt'd".  See "gofmt" or better use "goimports" or better yet configure IDE to do this formatting on file save.

## Hints

- First read the entire documentation for the `http` package.  There are many useful helper functions that make the problem easier.
- Consider using `http.ServeMux` to route endpoints
- Do not forget to read the entire HTTP request body and close the the body.
- If the digests of the `testdata` files do not match the values in the unit tests then your GIT installation changed the line endings.  You need to preserve UNIX style line endings.
- You can test the server with `curl`, e.g., `curl -v -X POST --data-binary "@go.mod" localhost:3333/blob/`

## Submission

- Commit and push your working code to your GIT repository.
- Ensure all tests pass otherwise you will receive no credit.
