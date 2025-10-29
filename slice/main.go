package main

import "fmt"

func f(s []int) []int {
	return append(s, 10)
}

func main() {
	xx := make([]int, 3, 5)
	xx = f(xx)
	yy := append(xx[:5])
	fmt.Println(yy)

	// xx := []int{1, 2, 3, 4, 5}
	// yy := make([]int, 2)
	// copy(yy, xx)
	// fmt.Println(yy)

	// var xx [5]int
	// fmt.Println(xx, len(xx), cap(xx))

	// const x = 5
	// a := [x]int{}
	// fmt.Println(a, len(a), cap(a))

	// b := []int{}
	// b = append(b, 1, 2, 3, 4)
	// fmt.Println(b, len(b), cap(b))

	// c := append(b[1:3], 10)
	// fmt.Println(c, len(c), cap(c))
	// c[0] = 777
	// fmt.Println(b, len(b), cap(b), c, len(c), cap(c))

	// var d []int
	// d = append(d, 1)
	// fmt.Println(d)

	// e := new([]int)
	// *e = append(*e, 1)
	// fmt.Println(e)

	// var kk []int
	// fmt.Println(kk == nil)

	// var a []int
	// b := []int{}
	// c := []int(nil)

	// d := make([]int, 0)
	// e := make([]int, 3, 5)

	// f := new([]int)
	// g := append(a)

	// fmt.Println(a, len(a), cap(a), a == nil)
	// fmt.Println(b, len(b), cap(b), b == nil)
	// fmt.Println(c, len(c), cap(c), c == nil)
	// fmt.Println(d, len(d), cap(d), d == nil)
	// fmt.Println(e, len(e), cap(e), e == nil)
	// fmt.Println(*f, len(*f), cap(*f), *f == nil)
	// fmt.Println(g, len(g), cap(g), g == nil)
}
