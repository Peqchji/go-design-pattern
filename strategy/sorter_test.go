package strategy_test

import (
	"design_pattern/internal/result"
	"design_pattern/strategy/sorter"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestAdaptiveSorterTestcase struct {
	name     string
	input    []int
	expected result.Result[[]int]
}

func TestAdaptiveSorter(t *testing.T) {
	table := []TestAdaptiveSorterTestcase{
		{
			name:  "Should sorting unsorted input",
			input: []int{-1, 2, 3, 2, 3, 4, 10, 1, 9},
			expected: result.Result[[]int]{
				Result: []int{-1, 1, 2, 2, 3, 3, 4, 9, 10},
				Error:  nil,
			},
		},
		{
			name:  "Should sorting sorted input",
			input: []int{1, 1, 2, 2, 3, 3, 4, 9, 10},
			expected: result.Result[[]int]{
				Result: []int{1, 1, 2, 2, 3, 3, 4, 9, 10},
				Error:  nil,
			},
		},
		{
			name:  "Should return empty slices",
			input: []int{},
			expected: result.Result[[]int]{
				Result: []int{},
				Error:  nil,
			},
		},
	}

	for _, tc := range table {
		quickSortStrategy := new(sorter.QuickSortStrategy[int])
		mergeSortStrategy := new(sorter.MergeSortStrategy[int])

		t.Run(tc.name, func(t *testing.T) {
			adaptiveSorter := sorter.NewAdaptiveSorter(
				2,
				mergeSortStrategy,
				quickSortStrategy,
			)

			actualRes, actualErr := adaptiveSorter.Sort(tc.input)

			if tc.expected.Error != nil {
				if assert.Error(t, actualErr) {
					assert.Contains(
						t,
						actualErr.Error(),
						tc.expected.Error.Error(),
					)
				}
			} else {
				assert.Equal(t, tc.expected.Result, actualRes)
			}
		})
	}
}


func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100000)
	}
	return arr
}


// LLM generated Bench mark
// ---------------------------------------------------------
// Round 1: Small Data (N=50)
// Expectation: QuickSort should win (less allocation overhead)
// ---------------------------------------------------------

func BenchmarkQuickSort_Size50(b *testing.B) {
	qs := new(sorter.QuickSortStrategy[int])
	input := generateRandomArray(50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// must copy to ensure we aren't sorting already sorted data
		// (QuickSort is fast on sorted data, which would cheat)
		qs.Sort(input)
	}
}

func BenchmarkMergeSort_Size50(b *testing.B) {
	ms := new(sorter.MergeSortStrategy[int])
	input := generateRandomArray(50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ms.Sort(input)
	}
}

// ---------------------------------------------------------
// Round 2: Large Data (N=10,000)
// Expectation: MergeSort *might* catch up due to QuickSort's worst cases
// or Deep Recursion, but usually QuickSort still wins in Go.
// Let's see if the Allocation cost of MergeSort kills it.
// ---------------------------------------------------------

func BenchmarkQuickSort_Size10k(b *testing.B) {
	qs := new(sorter.QuickSortStrategy[int])
	input := generateRandomArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qs.Sort(input)
	}
}

func BenchmarkMergeSort_Size10k(b *testing.B) {
	ms := new(sorter.MergeSortStrategy[int])
	input := generateRandomArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ms.Sort(input)
	}
}
