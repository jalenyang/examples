package main

import (
	"fmt"
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	max := 0
	var i, j = 0, 0
	arr := ""
	for {
		if i < n && j < n {
			if !strings.Contains(arr, string(s[j])) {
				arr += string(s[j])
				j++
				max = int(math.Max(float64(max), float64(j-i)))
			} else {
				arr = arr[1 : len(arr)]
				i++
			}
		} else {
			break
		}
	}
	return max
}

func main() {
	max := lengthOfLongestSubstring("abcdefffaabbcdefghia")
	fmt.Println(max)
}
