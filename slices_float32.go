package slices

import "strconv"

type _Float32 struct{}

var Float32 _Float32

func (_ *_Float32) Append(s []float32, args ...float32) []float32 {
	return append(s, args...)
}

func (_ *_Float32) AppendSlice(s []float32, s2 []float32) []float32 {
	return append(s, s2...)
}

func (_ *_Float32) Insert(s []float32, i int, args ...float32) []float32 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Float32) InsertSlice(s []float32, i int, s2 []float32) []float32 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Float32) Cut(s []float32, i, j int) ([]float32, []float32) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Float32) Extract(s []float32, i int) (float32, []float32) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Float32) ExtractBy(s []float32, f func(a float32) bool) ([]float32, []float32) {
	extracted := []float32{}
	remaining := []float32{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Float32) ExtractFirstBy(s []float32, f func(a float32) bool) ([]float32, []float32) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []float32{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []float32{}, s
}

func (_ *_Float32) ExtractLastBy(s []float32, f func(a float32) bool) ([]float32, []float32) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []float32{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []float32{}, s
}

func (_ *_Float32) Remove(s []float32, i int) []float32 {
	_, ns := Float32.Extract(s, i)
	return ns
}

func (_ *_Float32) RemoveBy(s []float32, f func(a float32) bool) []float32 {
	_, remaining := Float32.ExtractBy(s, f)
	return remaining
}

func (_ *_Float32) RemoveFirstBy(s []float32, f func(a float32) bool) []float32 {
	_, remaining := Float32.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Float32) RemoveLastBy(s []float32, f func(a float32) bool) []float32 {
	_, remaining := Float32.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Float32) Push(s []float32, n float32) []float32 {
	return append(s, n)
}

func (_ *_Float32) Pop(s []float32) (float32, []float32) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_Float32) Enqueue(s []float32, n float32) []float32 {
	return append([]float32{n}, s...)
}

func (_ *_Float32) Dequeue(s []float32) (float32, []float32) {
	return s[0], s[1:]
}
