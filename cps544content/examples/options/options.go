package options

import "fmt"

type MyObject struct {
	a bool
	b int
	c string
}

type Option func(o *MyObject) error

// Constructor
func New(options ...Option) (*MyObject, error) {
	// Set defaults
	obj := &MyObject{
		a: true,
		b: 13,
	}

	for _, o := range options {
		if err := o(obj); err != nil {
			return nil, err
		}
	}

	// check compatibility of options

	// do any extra initialization

	return obj, nil
}

func WithA() Option {
	return func(o *MyObject) error {
		o.a = true
		return nil
	}
}

func WithoutA() Option {
	return func(o *MyObject) error {
		o.a = false
		return nil
	}
}

func WithB(b int) Option {
	return func(o *MyObject) error {
		o.b = b
		return nil
	}
}

const optimizationFlag = "-o go-fast"

func WithOptimization() Option {
	return func(o *MyObject) error {
		o.c += optimizationFlag
		o.b++
		return nil
	}
}

func basicUsage() {
	obj, err := New(WithB(45), WithOptimization())
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}

func usage(overrides []Option) {
	o := []Option{WithA(), WithB(45)}
	o = append(o, overrides...)
	obj, err := New(o...)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
