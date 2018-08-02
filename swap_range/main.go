package main

import "fmt"

func swapRange(r []int, fm, fn, sm, sn int) []int {
	if fn < fm || sn < sm {
		return r
	}
	if lapped(fm, fn, sm, sn) {
		return r
	}
	if fm > sm {
		return swapSLRange(r, sm, sn, fm, fn)
	}
	return swapSLRange(r, fm, fn, sm, sn)
}

func swapSLRange(r []int, fm, fn, sm, sn int) []int {
	if sm-fn == 1 {
		return swapBorderRange(r, fm, fn, sm, sn)
	}

	swapBorderRange(r, fm, sm-1, sm, sn)

	a := fm + sn - sm + 1
	b := a + fn - fm
	c := b + 1
	d := sn

	return swapBorderRange(r, a, b, c, d)
}

func lapped(fm, fn, sm, sn int) bool {
	return (fm >= sm && fm <= sn) || (sm >= fm && sm <= fn)
}

func makeController(fm, fn, sm, sn int) (func() (*int, *int), func(*int) bool, func(*int, *int), func() (int, int, int, int), bool) {

	shortRangeLen := fn - fm + 1
	longRangeLen := sn - sm + 1

	var stopIndex = -1

	if fm <= sm {
		stopIndex = sn - (longRangeLen % shortRangeLen)
	} else {
		stopIndex = sm + (longRangeLen % shortRangeLen)
	}

	begin := func() (*int, *int) {
		var a, b int
		if fm <= sm {
			a, b = fm, sm
		} else {
			a, b = fn, sn
		}
		return &a, &b
	}

	stopCondition := func(i *int) bool {
		if fm <= sm {
			return *i <= stopIndex
		}
		return *i >= stopIndex
	}

	adjust := func(i, j *int) {
		if fm <= sm {
			*i = *i + 1
			*j = *j + 1
		} else {
			*i = *i - 1
			*j = *j - 1
		}
	}

	finished := longRangeLen%shortRangeLen == 0

	nextIndexs := func() (int, int, int, int) {
		if finished {
			return -1, -1, -1, -1
		}

		if fm <= sm {
			return stopIndex + 1, stopIndex + longRangeLen%shortRangeLen, stopIndex - shortRangeLen + 1, stopIndex

		} else {
			return stopIndex - longRangeLen%shortRangeLen, stopIndex - 1, stopIndex, stopIndex + shortRangeLen - 1
		}

	}

	return begin, stopCondition, adjust, nextIndexs, finished

}

func swapBorderRange(r []int, fm, fn, sm, sn int) []int {
	if fn < fm || sn < sm {
		return r
	}
	if lapped(fm, fn, sm, sn) {
		return r
	}

	if fn-fm <= sn-sm {
		return swapBorderSLRange(r, fm, fn, sm, sn)
	}
	return swapBorderSLRange(r, sm, sn, fm, fn)

}

//swapSLRange swap short long range
func swapBorderSLRange(r []int, fm, fn, sm, sn int) []int {
	if fn < fm || sn < sm {
		return r
	}
	if lapped(fm, fn, sm, sn) {
		return r
	}

	begin, stopCondition, adjust, next, finished := makeController(fm, fn, sm, sn)

	for i, j := begin(); stopCondition(j); adjust(i, j) {
		r[*i], r[*j] = r[*j], r[*i]
	}

	if finished {
		return r
	} else {
		a, b, c, d := next()
		return swapBorderSLRange(r, a, b, c, d)
	}

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(swapRange(a, 4, 4, 0, 0))
}
