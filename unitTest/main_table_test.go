package main

import (
	"fmt"
	"testing"
)

//Teste de Mesa
func TestTableDriven(t *testing.T) {
	tables := []struct {
		x    int
		y    int
		want int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		testName := fmt.Sprintf("%d,%d", table.x, table.y)

		t.Run(testName, func(t *testing.T) {
			total := Add(table.x, table.y)
			if total != table.want {
				t.Errorf("Sum of (%d + %d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.want)
			}
		})
	}
}
