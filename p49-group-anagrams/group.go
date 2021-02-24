package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	m := make(map[string][]string)

	for _, s := range strs {
		a := mySort(s)
		//fmt.Println(a)
		m[a] = append(m[a], s)
	}

	//fmt.Println("m: ", m)

	for _, v := range m {
		//fmt.Println(v)
		anagrams = append(anagrams, v)
	}

	return anagrams
}

func mySort(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
