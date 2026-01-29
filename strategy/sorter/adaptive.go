package sorter

import (
	"fmt"
)

type AdaptiveSorter struct {
	thredshold       int
	lowerSortStrategy  SortStrategy[int]
	higherSortStrategy SortStrategy[int]
}

func NewAdaptiveSorter(
	threshold int, 
	lowerSortStrategy SortStrategy[int], 
	higherSortStrategy SortStrategy[int],
) *AdaptiveSorter {

	return &AdaptiveSorter{
		thredshold: threshold,
		lowerSortStrategy: lowerSortStrategy,
		higherSortStrategy: higherSortStrategy,
	}
}

// Your test expects (Result, error).
// Since we removed the 'a' case, this effectively never errors.
func (as *AdaptiveSorter) Sort(input []int) ([]int, error) {
	if input == nil {
		return []int{}, nil
	}

	var sorter SortStrategy[int] = as.lowerSortStrategy
	if len(input) > as.thredshold {
		sorter = as.higherSortStrategy
	}

	fmt.Printf("Using %s\n", sorter.Name())
	
	return sorter.Sort(input), nil
}
