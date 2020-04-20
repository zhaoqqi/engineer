package _7_linkedlist

/*
单链表反转
时间复杂度：O(N)
*/
func (this *LinkedList) Reverse() *LinkedList {
	if nil == this.head.next {
		return nil
	}
	if this.length == 1 {
		return this
	}
	var pre *LinkedNode = nil
	cur = this.head.GetNext()
	next = this.head.GetNext().GetNext()
	for nil != cur {
		tmp := cur.GetNext()
		if nil == tmp {
			break
		}
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}
	this.head.next = pre
	return this
}

/*
链表中环的检测
思路1：使用快慢两个指针顺序遍历单链表，如果未相遇且都遍历结束则单链表没有环；如果有环则快慢指针必相遇。
*/
func (this *LinkedList) HasRing1() bool {
	if nil == this.head.next || this.length == 1 {
		return false
	}
	slow := this.head
	quick := this.head
	for slow != nil && quick != nil {
		slow = slow.GetNext()
		quick = quick.GetNext().GetNext()
		if slow == quick {
			return true
		}
	}
	return false
}


/*
链表中环的检测
思路2：足迹法，遍历单链表且使用散列表记录出现过的节点，如果有节点出现两次则说明有环。只能假设值重复的情况。
*/
func (this *LinkedList) HadRing2() bool {
	if nil == this.head.next || this.length == 1 {
		return false
	}
	hashMap := make(map[interface{}]int)
	cur := this.head.GetNext()
	index := 0
	for nil != cur {
		if _, ok := hashMap[cur.GetValue().(string)]; ok {
			return true
		} else {
			hashMap[cur.GetValue().(string)] = index++
		}
	}
	return false
}

/*
两个有序的链表合并
时间复杂度：O(n)
空间复杂度：O(n)
*/
func (this *LinkedList) MergeSortedList(l1, l2 *LinkedList) {
	if nil == l1 || nil == l1.head || nil == l1 {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := NewLinkedList{head: &ListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for nil != curl1 && nil != curl2 {
		if curl1.GetValue.(int) > curl2.GetValue.(int) {
			cur.next = curl2
		} else {
			cur.next = curl1
		}
		curl1 = curl1.GetNext()
		curl2 = curl2.GetNext()
	}
	if nil != curl1 {
		cur.next = curl1
	} else if nil != curl2 {
		cur.next = curl2
	}

	return l
}

/*
删除链表倒数第n个节点
思路：删除倒数第n个节点，就是删除正数第 l.length-n+1 个节点，遍历链表到 l.length-n+1 到位置，删除该节点
*/
func (list *LinkedList) DeleteBottomN(n int) {

}

/*
删除链表倒数第n个节点
思路：王争到思路（挺巧妙到）,fast先遍历n次后，slow和fast再同时遍历直到fast为nil，这时slow称为目标节点到pre,执行删除操作即可
*/
func (list *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == list.head || nil == list.head.next {
		return
	}
	fast := list.head
	for i := 1; i <= n && nil != fast; i++ {
		fast = fast.next
	}
	slow := this.head
	for nil != fast {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

/*
求链表的中间节点
*/
func (this *LinkedList) FindMiddleNode() *ListNode {
	if nil == this.head || nil == this.head.next {
		return nil
	}
	if nil == this.head.next.next {
		return this.head.next
	}
	slow, fast := this.head, this.head
	for nil != slow && nil != fast {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

