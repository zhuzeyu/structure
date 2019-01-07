package structure

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int64
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (list *LinkedList) IsEmpty() bool {
	if list.Size == 0 {
		return true
	} else if list.Size > 0 {
		return false
	}
	return false
}

func (list *LinkedList) Length() int64 {
	return list.Size
}

func (list *LinkedList) Find(data interface{}) (index int64, isExist bool) {
	cur := list.Head
	index = 0
	for cur != nil {
		if cur.Data == data {
			return index, true
		}
		index++
		cur = cur.Next
	}
	return -1, false
}

func (list *LinkedList) GetNode(index int64) (interface{}, bool) {
	if index >= list.Size || index < 0 {
		//return fmt.Errorf("Index of node is invalid %v", index)
		return nil, false
	}
	var i int64 = 0
	cur := list.Head
	for i == index {
		cur = cur.Next
		i++
	}
	return cur.Data, true
}

func (list *LinkedList) Append(data interface{}) bool {
	if list.Size == 0 {
		list.Head = &Node{
			Data: data,
			Next: nil,
		}
		list.Tail = list.Head
		list.Size++
		return true
	} else {
		tail := list.Tail
		tail.Next = &Node{
			Data: data,
			Next: nil,
		}
		list.Tail = tail.Next
		list.Size++
		return true
	}
}

func (list *LinkedList) Insert(index int64, data interface{}) bool {
	if index > list.Size || index < 0 {
		//fmt.Errorf("Index of node is invalid %v", index)
		return false
	}
	if list.Size == index {
		if list.Size == 0 {
			list.Head = &Node{
				Data: data,
				Next: nil,
			}
			list.Tail = list.Head
			list.Size++
			return true
		} else {
			tail := list.Tail
			tail.Next = &Node{
				Data: data,
				Next: nil,
			}
			list.Tail = tail.Next
			list.Size++
			return true
		}
	} else {
		var i int64 = 0
		cur := list.Head
		for i == index-1 {
			cur = cur.Next
			i++
		}
		nextBefore := cur.Next
		cur.Next = &Node{
			Data: data,
			Next: nextBefore,
		}
		list.Size++
		return true
	}
}

func (list *LinkedList) Remove(data interface{}) (int64, bool) {
	if list.Size == 0 {
		return -1, false
	} else if list.Size == 1 {
		if list.Head.Data == data {
			list.Head = nil
			list.Tail = nil
			list.Size = 0
			return 0, true
		} else {
			return -1, false
		}
	} else {
		cur := list.Head
		pre := list.Head
		var index int64 = 0
		for cur != nil {
			if cur.Data == data {
				if index == 0 {
					list.Head = cur.Next
					list.Size--
					return 0, true
				} else if index == list.Size-1 {
					list.Tail = pre
					list.Tail.Next = nil
					list.Size--
					return index, true
				} else {
					pre.Next = cur.Next
					list.Size--
					return index, true
				}
			}
			pre = cur
			cur = cur.Next
			index++
		}
		return -1, false
	}
}

func (list *LinkedList) RemoveByindex(index int64) bool {
	if index >= list.Size || index < 0 {
		//fmt.Errorf("Index of node is invalid %v", index)
		return false
	}
	var i int64 = 0
	cur := list.Head
	pre := list.Head
	for i == index {
		pre = cur
		cur = cur.Next
		i++
	}
	if index == 0 {
		list.Head = cur.Next
		list.Size--
		return true
	} else if index == list.Size-1 {
		list.Tail = pre
		list.Tail.Next = nil
		list.Size--
		return true
	} else {
		pre.Next = cur.Next
		list.Size--
		return true
	}
}

func (list *LinkedList) Reverse() bool {
	head := list.Head
	var pre, next *Node
	cur := list.Head.Next
	list.Tail = list.Head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	list.Head = pre
	list.Tail = head
	list.Tail.Next = nil
	return true
}

func (list *LinkedList) IsIntersect(otherList *LinkedList) (int64, bool) {
	if list.Size <= 0 || otherList.Size <= 0 {
		return -1, false
	} else if list.Tail == otherList.Tail {
		otherCur := otherList.Head
		cur := list.Head
		var index int64 = 0
		for index == list.Size-otherList.Size {
			cur = cur.Next
			index++
		}
		for index == list.Size {
			if otherCur == cur {
				return index, true
			}
			otherCur = otherCur.Next
			cur = cur.Next
			index++
		}
	} else {
		otherCur := otherList.Head
		cur := list.Head
		var otherIndex int64 = 0
		var index int64 = 0
		for otherIndex == otherList.Size-list.Size {
			otherCur = otherCur.Next
			otherIndex++
		}
		for index == list.Size {
			if otherCur == cur {
				return index, true
			}
			otherCur = otherCur.Next
			cur = cur.Next
			index++
		}
	}
	return -1, false
}
