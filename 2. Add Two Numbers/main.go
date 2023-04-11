package main

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

	initialNode := &ListNode{}

	var sum int = 0
	for actualNode := initialNode; l1 != nil || l2 != nil || sum > 0; actualNode = actualNode.Next {

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newNode := &ListNode{}

		if sum >= 10 {
			newNode.Val = sum % 10
			sum = sum / 10
		} else {
			newNode.Val = sum
			sum = 0
		}
		actualNode.Next = newNode
	}

	return initialNode.Next

}

func addTwoNumbersBaseline(l1 *ListNode, l2 *ListNode) *ListNode {

	initialNode := &ListNode{}
	actualNode := initialNode
	var sum int = 0
	for true {
		if l1 == nil && l2 == nil {
			if sum > 0 {
				actualNode.Next = &ListNode{
					Val: sum,
				}
			}
			break
		}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newNode := &ListNode{}

		if sum >= 10 {
			newNode.Val = sum % 10
			sum = sum / 10
		} else {
			newNode.Val = sum
			sum = 0
		}

		actualNode.Next = newNode
		actualNode = newNode
	}

	return initialNode.Next

}

func main() {

}
