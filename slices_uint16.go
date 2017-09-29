package slices

import "strconv"

type _UInt16 struct{}

var UInt16 _UInt16

func (_ *_UInt16) Append(s []uint16, args ...uint16) []uint16 {
	return append(s, args...)
}

func (_ *_UInt16) AppendSlice(s []uint16, s2 []uint16) []uint16 {
	return append(s, s2...)
}

func (_ *_UInt16) Insert(s []uint16, i int, args ...uint16) []uint16 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt16) InsertSlice(s []uint16, i int, s2 []uint16) []uint16 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt16) Cut(s []uint16, i, j int) ([]uint16, []uint16) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_UInt16) Extract(s []uint16, i int) (uint16, []uint16) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_UInt16) ExtractBy(s []uint16, f func(a uint16) bool) ([]uint16, []uint16) {
	extracted := []uint16{}
	remaining := []uint16{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_UInt16) ExtractFirstBy(s []uint16, f func(a uint16) bool) ([]uint16, []uint16) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []uint16{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint16{}, s
}

func (_ *_UInt16) ExtractLastBy(s []uint16, f func(a uint16) bool) ([]uint16, []uint16) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []uint16{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint16{}, s
}

func (_ *_UInt16) Remove(s []uint16, i int) []uint16 {
	_, ns := UInt16.Extract(s, i)
	return ns
}

func (_ *_UInt16) RemoveBy(s []uint16, f func(a uint16) bool) []uint16 {
	_, remaining := UInt16.ExtractBy(s, f)
	return remaining
}

func (_ *_UInt16) RemoveFirstBy(s []uint16, f func(a uint16) bool) []uint16 {
	_, remaining := UInt16.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_UInt16) RemoveLastBy(s []uint16, f func(a uint16) bool) []uint16 {
	_, remaining := UInt16.ExtractLastBy(s, f)
	return remaining
}

func (_ *_UInt16) Push(s []uint16, n uint16) []uint16 {
	return append(s, n)
}

func (_ *_UInt16) Pop(s []uint16) (uint16, []uint16) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_UInt16) Enqueue(s []uint16, n uint16) []uint16 {
	return append([]uint16{n}, s...)
}

func (_ *_UInt16) Dequeue(s []uint16) (uint16, []uint16) {
	return s[0], s[1:]
}
