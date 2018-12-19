package selectionsort

func selectionSort(items []int) []int {
	length := len(items)
	for i := 0; i < length; i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			if items[minIdx] > items[j] {
				minIdx = j
			}
		}
		if items[i] > items[minIdx] {
			items[i], items[minIdx] = items[minIdx], items[i]
		}
	}
	return items
}
