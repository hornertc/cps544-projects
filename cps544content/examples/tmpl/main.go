package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type S struct {
	size      int
	CreatedAt time.Time `json:"created_at"`
}

func ago(t time.Time) (time.Duration, error) {
	d := time.Since(t)

	// Disallow negative durations (to show how errors can be returned)
	if d < time.Duration(0) {
		return d, fmt.Errorf("negative duration")
	}

	return d, nil
}

func computeValues() any {
	created, err := time.Parse(time.RFC3339, "2023-11-01T09:04:05Z")
	if err != nil {
		panic(err)
	}

	return map[string]any{
		"Number":        45,
		"Name":          "John",
		"Prime Numbers": []int{1, 3, 5, 7, 11},
		"Structure": &S{
			size:      13,
			CreatedAt: created,
		},
	}
}

func main() {

	t, err := template.New("root").
		Funcs(template.FuncMap{"Ago": ago}).
		ParseGlob(filepath.Join(os.Args[1], "*.tmpl"))
	if err != nil {
		log.Fatal(err)
	}

	values := computeValues()

	// execute the template
	if err := t.ExecuteTemplate(os.Stdout, "main.tmpl", values); err != nil {
		log.Fatal(err)
	}
}
