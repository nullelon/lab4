package lab4

import (
	"math/rand"
	"testing"
)

func BenchData(n int, isRandom bool) []int {
	data := make([]int, n)

	for i := 0; i < n; i++ {
		if isRandom {
			data[i] = rand.Int()
		} else {
			data[i] = i
		}
	}
	return data
}

func foo(sorter Sorter, data []int) func(*testing.B) {
	return func(b *testing.B) {
		d := make([]int, len(data))
		b.StopTimer()
		b.ResetTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			copy(d, data)
			sorter.Sort(d)
		}
	}
}

func doBench(sorter Sorter, b *testing.B) {
	dataRandom := BenchData(1000000, true)
	dataGrad := BenchData(1000000, false)
	b.Run("rand_1k", foo(sorter, dataRandom[:1000]))
	b.Run("rand_10k", foo(sorter, dataRandom[:10000]))
	b.Run("rand_100k", foo(sorter, dataRandom[:100000]))
	b.Run("rand_1000k", foo(sorter, dataRandom))
	b.Run("grad_1k", foo(sorter, dataGrad[:1000]))
	b.Run("grad_10k", foo(sorter, dataGrad[:10000]))
	b.Run("grad_100k", foo(sorter, dataGrad[:100000]))
	b.Run("grad_1000k", foo(sorter, dataGrad))
}

func BenchmarkBubbleSorter(b *testing.B) {
	doBench(BubbleSorter{}, b)
}

func BenchmarkInsertionSorter(b *testing.B) {
	doBench(InsertionSorter{}, b)
}

func BenchmarkSelectionSorter(b *testing.B) {
	doBench(SelectionSorter{}, b)
}

func BenchmarkQuickSorter(b *testing.B) {
	doBench(QuickSort{}, b)
}

func BenchmarkMergeSorter(b *testing.B) {
	doBench(MergeSort{}, b)
}

func BenchmarkShellSorter(b *testing.B) {
	doBench(ShellSort{}, b)
}

func BenchmarkHeapSorter(b *testing.B) {
	doBench(HeapSort{}, b)
}
