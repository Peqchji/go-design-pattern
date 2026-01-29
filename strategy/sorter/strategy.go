package sorter

import "cmp"

type SortStrategy[T cmp.Ordered] interface {
	Name() string
	Sort(arr []T) []T
}

type QuickSortStrategy[T cmp.Ordered] struct {}

func (q *QuickSortStrategy[T]) Name() string { return "Quick Sort Strategy" }

func (q *QuickSortStrategy[T]) Sort(arr []T) []T {
    n := len(arr)

	if n < 2 {
		return arr
	}
    
    result := make([]T, n)
    copy(result, arr)

	q.quickSort(result, 0, len(arr) - 1)

	return result
}

func (q *QuickSortStrategy[T]) quickSort(arr []T, low, high int) {
	if low < high {
		p := q.partition(arr, low, high)

		q.quickSort(arr, low, p - 1)
		q.quickSort(arr, p + 1, high)
	}
}

func (q *QuickSortStrategy[T]) partition(arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i + 1], arr[high] = arr[high], arr[i + 1]

	return i + 1
}



type MergeSortStrategy[T cmp.Ordered] struct{}

func (m *MergeSortStrategy[T]) Name() string { return "Merge Sort Strategy" }

func (m *MergeSortStrategy[T]) Sort(arr []T) []T {
    if len(arr) < 2 {
        return arr
    }
    // Merge Sort is naturally creating new slices usually
    return m.mergeSort(arr)
}

func (m *MergeSortStrategy[T]) mergeSort(arr []T) []T {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := m.mergeSort(arr[:mid])
    right := m.mergeSort(arr[mid:])
	
    return m.merge(left, right)
}

func (m *MergeSortStrategy[T]) merge(left, right []T) []T {
    result := make([]T, 0, len(left)+len(right))
    i, j := 0, 0
    
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}