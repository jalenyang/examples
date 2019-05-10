package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make(map[int][]string)
	upRow := 0
	downRow := numRows - 1
	for _, v := range s {
		if upRow < numRows {
			res[upRow] = append(res[upRow], string(v))
			upRow++
			if upRow == numRows {
				if numRows == 2 {
					upRow = 0
				} else {
					downRow = numRows - 1
				}
			}
		} else if downRow > 1 {
			downRow--
			res[downRow] = append(res[downRow], string(v))
			if downRow == 1 {
				upRow = 0
			}
		}
	}
	var str string
	for i := 0; i < len(res); i++ {
		str += strings.Join(res[i], "")
	}
	return str
}

func main() {
	fmt.Printf(convert("PAYPALISHIRING", 2))
}
