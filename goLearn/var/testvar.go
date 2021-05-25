package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	fmt.Printf("%T", a)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("star ....")
	var com *ListNode

	if l1.Val < l2.Val {
		com.Val = l1.Val
		if l1.Next==nil {
			com.Next=l2
			return com
		}
		com.Next = mergeTwoLists(l1.Next, l2)

	} else {
		com.Val = l2.Val
		if l2.Next==nil {
			com.Next=l1
			return com
		}
		com.Next = mergeTwoLists(l1, l2.Next)
	}
	return com
}

func listNode() {

}
