package slices

import "strconv"

type _Int32 struct{}

var Int32 _Int32

func (_ *_Int32) Append(s []int32, args ...int32) []int32 {
	return append(s, args...)
}

func (_ *_Int32) AppendSlice(s []int32, s2 []int32) []int32 {
	return append(s, s2...)
}

func (_ *_Int32) Insert(s []int32, i int, args ...int32) []int32 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Int32) InsertSlice(s []int32, i int, s2 []int32) []int32 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Int32) Cut(s []int32, i, j int) ([]int32, []int32) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Int32) Extract(s []int32, i int) (int32, []int32) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Int32) ExtractBy(s []int32, f func(a int32) bool) ([]int32, []int32) {
	extracted := []int32{}
	remaining := []int32{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Int32) ExtractFirstBy(s []int32, f func(a int32) bool) ([]int32, []int32) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []int32{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int32{}, s
}

func (_ *_Int32) ExtractLastBy(s []int32, f func(a int32) bool) ([]int32, []int32) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []int32{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []int32{}, s
}

func (_ *_Int32) Remove(s []int32, i int) []int32 {
	_, ns := Int32.Extract(s, i)
	return ns
}

func (_ *_Int32) RemoveBy(s []int32, f func(a int32) bool) []int32 {
	_, remaining := Int32.ExtractBy(s, f)
	return remaining
}

func (_ *_Int32) RemoveFirstBy(s []int32, f func(a int32) bool) []int32 {
	_, remaining := Int32.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Int32) RemoveLastBy(s []int32, f func(a int32) bool) []int32 {
	_, remaining := Int32.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Int32) Push(s []int32, n int32) []int32 {
	return append(s, n)
}

func (_ *_Int32) Pop(s []int32) (int32, []int32) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_Int32) Enqueue(s []int32, n int32) []int32 {
	return append([]int32{n}, s...)
}

func (_ *_Int32) Dequeue(s []int32) (int32, []int32) {
	return s[0], s[1:]
}
