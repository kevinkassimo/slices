package slices

import "strconv"

type _Bool struct {}
var Bool _Bool

func (_ *_Bool) Append(s []bool, args ...bool) []bool {
	return append(s, args...)
}

func (_ *_Bool) AppendSlice(s []bool, s2 []bool) []bool {
	return append(s, s2...)
}

func (_ *_Bool) Insert(s []bool, i int, args ...bool) []bool {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Bool) InsertSlice(s []bool, i int, s2 []bool) []bool {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Bool) Cut(s []bool, i, j int) ([]bool, []bool) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = false // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Bool) Extract(s []bool, i int) (bool, []bool) {
	if len(s) <= i {
		crash("cannot remove at index " + strconv.Itoa(i) +  ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = false
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Bool) ExtractBy(s []bool, f func(a bool) bool) ([]bool, []bool) {
	extracted := []bool{}
	remaining := []bool{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Bool) ExtractFirstBy(s []bool, f func(a bool) bool) ([]bool, []bool) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []bool{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []bool{}, s
}

func (_ *_Bool) ExtractLastBy(s []bool, f func(a bool) bool) ([]bool, []bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []bool{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []bool{}, s
}

func (_ *_Bool) Remove(s []bool, i int) []bool {
	_, ns := Bool.Extract(s, i)
	return ns
}

func (_ *_Bool) RemoveBy(s []bool, f func(a bool) bool) []bool {
	_, remaining := Bool.ExtractBy(s, f)
	return remaining
}

func (_ *_Bool) RemoveFirstBy(s []bool, f func(a bool) bool) []bool {
	_, remaining := Bool.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Bool) RemoveLastBy(s []bool, f func(a bool) bool) []bool {
	_, remaining := Bool.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Bool) Push(s []bool, n bool) []bool {
	return append(s, n)
}

func (_ *_Bool) Pop(s []bool) (bool, []bool) {
	last := len(s)-1
	return s[last], s[:last]
}

func (_ *_Bool) Enqueue(s []bool, n bool) []bool {
	return append([]bool{n}, s...)
}

func (_ *_Bool) Dequeue(s []bool) (bool, []bool) {
	return s[0], s[1:]
}
