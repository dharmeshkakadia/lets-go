package main

import "fmt"
import "strings"

func index(list []string, s string) int {
	for i, val := range list {
		if val == s {
			return i
		}
	}

	return -1
}

func contains(list []string, s string) bool {
	return index(list, s) != -1
}

func any(list []string, f func(string) bool) bool {
	for _, s := range list {
		if f(s) {
			return true
		}
	}
	return false
}

func all(list []string, f func(string) bool) bool {
	for _, s := range list {
		if !f(s) {
			return false
		}
	}
	return true
}

func filter(list []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, val := range list {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

func apply(list []string, f func(string) string) []string {
	result := make([]string, len(list))
	for i, val := range list {
		result[i] = f(val)
	}
	return result
}

func main() {
	var list = []string{"apple", "pear", "plum", "peach"}
	fmt.Println(index(list, "apple"))
	fmt.Println(contains(list, "apple"))

	fmt.Println(any(list, func(s string) bool {
		return strings.HasPrefix(s, "a")
	}))

	fmt.Println(all(list, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(filter(list, func(s string) bool {
		return strings.Contains(s, "e")
	}))

	fmt.Println(apply(list, strings.ToUpper))
}
