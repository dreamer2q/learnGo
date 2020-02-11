package popcount

func PopCount(x uint64) int {

	var r int
	for i := 0; i < 64; i++ {
		r += int(x >> i & 0x01)
	}

	return r
}
