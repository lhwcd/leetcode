package bubblesort

func bubbleSort(items []int) []int {
	//
	for i := len(items) - 1; i >= 0; i-- {
		//
		swapped := false
		for j := 0; j < i; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
				swapped = true
			}
		}
		//
		if swapped == false {
			break
		}
	}
	return items
}
