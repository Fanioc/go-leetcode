package main

import "fmt"

func lengthOfLongestSubstringT1(s string) int {
  charMap := map[rune]int{}
  num, max := 0, 0
  for _, char := range s {
    _, ok := charMap[char]
    
    if ok {
      if num > max { //
        max = num
      }
      num = 1                         //recode in first
      charMap = map[rune]int{char: 1} // recode self
    } else {
      charMap[char] = 1
      num++
    }
    if num > max { //
      max = num
    }
  }
  return max
}

//回溯 5% 5%
func lengthOfLongestSubstringT2(s string) int {
  charMap := map[uint8]int{}
  num, max := 0, 0
  for index := 0; index < len(s); index++ {
    curIndex, ok := charMap[s[index]]
    if ok {
      if num > max { //
        max = num
      }
      index = curIndex + 1
      num = 1                                  //recode
      charMap = map[uint8]int{s[index]: index} // recode self
      
    } else {
      charMap[s[index]] = index
      num++
    }
    if num > max {
      max = num
    }
  }
  return max
}

//Hash 90% 42%
func lengthOfLongestSubstring(s string) int {
  charMap := map[rune]int{}
  num, max, start := 0, 0, 0
  
  for key, value := range s {
    index, ok := charMap[value]
    if ok && index >= start { //有重复的数据,排除start前面的字符
      start = index + 1
    }
    charMap[value] = key
    num = key - start + 1
    if num > max { //记录最大数
      max = num
    }
  }
  return max
}

func main() {
  fmt.Printf("%d\n", lengthOfLongestSubstring("aab"))
  fmt.Printf("%d\n", lengthOfLongestSubstring("ab"))
  fmt.Printf("%d\n", lengthOfLongestSubstring("sdfvdsf"))
}
