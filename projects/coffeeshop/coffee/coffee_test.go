package coffee

import (
	"testing"
)

// init some test data
func init() {
	Coffees = CoffeeList{
		List: []CoffeeDetails{
			CoffeeDetails{"Latte", 2.5},
			CoffeeDetails{"Cappuccino", 2.75},
			CoffeeDetails{"Flat White", 2.25},
		},
	}
}

// make a test case func to test the IsCoffeeAvailable function
func TestIsCoffeeAvailable(t *testing.T) {
	// test case 1
	// check if the function returns true when the coffee type is available
	if !IsCoffeeAvailable("Latte") {
		t.Error("Expected true, got false")
	}

	// test case 2
	// check if the function returns false when the coffee type is not available
	if IsCoffeeAvailable("Mocha") {
		t.Error("Expected false, got true")
	}

	type testcase struct {
		coffeeType string
		want       bool
	}

	cases := []testcase{
		{"Latte", true},
		{"Cappuccino", true},
		{"Flat White", true},
		{"Mocha", false},
		{"", false},
	}

	for _, c := range cases {
		got := IsCoffeeAvailable(c.coffeeType)
		if got != c.want {
			t.Errorf("Expected %v, got %v", c.want, got)
		}
	}
}
