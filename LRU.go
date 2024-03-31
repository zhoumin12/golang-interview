package main

/**
最近最少使用算法LRU
1.获取页的时候需要更新最近访问的页
2.放入新的页的时候判断是否要淘汰页面，并进行淘汰。
*/

type ListNode struct {
	val  int
	next *ListNode
	pre  *ListNode
}

type LUR struct {
	size  int
	cache map[int]*ListNode
	head  *ListNode
	tail  *ListNode
}

func (L *LUR) init(size int) {
	L.size = size
	L.cache = make(map[int]*ListNode, L.size)
}

func (L *LUR) put(key int) {
	var node *ListNode
	if v, ok := L.cache[key]; ok {
		if L.head != v {
			v.pre.next = v.next
			v.next.pre = v.pre
		}
	} else {
		node = &ListNode{val: key}
		L.cache[key] = node
	}
	L.head.pre = node
	node.next = L.head
	node.pre = nil
	L.head = node
	if len(L.cache) < L.size {
		if len(L.cache) == 0 {
			L.tail = node
		}
	} else {
		//淘汰最后一个节点
		deleteNode := L.tail
		L.tail.pre.next = nil
		delete(L.cache, deleteNode.val)
	}
}

func (L *LUR) get(key int) int {
	if v, ok := L.cache[key]; ok {
		v.pre.next = v.next
		v.next.pre = v.pre
		L.head.pre = v
		v.next = L.head
		v.pre = nil
		L.head = v
		return v.val
	} else {
		return -1
	}
}
