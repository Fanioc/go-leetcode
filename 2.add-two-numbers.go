package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
  Val  int
  Next *ListNode
}

//考虑大整数的问题
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  num1, num2, i := 0, 0, 1
  for l1 != nil {
    num1 = num1 + l1.Val*i
    i *= 10
    l1 = l1.Next
  }
  //fmt.Printf("num1:%d\n", num1)
  i = 1
  for l2 != nil {
    num2 = num2 + l2.Val*i
    i *= 10
    l2 = l2.Next
  }
  //fmt.Printf("num2:%d\n", num2)
  num3 := num1 + num2
  //fmt.Printf("num3:%d\n", num3)
  ls := ListNode{0, nil}
  lsp := &ls
  for num3 != 0 {
    lsp.Val = num3 % 10
    //fmt.Printf("currNodeVal:%d\n", lsp.Val)
    num3 /= 10
    if num3 != 0 {
      lsp.Next = &ListNode{}
      lsp = lsp.Next
    } else {
      break
    }
  }
  return &ls
}

func main() {
  L1 := ListNode{
    1, &ListNode{
      8, nil,
    },
  }
  L2 := ListNode{
    0, nil,
  }
  fmt.Print(addTwoNumbers(&L1, &L2))
}
