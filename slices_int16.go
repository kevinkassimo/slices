package slices

import "strconv"

type _Int16 struct{}

var Int16 _Int16

func (_ *_Int16) Append(s []int16, args ...int16) []int16 {
	return append(s, args...)
}

func (_ *_Int16) AppendSlice(s []int16, s2 []int16) []int16 {
	return append(s, s2...)
}

func (_ *_Int16) Insert(s []int16, i int, args ...int16) []int16 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Int16) InsertSlice(s []int16, i int, s2 []int16) []int16 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Int16) Cut(s []int16, i, j int) ([]int16, []int16) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Int16) Extract(s []int16, i int) (int16, []int16) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Int16) ExtractBy(s []int16, f func(a int16) bool) ([]int16, []int16) {
	extracted := []int16{}
	remaining := []int16{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Int16) ExtractFirstBy(s []int16, f func(a int16) bool) ([]int16, []int16) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []int16{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int16{}, s
}

func (_ *_Int16) ExtractLastBy(s []int16, f func(a int16) bool) ([]int16, []int16) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []int16{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int16{}, s
}

func (_ *_Int16) Remove(s []int16, i int) []int16 {
	_, ns := Int16.Extract(s, i)
	return ns
}

func (_ *_Int16) RemoveBy(s []int16, f func(a int16) bool) []int16 {
	_, remaining := Int16.ExtractBy(s, f)
	return remaining
}

func (_ *_Int16) RemoveFirstBy(s []int16, f func(a int16) bool) []int16 {
	_, remaining := Int16.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Int16) RemoveLastBy(s []int16, f func(a int16) bool) []int16 {
	_, remaining := Int16.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Int16) Push(s []int16, n int16) []int16 {
	return append(s, n)
}

func (_ *_Int16) Pop(s []int16) (int16, []int16) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_Int16) Enqueue(s []int16, n int16) []int16 {
	return append([]int16{n}, s...)
}

func (_ *_Int16) Dequeue(s []int16) (int16, []int16) {
	return s[0], s[1:]
}
