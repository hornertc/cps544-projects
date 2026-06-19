package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go Shell (gos)")

	script := os.Stdin
	if len(os.Args) >= 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		script = f
	}

	scanner := bufio.NewScanner(script)

	for true {
		fmt.Print(os.Getenv("MYPS1"))
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if err := processLine(line); err != nil {
			log.Printf("Error on line %q, %s", line, err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLine(line string) error {
	words := strings.Split(line, " ") // FIXME this ignores quotes and escaping

	env := os.Environ()
	args := []string{}
	fds := []*os.File{
		os.Stdin,
		os.Stdout,
		os.Stderr,
	}

	// TODO pipe |
loop:
	for _, word := range words {
		switch {
		case len(word) == 0:
			// whitespace
			continue loop
		case strings.HasPrefix(word, "#"):
			// comment
			break loop
		case len(args) == 0 && strings.Contains(word, "="):
			// environment variable
			env = append(env, word)
		case strings.HasPrefix(word, ">>"):
			// append stdout to file
			fn := word[2:]
			f, err := os.OpenFile(fn, os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fds[1] = f

		case strings.HasPrefix(word, ">"):
			// write stdout to file
			fn := word[1:]
			f, err := os.Create(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fds[1] = f

		case strings.HasPrefix(word, "<"):
			// read stdin from file
			fn := word[1:]
			f, err := os.Open(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fds[0] = f

		case strings.HasPrefix(word, "2>"):
			// write stderr to file
			fn := word[2:]
			f, err := os.Create(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fds[2] = f

		case strings.HasPrefix(word, "3>"): // TODO expand to arbitrary file descriptor numbers
			// write stderr to file
			fn := word[2:]
			f, err := os.Create(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			fds[3] = f

		default:
			args = append(args, word)
		}
	}

	if len(args) == 0 {
		return nil
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Env = env
	cmd.Stdin = fds[0]
	cmd.Stdout = fds[1]
	cmd.Stderr = fds[2]
	cmd.ExtraFiles = fds[3:]

	return cmd.Run()
}
