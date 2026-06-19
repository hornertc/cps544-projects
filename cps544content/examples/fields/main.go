package main

type Foo struct {
	Field1 bool
	Field2 *bool // nil, true, false
}

// {"Field1": true, "Field2": false}  // Field2 -> false
// {"Field1": true}  // Field2 = nil

func f(x *int, y string, z int) (*Foo, error) {
	if y != "" {

	}

	if x != nil {
		xVal := *x
	}

	if err != nil {
		return nil, err
	}
	foo := Foo{}

	return &foo, nil
}

func main() {

	f(nil, "", 5)
}
