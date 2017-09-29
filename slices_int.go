package slices

import "strconv"

type _Int struct{}

var Int _Int

func (_ *_Int) Append(s []int, args ...int) []int {
	return append(s, args...)
}

func (_ *_Int) AppendSlice(s []int, s2 []int) []int {
	return append(s, s2...)
}

func (_ *_Int) Insert(s []int, i int, args ...int) []int {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Int) InsertSlice(s []int, i int, s2 []int) []int {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Int) Cut(s []int, i, j int) ([]int, []int) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Int) Extract(s []int, i int) (int, []int) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Int) ExtractBy(s []int, f func(a int) bool) ([]int, []int) {
	extracted := []int{}
	remaining := []int{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Int) ExtractFirstBy(s []int, f func(a int) bool) ([]int, []int) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []int{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int{}, s
}

func (_ *_Int) ExtractLastBy(s []int, f func(a int) bool) ([]int, []int) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []int{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int{}, s
}

func (_ *_Int) Remove(s []int, i int) []int {
	_, ns := Int.Extract(s, i)
	return ns
}

func (_ *_Int) RemoveBy(s []int, f func(a int) bool) []int {
	_, remaining := Int.ExtractBy(s, f)
	return remaining
}

func (_ *_Int) RemoveFirstBy(s []int, f func(a int) bool) []int {
	_, remaining := Int.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Int) RemoveLastBy(s []int, f func(a int) bool) []int {
	_, remaining := Int.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Int) Push(s []int, n int) []int {
	return append(s, n)
}

func (_ *_Int) Pop(s []int) (int, []int) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_Int) Enqueue(s []int, n int) []int {
	return append([]int{n}, s...)
}

func (_ *_Int) Dequeue(s []int) (int, []int) {
	return s[0], s[1:]
}
