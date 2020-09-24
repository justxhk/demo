package math

type LinkedList struct {
	Data int
	Next *LinkedList
}

// 初始化链表
func LinkedListBuilder(arr ...int) *LinkedList {
	linkedList := &LinkedList{
		Data: arr[0],
	}
	curNode := linkedList
	for i := 1; i < len(arr); i++ {
		curNode.Next = &LinkedList{
			Data: arr[i],
		}
		curNode = curNode.Next
	}
	return linkedList
}
