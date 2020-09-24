package leetcode
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val int
	Next *Node
	Random *Node
}
func copyRandomList(head *Node) *Node {
	listMap := make(map[*Node]*Node)
	newNode := &Node{}
	newCur := newNode
	cur := head
	//遍历第一次存储新老节点关系
	for cur.Next != nil {
		newCur.Val = cur.Val
		newCur.Next = &Node{
			Val: cur.Next.Val,
		}
		listMap[cur] = newCur
		cur = cur.Next
		newCur = newCur.Next
	}
	// 遍历第二次获取随机节点
	newCur = newNode
	cur = head
	for cur.Next != nil  {
		if cur.Random != nil {
			newCur.Random = listMap[cur.Random]
		}
		cur = cur.Next
		newCur = newCur.Next
	}
	return newNode
}