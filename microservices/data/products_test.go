package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Ramesh",
		Price: 10.5,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
