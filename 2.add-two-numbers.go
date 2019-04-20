package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
  Val  int
  Next *ListNode
}

//no consider big int
func addTwoNumbersT1(l1 *ListNode, l2 *ListNode) *ListNode {
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

//carry add perfect
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  num1, num2, carryFlag := 0, 0, 0
  headLs := ListNode{0, nil}
  lsp := &headLs
  for l1 != nil || l2 != nil || carryFlag == 1 {
    headLs.Val++
    lsp.Next = &ListNode{}
    lsp = lsp.Next
    
    if l2 != nil {
      num2 = l2.Val
      l2 = l2.Next
    } else {
      num2 = 0
    }
    
    if l1 != nil {
      num1 = l1.Val
      l1 = l1.Next
    } else {
      num1 = 0
    }
    
    lsp.Val = (num1 + num2) % 10
    if carryFlag == 1 {
      if carryFlag+lsp.Val >= 10 {
        carryFlag = 1
        lsp.Val = 0
      } else {
        carryFlag = 0
        lsp.Val++
      }
    }
    
    if num1+num2+carryFlag >= 10 {
      carryFlag = 1 //判断进位
    } else {
      carryFlag = 0
    }
  }
  
  return headLs.Next
}

func main() {
  L1 := ListNode{
    2, &ListNode{
      4, &ListNode{
        3, nil,
      },
    },
  }
  L2 := ListNode{
    5, &ListNode{
      6, &ListNode{
        4, nil,
      },
    },
  }
  fmt.Print(addTwoNumbers(&L1, &L2))
}
