package test

import (
	"testing"
)

func TestAvg(t *testing.T) {
	type AvgResult struct {
		Num    []int
		Result int
	}
	avgValues := []AvgResult{
		{Num: []int{1, 3}, Result: 2},
		{Num: []int{1, 3, 4, 5, 1}, Result: 2},
		{Num: []int{}, Result: 0},
		{Num: []int{1, 7, 9}, Result: 5},
	}

	for _, val := range avgValues {
		if avg := Avg(val.Num...); avg != val.Result {
			t.Fatalf("expected the avg of %v to equal %d but got %d\n", val.Num, val.Result, avg)
		}
	}
}
