package main

import "fmt"

func abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}

func reverse(x int) int {
	var l []int
	x_abs := abs(x) 
	flag := x_abs != x
	x = x_abs

    for x > 0{
	   l = append(l, x%10)
	   x /= 10
    }
	ret := 0
	for i, e := range l{
		if e >  0 || i > 0{
			f := 1
			for j:=0; j< len(l) -i -1; j++{
				f *= 10
			}
			ret += e * f
		}
	}
	if flag{
		return -ret
	}else{
		return ret
	}
}

func main() {
	x := 10
	fmt.Println(reverse(x))
}
