package pkg

func Partition(a []Pizza, lo, hi int) int {
	p := a[hi].IngQty
	for j := lo; j < hi; j++ {
		if a[j].IngQty < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func QuickSort(a []Pizza, lo, hi int) {
	if lo > hi {
		return
	}

	p := Partition(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}
