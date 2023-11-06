package solution

type Node struct {
	Val  int
	Key  int
	Next *Node
}

type MyHashMap struct {
	list []*Node
	size int
}

func Constructor() MyHashMap {
	return MyHashMap{
		list: make([]*Node, 10000),
		size: 10000,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	i := key % this.size

	if this.list[i] == nil {
		this.list[i] = &Node{
			Key: key,
			Val: value,
		}
		return
	}

	curr := this.list[i]

	for curr.Next != nil {
		if curr.Key == key {
			curr.Val = value
			return
		}
		curr = curr.Next
	}

	if curr.Key == key {
		curr.Val = value
		return
	}

	curr.Next = &Node{
		Key: key,
		Val: value,
	}
}

func (this *MyHashMap) Get(key int) int {
	i := key % this.size

	curr := this.list[i]

	for curr != nil {
		if curr.Key == key {
			return curr.Val
		}
		curr = curr.Next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	i := key % this.size

	curr := this.list[i]

	if curr == nil {
		return
	}

	if curr.Key == key {
		this.list[i] = curr.Next
		return
	}

	prev := curr
	curr = curr.Next

	for curr != nil {
		if curr.Key == key {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}
}
