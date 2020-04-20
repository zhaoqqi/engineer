package _6_linkedlist

/*
思路1：开一个栈存放链表前半段
时间复杂度：O(N)
空间复杂度：O(N)
*/
func isPalindrome1(l *LinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	s := make([]string, 0, lLen/2)
	cur := l.head
	for i:=uint(1); i<=lLen; i++ {
		cur = cur.next
		if lLen%2 != 0 && i == (lLen/2+1) {
			continue
		}
		if i <= lLen/2 {
			s = append(s, cur.GetValue().(string))
		} else {
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}
	return true
}

/*
思路2：找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/
func isPalindrome2(l *LinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	isPalindrome = true
	step := lLen / 2
	cur := l.head.GetNext()
	var pre *ListNode = nil
	for i:=uint(1); i<=step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
	}
	mid := cur

	var left, right *ListNode = pre, nil
	if lLen%2 != 0 { // 链表长度为奇数
		right = mid.GetNext()
	} else { // 链表长度为偶数
		right = mid
	}

	for nil != left && nil != right {
		if left.GetValue.(string) != right.GetValue.(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	// 复原列表
	cur = pre
	pre = mid
	for nil != cur {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre
	return isPalindrome
}

/*
思路3:使用快慢指针法确认链表的中间节点，并且在遍历过程中反转前半段链表，反向遍历前半段链表 vs 正向遍历后半段链表 -> 检查是否回文字符串
时间复杂度：O(N)
*/



