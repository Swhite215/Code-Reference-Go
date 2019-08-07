package tests

import (
	"testProject/utilities"
	"testing"	
)

//Singular Test for Utility Function
func TestSum(t *testing.T) {
	total := utilities.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %b, want: %b", total, 10)
	}
}

func TestDifference(t *testing.T) {
	difference := utilities.Difference(10, 5)

	if difference != 5 {
		t.Errorf("Difference was incorrect, got %b, want: %b", difference, 5)
	}
}

func TestProduct(t *testing.T) {
	product := utilities.Product(10, 10)

	if product != 100 {
		t.Errorf("Product was incorrect, got %b, want: %b", product, 100)
	}
}

func TestQuotient(t *testing.T) {
	quotient := utilities.Quotient(100, 5)

	if quotient != 20 {
		t.Errorf("Quotient was incorrect, got %b, want: %b", quotient, 20)
	}
}

// Multiple Tests for Utility Function
type testpair struct {
	values []float64
	expectedResult float64
}

var sumTests = []testpair{
	{[]float64{1,2}, 3},
	{[]float64{5.5,1}, 6.5},
	{[]float64{105,20}, 125},
}

func TestMultipleSum(t *testing.T) {
	for _, pair := range sumTests {
		total := utilities.Sum(pair.values[0], pair.values[1])

		if total != pair.expectedResult {
			t.Error("For", pair.values, "expected", pair.expectedResult, "got", total)
		}
	}
}

var differenceTests = []testpair{
	{[]float64{1,2}, -1},
	{[]float64{5.5,1}, 4.5},
	{[]float64{105,20}, 85},
}

func TestMultipleDifference(t *testing.T) {
	for _, pair := range differenceTests {
		difference := utilities.Difference(pair.values[0], pair.values[1])

		if difference != pair.expectedResult {
			t.Error("For", pair.values, "expected", pair.expectedResult, "got", difference)
		}
	}
}

var productTests = []testpair{
	{[]float64{1,2}, 2},
	{[]float64{5.5,1}, 5.5},
	{[]float64{105,20}, 2100},
}

func TestMultipleProduct(t *testing.T) {
	for _, pair := range productTests {
		product := utilities.Product(pair.values[0], pair.values[1])

		if product != pair.expectedResult {
			t.Error("For", pair.values, "expected", pair.expectedResult, "got", product)
		}
	}
}

var quotientTests = []testpair{
	{[]float64{1,2}, .5},
	{[]float64{5.5,1}, 5.5},
	{[]float64{100,20}, 5},
}

func TestMultipleQuotient(t *testing.T) {
	for _, pair := range quotientTests {
		quotient := utilities.Quotient(pair.values[0], pair.values[1])

		if quotient != pair.expectedResult {
			t.Error("For", pair.values, "expected", pair.expectedResult, "got", quotient)
		}
	}
}
