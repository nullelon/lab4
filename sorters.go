package lab4

type Sorter interface {
	Sort(data []int)
}

type (
	BubbleSorter    struct{}
	SelectionSorter struct{}
	InsertionSorter struct{}
)

func (s BubbleSorter) Sort(data []int) {
	for true {
		isSorted := true
		for i := 1; i < len(data); i++ {
			if data[i] < data[i-1] {
				var temp = data[i]
				a := &temp
				_ = a
				data[i] = data[i-1]
				data[i-1] = temp
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

func (s InsertionSorter) Sort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			for j := 0; j < i; j++ {
				if data[i] < data[j] {
					temp := data[i]
					copy(data[j+1:i+1], data[j:])
					data[j] = temp
					break
				}
			}
		}
	}
}

func (s SelectionSorter) Sort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}
