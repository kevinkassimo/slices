package slices

import "strconv"

type _Int8 struct{}

var Int8 _Int8

func (_ *_Int8) Append(s []int8, args ...int8) []int8 {
	return append(s, args...)
}

func (_ *_Int8) AppendSlice(s []int8, s2 []int8) []int8 {
	return append(s, s2...)
}

func (_ *_Int8) Insert(s []int8, i int, args ...int8) []int8 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Int8) InsertSlice(s []int8, i int, s2 []int8) []int8 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Int8) Cut(s []int8, i, j int) ([]int8, []int8) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Int8) Extract(s []int8, i int) (int8, []int8) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Int8) ExtractBy(s []int8, f func(a int8) bool) ([]int8, []int8) {
	extracted := []int8{}
	remaining := []int8{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Int8) ExtractFirstBy(s []int8, f func(a int8) bool) ([]int8, []int8) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []int8{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int8{}, s
}

func (_ *_Int8) ExtractLastBy(s []int8, f func(a int8) bool) ([]int8, []int8) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []int8{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int8{}, s
}

func (_ *_Int8) Remove(s []int8, i int) []int8 {
	_, ns := Int8.Extract(s, i)
	return ns
}

func (_ *_Int8) RemoveBy(s []int8, f func(a int8) bool) []int8 {
	_, remaining := Int8.ExtractBy(s, f)
	return remaining
}

func (_ *_Int8) RemoveFirstBy(s []int8, f func(a int8) bool) []int8 {
	_, remaining := Int8.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Int8) RemoveLastBy(s []int8, f func(a int8) bool) []int8 {
	_, remaining := Int8.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Int8) Push(s []int8, n int8) []int8 {
	return append(s, n)
}

func (_ *_Int8) Pop(s []int8) (int8, []int8) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_Int8) Enqueue(s []int8, n int8) []int8 {
	return append([]int8{n}, s...)
}

func (_ *_Int8) Dequeue(s []int8) (int8, []int8) {
	return s[0], s[1:]
}
