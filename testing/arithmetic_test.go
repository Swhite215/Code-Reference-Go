package tests

import (
	"testProject/utilities"
	"testing"	
)

func TestSum(t *testing.T) {
	total := utilities.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestDifference(t *testing.T) {
	difference := utilities.Difference(10, 5)

	if difference != 5 {
		t.Errorf("Difference was incorrect, got %d, want: %d", difference, 5)
	}
}

func TestProduct(t *testing.T) {
	product := utilities.Product(10, 10)

	if product != 100 {
		t.Errorf("Product was incorrect, got %d, want: %d", product, 100)
	}
}

func TestQuotient(t *testing.T) {
	quotient := utilities.Quotient(100, 5)

	if quotient != 20 {
		t.Errorf("Quotient was incorrect, got %d, want: %d", quotient, 20)
	}
}
