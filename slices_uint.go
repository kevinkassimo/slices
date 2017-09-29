package slices

import "strconv"

type _UInt struct{}

var UInt _UInt

func (_ *_UInt) Append(s []uint, args ...uint) []uint {
	return append(s, args...)
}

func (_ *_UInt) AppendSlice(s []uint, s2 []uint) []uint {
	return append(s, s2...)
}

func (_ *_UInt) Insert(s []uint, i int, args ...uint) []uint {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt) InsertSlice(s []uint, i int, s2 []uint) []uint {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_UInt) Cut(s []uint, i, j int) ([]uint, []uint) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_UInt) Extract(s []uint, i int) (uint, []uint) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_UInt) ExtractBy(s []uint, f func(a uint) bool) ([]uint, []uint) {
	extracted := []uint{}
	remaining := []uint{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_UInt) ExtractFirstBy(s []uint, f func(a uint) bool) ([]uint, []uint) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []uint{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint{}, s
}

func (_ *_UInt) ExtractLastBy(s []uint, f func(a uint) bool) ([]uint, []uint) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []uint{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []uint{}, s
}

func (_ *_UInt) Remove(s []uint, i int) []uint {
	_, ns := UInt.Extract(s, i)
	return ns
}

func (_ *_UInt) RemoveBy(s []uint, f func(a uint) bool) []uint {
	_, remaining := UInt.ExtractBy(s, f)
	return remaining
}

func (_ *_UInt) RemoveFirstBy(s []uint, f func(a uint) bool) []uint {
	_, remaining := UInt.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_UInt) RemoveLastBy(s []uint, f func(a uint) bool) []uint {
	_, remaining := UInt.ExtractLastBy(s, f)
	return remaining
}

func (_ *_UInt) Push(s []uint, n uint) []uint {
	return append(s, n)
}

func (_ *_UInt) Pop(s []uint) (uint, []uint) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_UInt) Enqueue(s []uint, n uint) []uint {
	return append([]uint{n}, s...)
}

func (_ *_UInt) Dequeue(s []uint) (uint, []uint) {
	return s[0], s[1:]
}
