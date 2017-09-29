package slices

import "strconv"

type _UInt32 struct{}

var UInt32 _UInt32

func (_ *_UInt32) Append(s []uint32, args ...uint32) []uint32 {
	return append(s, args...)
}

func (_ *_UInt32) AppendSlice(s []uint32, s2 []uint32) []uint32 {
	return append(s, s2...)
}

func (_ *_UInt32) Insert(s []uint32, i int, args ...uint32) []uint32 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt32) InsertSlice(s []uint32, i int, s2 []uint32) []uint32 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt32) Cut(s []uint32, i, j int) ([]uint32, []uint32) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_UInt32) Extract(s []uint32, i int) (uint32, []uint32) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_UInt32) ExtractBy(s []uint32, f func(a uint32) bool) ([]uint32, []uint32) {
	extracted := []uint32{}
	remaining := []uint32{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_UInt32) ExtractFirstBy(s []uint32, f func(a uint32) bool) ([]uint32, []uint32) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []uint32{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint32{}, s
}

func (_ *_UInt32) ExtractLastBy(s []uint32, f func(a uint32) bool) ([]uint32, []uint32) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []uint32{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint32{}, s
}

func (_ *_UInt32) Remove(s []uint32, i int) []uint32 {
	_, ns := UInt32.Extract(s, i)
	return ns
}

func (_ *_UInt32) RemoveBy(s []uint32, f func(a uint32) bool) []uint32 {
	_, remaining := UInt32.ExtractBy(s, f)
	return remaining
}

func (_ *_UInt32) RemoveFirstBy(s []uint32, f func(a uint32) bool) []uint32 {
	_, remaining := UInt32.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_UInt32) RemoveLastBy(s []uint32, f func(a uint32) bool) []uint32 {
	_, remaining := UInt32.ExtractLastBy(s, f)
	return remaining
}

func (_ *_UInt32) Push(s []uint32, n uint32) []uint32 {
	return append(s, n)
}

func (_ *_UInt32) Pop(s []uint32) (uint32, []uint32) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_UInt32) Enqueue(s []uint32, n uint32) []uint32 {
	return append([]uint32{n}, s...)
}

func (_ *_UInt32) Dequeue(s []uint32) (uint32, []uint32) {
	return s[0], s[1:]
}
