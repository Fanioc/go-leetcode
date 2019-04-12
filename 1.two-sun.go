package main

func twoSum(nums []int, target int) []int {
  list := map[int]int{}
  for key, value := range nums {
    // if value >= target { //避免0,0情况  不能添加=号
    
    left := target - value
    index, ok := list[left]
    if ok {
      return []int{index, key}
    } else {
      list[value] = key
    }
    
  }
  return []int{0, 0}
}
