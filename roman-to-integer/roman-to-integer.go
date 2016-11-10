package main

import "fmt"

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	ret := 0
	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	strs := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	for _, k := range strs {
		if s == k {
			ret += m[k]
			break
		}
		point := len(s) - len(k)
		if point < 0 {
			continue
		}
		for string(s[point:]) == k && len(s) >= len(k) {
			ret += m[k]
			s = string(s[:point])
			point = len(s) - len(k)
			if point < 0 {
				break
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(romanToInt("MCMXCVI"))
}
