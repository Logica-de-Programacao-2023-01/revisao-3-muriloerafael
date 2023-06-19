package bonus

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func mergeTwoLists(a, b *Node) *Node {
	var temp *Node

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.Value <= b.Value {
		temp = a
		temp.Next = mergeTwoLists(a.Next, b)
	} else {
		temp = b
		temp.Next = mergeTwoLists(a, b.Next)
	}

	return temp
}

func MergeTwoLists(list1 *LinkedList, list2 *LinkedList) *LinkedList {
	return &LinkedList{
		Head: mergeTwoLists(list1.Head, list2.Head),
	}
}
