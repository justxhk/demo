package math

type Sort struct {
	arr    []int
	length int
}

func SortInit(arr []int) Sort {
	return Sort{arr: arr, length: len(arr)}
}

func (s Sort) Len() int {
	return s.length
}

func (s Sort) QuickSort(left, right int) {
	if left >= right {
		return
	}
	p := s.partition(left, right)
	s.QuickSort(p, right)
	s.QuickSort(left, p)
}

func (s Sort) partition(left, right int) int {
	m := s.arr[right-1]
	p := left
	for i := left; i < right-1; i++ {
		if s.arr[i] < m {
			s.arr[i], s.arr[p] = s.arr[p], s.arr[i]
			p++
		}
	}
	s.arr[right], s.arr[p] = s.arr[p], s.arr[right]
	return p
}

func (s Sort) BubbleSort() {
	for i := 0; i < s.length-1; i++ {
		flag := true
		for j := 0; j < s.length-1-i; j++ {
			if s.arr[j] > s.arr[j+1] {
				s.arr[j], s.arr[j+1] = s.arr[j+1], s.arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func (s Sort) InsertionSort() {
	for i := 1; i < s.length; i++ {

	}
}

func (s Sort) SelectionSort() {

}
