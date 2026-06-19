package main

import (
	"errors"
	"fmt"
	"os"
)

// want this to be const
var ErrUnknownChild = errors.New("unknown child")

type ErrMissingStuff struct {
	Reason  string
	Missing []string
}

// Error implements the error interface
func (e *ErrMissingStuff) Error() string {
	return fmt.Sprintf("missing %q because %s", e.Missing, e.Reason)
}

func main() {

	err := getReadyForDay(os.Args[1])
	if err != nil {
		fmt.Println(err)

		fmt.Printf("%T\n", err)
		fmt.Println(err == ErrUnknownChild, errors.Is(err, ErrUnknownChild))

		errMissing := &ErrMissingStuff{}
		if errors.As(err, &errMissing) {
			// use errMissing
			fmt.Println("The error was ErrMissingStuff", errMissing.Missing)
		}
		os.Exit(1)
	}
	fmt.Println("Children are ready for school")
}

func getReadyForDay(name string) error {
	if err := getReadyForSchool(name); err != nil {
		return fmt.Errorf("%q failed to get ready for the day: %w", name, err)
		// return err
	}
	// do more stuff, feed them

	return nil
}

func getReadyForSchool(name string) error {
	switch name {
	case "Cathy":
	case "Joe":
	case "John":
		return &ErrMissingStuff{
			"Lost",
			[]string{"Backpack", "Shoes", "Socks"},
		}
	case "Bobby":
		return &ErrMissingStuff{
			"Stolen",
			[]string{"Wallet"},
		}
	default:
		return ErrUnknownChild
	}

	return nil
}
