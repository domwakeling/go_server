package data

import "testing"

func TestChecksValidation(t *testing.T) {
	v := NewValidation()
	p := &Product{}

	// should fail on an empty product
	err := v.Validate(p)
	if err == nil {
		t.Fatal(err)
	}

	// should fail because the SKU is invalid
	p.Name = "some name"
	p.Price = 1.0
	p.SKU = "abc1234"
	err = v.Validate(p)
	if err == nil {
		t.Fatal(err)
	}

	// should pass with a valid SKU
	p.SKU = "abc-abcd-ab"
	err = v.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
