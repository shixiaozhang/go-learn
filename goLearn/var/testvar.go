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
		com.Next = l1.Next
	
	} else {
		com.Val = l2.Val
		com.Next = l2.Next
		
	}
	
	return com
}

func listNode()  {
	
}