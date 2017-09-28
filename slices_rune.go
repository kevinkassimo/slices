package slices

import "strconv"

type _Rune struct {}
var Rune _Rune

func (_ *_Rune) Append(s []rune, args ...rune) []rune {
	return append(s, args...)
}

func (_ *_Rune) AppendSlice(s []rune, s2 []rune) []rune {
	return append(s, s2...)
}

func (_ *_Rune) Insert(s []rune, i int, args ...rune) []rune {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Rune) InsertSlice(s []rune, i int, s2 []rune) []rune {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Rune) Cut(s []rune, i, j int) ([]rune, []rune) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Rune) Extract(s []rune, i int) (rune, []rune) {
	if len(s) <= i {
		crash("cannot remove at index " + strconv.Itoa(i) +  ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Rune) ExtractBy(s []rune, f func(a rune) bool) ([]rune, []rune) {
	extracted := []rune{}
	remaining := []rune{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Rune) ExtractFirstBy(s []rune, f func(a rune) bool) ([]rune, []rune) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []rune{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []rune{}, s
}

func (_ *_Rune) ExtractLastBy(s []rune, f func(a rune) bool) ([]rune, []rune) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []rune{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []rune{}, s
}

func (_ *_Rune) Remove(s []rune, i int) []rune {
	_, ns := Rune.Extract(s, i)
	return ns
}

func (_ *_Rune) RemoveBy(s []rune, f func(a rune) bool) []rune {
	_, remaining := Rune.ExtractBy(s, f)
	return remaining
}

func (_ *_Rune) RemoveFirstBy(s []rune, f func(a rune) bool) []rune {
	_, remaining := Rune.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Rune) RemoveLastBy(s []rune, f func(a rune) bool) []rune {
	_, remaining := Rune.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Rune) Push(s []rune, n rune) []rune {
	return append(s, n)
}

func (_ *_Rune) Pop(s []rune) (rune, []rune) {
	last := len(s)-1
	return s[last], s[:last]
}

func (_ *_Rune) Enqueue(s []rune, n rune) []rune {
	return append([]rune{n}, s...)
}

func (_ *_Rune) Dequeue(s []rune) (rune, []rune) {
	return s[0], s[1:]
}
