package main

func twoSum(nums []int, target int) []int {
  list := map[int]int{}
  for key, value := range nums {
    // if value >= target { //avoid case {0,0} and minus int
    
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
