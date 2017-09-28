package slices

import "strconv"

type _Byte struct {}
var Byte _Byte

func (_ *_Byte) Append(s []byte, args ...byte) []byte {
	return append(s, args...)
}

func (_ *_Byte) AppendSlice(s []byte, s2 []byte) []byte {
	return append(s, s2...)
}

func (_ *_Byte) Insert(s []byte, i int, args ...byte) []byte {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Byte) InsertSlice(s []byte, i int, s2 []byte) []byte {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Byte) Cut(s []byte, i, j int) ([]byte, []byte) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Byte) Extract(s []byte, i int) (byte, []byte) {
	if len(s) <= i {
		crash("cannot remove at index " + strconv.Itoa(i) +  ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Byte) ExtractBy(s []byte, f func(a byte) bool) ([]byte, []byte) {
	extracted := []byte{}
	remaining := []byte{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Byte) ExtractFirstBy(s []byte, f func(a byte) bool) ([]byte, []byte) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []byte{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []byte{}, s
}

func (_ *_Byte) ExtractLastBy(s []byte, f func(a byte) bool) ([]byte, []byte) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []byte{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []byte{}, s
}

func (_ *_Byte) Remove(s []byte, i int) []byte {
	_, ns := Byte.Extract(s, i)
	return ns
}

func (_ *_Byte) RemoveBy(s []byte, f func(a byte) bool) []byte {
	_, remaining := Byte.ExtractBy(s, f)
	return remaining
}

func (_ *_Byte) RemoveFirstBy(s []byte, f func(a byte) bool) []byte {
	_, remaining := Byte.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Byte) RemoveLastBy(s []byte, f func(a byte) bool) []byte {
	_, remaining := Byte.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Byte) Push(s []byte, n byte) []byte {
	return append(s, n)
}

func (_ *_Byte) Pop(s []byte) (byte, []byte) {
	last := len(s)-1
	return s[last], s[:last]
}

func (_ *_Byte) Enqueue(s []byte, n byte) []byte {
	return append([]byte{n}, s...)
}

func (_ *_Byte) Dequeue(s []byte) (byte, []byte) {
	return s[0], s[1:]
}
