package main

func findMedianSortedArraysT1(nums1 []int, nums2 []int) float64 {
  len1, len2 := len(nums1), len(nums2)
  m, n, sum := 0, 0, 0.0
  if (len1+len2)&0x1 == 1 {
    //奇数列
    for i := 0; i < (len1+len2)/2+1; i++ {
      if nums1[m] < nums2[n] {
        sum = float64(nums2[n])
        if m < len1-1 {
          m++
        }
      } else {
        sum = float64(nums1[m])
        if n < len2-1 {
          n++
        }
      }
    }
    return sum
  } else {
    //偶数列
    m, n := 0, 0
    for i := 0; i <= (len1+len2)/2+1; i++ {
      sum = (float64(nums2[n]) + float64(nums1[m])) / 2.0
      if nums1[m] < nums2[n] {
        if m < len1-1 {
          m++
        }
      } else {
        if n < len2-1 {
          n++
        }
      }
    }
    return sum
  }
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  len1, len2 := len(nums1), len(nums2)
  m, n, sum := 0, 0, 0.0
  
  for i := 0; i <= (len1+len2)/2+1; i++ {
    if len1 > m && len2 > n {
      if nums1[m] < nums2[n] {
        if len1-m > 1 {
          m++
        }else{
        
        }
        m++
      } else {
        if len2-n > 1 {
          m++
        }
      }
    } else if len1 > m && len2 <= n { //第二条数组没有数了
    
    } else if len1 <= m && len2 > n {
    
    }
  }
  
}
