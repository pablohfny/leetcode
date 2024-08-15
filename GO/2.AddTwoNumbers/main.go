package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	sum := 0

// 	if l1 != nil && l2 != nil {
// 		sum = l1.Val + l2.Val
// 	} else if l1 != nil {
// 		sum = l1.Val
// 	} else {
// 		sum = l2.Val
// 	}

// 	newList := &ListNode{
// 		Val:  sum,
// 		Next: nil,
// 	}

// 	carry := 0

// 	if sum > 9 {
// 		newList.Val = sum - 10
// 		carry = 1
// 	}

// 	if l1 != nil && l1.Next != nil && l2 != nil && l2.Next != nil {
// 		l1.Next.Val += carry
// 		newList.Next = addTwoNumbers(l1.Next, l2.Next)
// 	} else if l1 != nil && l1.Next != nil {
// 		l1.Next.Val += carry
// 		newList.Next = addTwoNumbers(l1.Next, nil)
// 	} else if l2 != nil && l2.Next != nil {
// 		l2.Next.Val += carry
// 		newList.Next = addTwoNumbers(l2.Next, nil)
// 	} else if carry > 0 {
// 		newList.Next = &ListNode{
// 			Val:  carry,
// 			Next: nil,
// 		}
// 	}

// 	return newList
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	aux := result
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
            l2 = l2.Next
		}

        if carry > 0 {
            sum += carry
        }

		carry = sum / 10
		aux.Next = &ListNode{
			Val: sum % 10,
		}
		aux = aux.Next
	}

	return result.Next
}

