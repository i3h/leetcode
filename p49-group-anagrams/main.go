package main

import "fmt"

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("str: ", str)
	anagrams := groupAnagrams(str)
	fmt.Println("anagrams: ", anagrams)
}
