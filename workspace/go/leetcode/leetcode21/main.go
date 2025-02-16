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
		Next: nil,
	}
	testElement1 := &ListNode{
		Val:  1,
		Next: testElement2,
	}
	result := mergeTwoLists(testElement1, testElement3)
	for result != nil {
		fmt.Println(result)
		result = result.Next
	}
}

// mergeTwoLists 合并两个有序链表(升序)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	resList := curr
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			next := p1.Next
			curr.Next = p1
			p1 = next
			curr = curr.Next
			continue
		}
		next := p2.Next
		curr.Next = p2
		p2 = next
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	}
	if p2 != nil {
		curr.Next = p2
	}
	return resList.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
