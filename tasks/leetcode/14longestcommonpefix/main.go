package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{""}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	fmt.Println(longestCommonPrefix([]string{"a", "ac"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))

}

func longestCommonPrefix(strs []string) string {
	result := ""
	temp := ""
	strZero := strs[0]
	for i := 0; i < len(strZero); i++ {
		checker := 0
		temp = string(strZero[i])
		for _, str := range strs[1:] {
			if len(str) < i+1 {
				return result
			}

			if string(str[i]) != temp {
				checker++
			}
		}
		if checker == 0 {
			result += temp
		} else {
			return strZero[:i]
		}
	}

	return result
}

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/
