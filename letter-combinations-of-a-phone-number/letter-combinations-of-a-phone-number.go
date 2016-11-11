package main

import "fmt"

func letterCombinations(digits string) []string {

	ret := []string{}
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	for _, s := range digits {
		key := string(s)
		vals, ok := m[key]
		if ok == false {
			return []string{}
		}
		if len(ret) == 0 {
			ret = vals
		} else {
			new_ret := []string{}
			for _, prefix := range ret {
				for _, suffix := range vals {
					new_ret = append(new_ret, prefix+suffix)
				}
			}
			ret = new_ret
		}
	}
	return ret
}

func main() {
	fmt.Println(letterCombinations("203"))
}
