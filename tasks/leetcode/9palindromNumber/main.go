package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(123))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arrInt := convertIntArr(x)
	result := checkPalindrome(arrInt)
	return result
}

func checkPalindrome(arrInt []int) bool {
	lenHalf := len(arrInt) / 2
	lenFull := len(arrInt)

	for i := 0; i < lenFull; i++ {
		if i > lenHalf {
			break
		}
		if arrInt[i] != arrInt[lenFull-i-1] {
			return false
		}
	}
	return true
}

func convertIntArr(x int) []int {
	var arrInt []int
	numb := x
	for numb > 0 {
		temp := numb % 10
		numb /= 10
		arrInt = append(arrInt, temp)
	}
	return arrInt
}
