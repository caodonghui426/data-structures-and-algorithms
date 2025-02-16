package main

import "fmt"

func main() {
	testElement5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	testElement4 := &ListNode{
		Val:  4,
		Next: testElement5,
	}
	testElement3 := &ListNode{
		Val:  3,
		Next: testElement4,
	}
	testElement2 := &ListNode{
		Val:  2,
		Next: testElement3,
	}
	testElement1 := &ListNode{
		Val:  1,
		Next: testElement2,
	}
	result := reverseKGroup(testElement1, 2)
	for result != nil {
		fmt.Println(result)
		result = result.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 如果 k 为 1 ，说明不用反转；只有一个节点也不用反转
	if k == 1 || head.Next == nil {
		return head
	}
	// 左侧节点，用于最后返回
	left := head
	// 左侧节点
	for i := 0; i < k-1; i++ {
		left = left.Next
	}
	// 动态右侧节点，用于每次连接新一组的左侧节点
	right := &ListNode{Next: head}
	for head != nil {
		tmp := head
		// 判断是否满足一组
		for i := 0; i < k-1; i++ {
			if head.Next == nil {
				right.Next = tmp
				return left
			}
			head = head.Next
		}
		// 反转指针
		next := head.Next
		head.Next = nil
		tmpLeft := reverseIteration(tmp)
		right.Next = tmpLeft
		right = tmp
		head = next
	}
	return left
}

// reverseIteration 迭代
func reverseIteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}
