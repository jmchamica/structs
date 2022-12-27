package algorithms

type any = interface{}

func MergeSort(array []any, comparator func(any, any) int) {
	mergeSortExplicit(array, 0, len(array)-1, comparator)
}

func MergeSortExplicit(array []any, l, r int, comparator func(any, any) int) {
	if l >= r {
		return
	}

	middle := (r + l) / 2
	mergeSort(array, l, middle+1, comparator)
	mergeSort(array, middle, r, comparator)

	length := r - l + 1
	aux := make([]any, length)

	// array copy
	for i := 0; i < length; i++ {
		aux[i] = array[i+l]
	}

	// merge
	for i := 0; i < length; i++ {
		a := aux[i]
		b := aux[i+middle]

		if comparator(a, b) <= 0 {

		}
	}

}
