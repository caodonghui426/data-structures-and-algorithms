package main

import "fmt"

func main() {
	testElement3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	testElement2 := &ListNode{
		Val:  2,
		Next: testElement3,
	}
	testElement1 := &ListNode{
		Val:  1,
		Next: testElement2,
	}
	// resultRecursion := reverseListRecursion(testElement1)
	// for resultRecursion != nil {
	// 	fmt.Println(resultRecursion)
	// 	resultRecursion = resultRecursion.Next
	// }
	resultIteration := reverseIteration(testElement1)
	for resultIteration != nil {
		fmt.Println(resultIteration)
		resultIteration = resultIteration.Next
	}
}

// reverseListRecursion 递归
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
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
