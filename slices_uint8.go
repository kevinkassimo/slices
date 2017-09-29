package slices

import "strconv"

type _UInt8 struct{}

var UInt8 _UInt8

func (_ *_UInt8) Append(s []uint8, args ...uint8) []uint8 {
	return append(s, args...)
}

func (_ *_UInt8) AppendSlice(s []uint8, s2 []uint8) []uint8 {
	return append(s, s2...)
}

func (_ *_UInt8) Insert(s []uint8, i int, args ...uint8) []uint8 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt8) InsertSlice(s []uint8, i int, s2 []uint8) []uint8 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt8) Cut(s []uint8, i, j int) ([]uint8, []uint8) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_UInt8) Extract(s []uint8, i int) (uint8, []uint8) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_UInt8) ExtractBy(s []uint8, f func(a uint8) bool) ([]uint8, []uint8) {
	extracted := []uint8{}
	remaining := []uint8{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_UInt8) ExtractFirstBy(s []uint8, f func(a uint8) bool) ([]uint8, []uint8) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []uint8{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint8{}, s
}

func (_ *_UInt8) ExtractLastBy(s []uint8, f func(a uint8) bool) ([]uint8, []uint8) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []uint8{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint8{}, s
}

func (_ *_UInt8) Remove(s []uint8, i int) []uint8 {
	_, ns := UInt8.Extract(s, i)
	return ns
}

func (_ *_UInt8) RemoveBy(s []uint8, f func(a uint8) bool) []uint8 {
	_, remaining := UInt8.ExtractBy(s, f)
	return remaining
}

func (_ *_UInt8) RemoveFirstBy(s []uint8, f func(a uint8) bool) []uint8 {
	_, remaining := UInt8.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_UInt8) RemoveLastBy(s []uint8, f func(a uint8) bool) []uint8 {
	_, remaining := UInt8.ExtractLastBy(s, f)
	return remaining
}

func (_ *_UInt8) Push(s []uint8, n uint8) []uint8 {
	return append(s, n)
}

func (_ *_UInt8) Pop(s []uint8) (uint8, []uint8) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_UInt8) Enqueue(s []uint8, n uint8) []uint8 {
	return append([]uint8{n}, s...)
}

func (_ *_UInt8) Dequeue(s []uint8) (uint8, []uint8) {
	return s[0], s[1:]
}
