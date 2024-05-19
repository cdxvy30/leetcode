package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int

	for i := 0; i < len(s); i++ {
		count[s[i]-'a'] += 1
		count[t[i]-'a'] -= 1
	}

	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	testcases := []struct {
		s string
		t string
	}{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"ab", "a"},
		{"aacc", "ccac"},
		{"a", "b"},
		{"a", "a"},
	}

	for _, testcase := range testcases {
		fmt.Printf("isAnagram(%v, %v) = %v\n", testcase.s, testcase.t, isAnagram(testcase.s, testcase.t))
	}
}
