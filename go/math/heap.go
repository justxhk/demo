package math

type heap struct {
	data []int
	cap  int
	len  int
}

func BuildMaxHeap(arr []int) {
	l := len(arr)
	for i := l / 2; i >= 0; i-- {
		Heapify(arr, i, l)
	}
}
func Heapify(arr []int, i int, len int) {
	for {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		maxPos := i
		if leftIndex < len && arr[maxPos] < arr[leftIndex] {
			maxPos = leftIndex
		}
		if rightIndex < len && arr[maxPos] < arr[rightIndex] {
			maxPos = rightIndex
		}
		if maxPos == i {
			break
		}
		swap(arr, i, maxPos)
		i = maxPos
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
