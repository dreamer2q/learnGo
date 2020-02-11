// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// for i := range pc {
	// fmt.Printf("pc[%d] = %b\t\t%d (10) = %b (2)\n", i, pc[i], i, i)
	// }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {

	r := 0
	for i := 0; i < 8; i++ {
		//	fmt.Printf("x=%b\tx>>%d*8=%b\tbyte(x>>%d*8)=%b\n", x, i, x>>(i*8), i, byte(x>>(i*8)))
		r += int(pc[byte(x>>(i*8))])
	}
	return r

	// return int(pc[byte(x>>(0*8))] +
	// pc[byte(x>>(1*8))] +
	// pc[byte(x>>(2*8))] +
	// pc[byte(x>>(3*8))] +
	// pc[byte(x>>(4*8))] +
	// pc[byte(x>>(5*8))] +
	// pc[byte(x>>(6*8))] +
	// pc[byte(x>>(7*8))])
}

//func main() {
//
// if len(os.Args[1:]) > 0 {
//
// for _, v := range os.Args[1:] {
// r, _ := strconv.ParseUint(v, 10, 64)
// fmt.Printf("%s\t%d\n", v, PopCount(r))
// }
//
// } else {
// fmt.Println("Usage:", os.Args[0], "numbers")
// return
// }
// }
//
//!-
