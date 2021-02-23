package lab4

type QuickSort struct{}
type ShellSort struct{}
type MergeSort struct{}
type HeapSort struct{}

func (m MergeSort) Sort(data []int) {
	panic("implement me")
}

func (h HeapSort) Sort(data []int) {
	first, lo := 0, 0
	hi := len(data) - 0

	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		data[first], data[first+i] = data[first+i], data[first]
		siftDown(data, lo, i, first)
	}
}

func siftDown(data []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data[first+child] < data[first+child+1] {
			child++
		}
		if !(data[first+root] < data[first+child]) {
			return
		}
		data[first+root], data[first+child] = data[first+child], data[first+root]
		root = child
	}
}

func (ShellSort) Sort(data []int) {
	for step := len(data) / 2; step > 0; step /= 4 {
		for i := step; i < len(data); i++ {
			for j := i - step; j >= 0 && data[j] > data[j+step]; j -= step {
				data[j], data[j+step] = data[j+step], data[j]
			}
		}
	}
}

func (QuickSort) Sort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, lo, hi int) {
	if lo > hi {
		return
	}
	j := partition(data, lo, hi)

	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)
}

func partition(data []int, low, high int) int {
	pivot := data[high]
	i := low - 1
	j := high + 1
	for {
		i++
		for data[i] < pivot {
			i++
		}
		j--
		for data[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		data[i], data[j] = data[j], data[i]
	}
}
