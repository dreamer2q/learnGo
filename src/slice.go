package main

import (
	"fmt"
)

func main() {

	//create slice
	var (
		a []int               //nil 0 0
		b = []int{}           //{} 0 0
		c = []int{1, 2, 3}    //{1,2,3} 3 3
		d = c[:2]             //{1,2} 2 3
		e = c[0:2:cap(c)]     //{1,2} 2 3
		f = c[:0]             //{} 0 3
		g = make([]int, 3)    //{0,0,0} 3 3
		h = make([]int, 2, 3) //{0,0} 2 3
		i = make([]int, 0, 3) //{} 0 3
	)

	fmt.Printf("%#v %d %d\n", a, len(a), cap(a))
	fmt.Printf("%#v %d %d\n", b, len(b), cap(b))
	fmt.Printf("%#v %d %d\n", c, len(c), cap(c))
	fmt.Printf("%#v %d %d\n", d, len(d), cap(d))
	fmt.Printf("%#v %d %d\n", e, len(e), cap(e))
	fmt.Printf("%#v %d %d\n", f, len(f), cap(f))
	fmt.Printf("%#v %d %d\n", g, len(g), cap(g))
	fmt.Printf("%#v %d %d\n", h, len(h), cap(h))
	fmt.Printf("%#v %d %d\n", i, len(i), cap(i))

	//visit slice
	for i := range a {
		fmt.Printf("a[%d]:%d ", i, a[i])
	}
	for i, v := range c {
		fmt.Printf("c[%d]:%d ", i, v)
	}
	fmt.Printf("\n")
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]:%d ", i, c[i])
	}
	fmt.Printf("\n")

	//add elememts to slice
	a = append(a, 1)
	a = append(a, 2, 3)
	a = append(a, []int{4, 5, 6}...)
	fmt.Printf("a %#v\n", a)
	a = append([]int{-4, -3, -2, -1}, a...)
	fmt.Printf("a %#v\n", a)
	a = append(a[:4], append([]int{0}, a[4:]...)...)
	fmt.Printf("a %#v\n", a)

	a = append(a, 666)
	copy(a[4+1:], a[4:])
	a[4] = 666
	fmt.Printf("a %#v\n", a)

	//del elements of slice
	a = a[:len(a)-1]
	fmt.Printf("%#v\n", a)
	a = a[1:]
	fmt.Printf("%#v\n", a)
	a = append(a[:3], a[4:]...)
	fmt.Printf("%#v\n", a)
}
