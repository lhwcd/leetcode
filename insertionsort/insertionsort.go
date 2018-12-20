package insertionsort

func insertionSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	//
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
			} else {
				break
			}
		}
	}
	return items
}
