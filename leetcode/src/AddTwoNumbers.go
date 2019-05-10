package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	dummyHead := ListNode{}
	curr := &dummyHead
	for {
		if l1 != nil || l2 != nil {
			v1 := 0
			v2 := 0
			if l1 != nil {
				v1 = l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				v2 = l2.Val
				l2 = l2.Next
			}
			val := v1 + v2 + flag
			flag = val / 10
			curr.Next = &ListNode{Val: val % 10}
			curr = curr.Next
		} else {
			break
		}
	}
	if flag > 0 {
		curr.Next = &ListNode{Val: flag}
	}
	return dummyHead.Next
}

func main() {
	n3 := ListNode{
		Val:  3,
		Next: nil,
	}
	n2 := ListNode{
		Val:  4,
		Next: &n3,
	}
	n1 := ListNode{
		Val:  2,
		Next: &n2,
	}

	n6 := ListNode{
		Val:  4,
		Next: nil,
	}
	n5 := ListNode{
		Val:  6,
		Next: &n6,
	}
	n4 := ListNode{
		Val:  5,
		Next: &n5,
	}
	addTwoNumbers(&n1, &n4)
	fmt.Println(12 / 10)
}
