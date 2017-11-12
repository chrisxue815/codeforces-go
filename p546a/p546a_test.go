package main

import (
	"testing"
)

func TestCalcBorrow(t *testing.T) {
	tables := []struct {
		expected int
		k        int
		n        int
		w        int
	}{
		{13, 3, 17, 4},
	}

	for _, table := range tables {
		actual := CalcBorrow(table.k, table.n, table.w)
		if actual != table.expected {
			t.Errorf("TestCalcBorrow failed. Actual: %d. Table: {%d, %d, %d, %d}",
				actual, table.expected, table.k, table.n, table.w)
		}
	}
}
