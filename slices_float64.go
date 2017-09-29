package slices

import "strconv"

type _Float64 struct{}

var Float64 _Float64

func (_ *_Float64) Append(s []float64, args ...float64) []float64 {
	return append(s, args...)
}

func (_ *_Float64) AppendSlice(s []float64, s2 []float64) []float64 {
	return append(s, s2...)
}

func (_ *_Float64) Insert(s []float64, i int, args ...float64) []float64 {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Float64) InsertSlice(s []float64, i int, s2 []float64) []float64 {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Float64) Cut(s []float64, i, j int) ([]float64, []float64) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Float64) Extract(s []float64, i int) (float64, []float64) {
	if len(s) <= i {
		crash("cannot extract at index " + strconv.Itoa(i) + ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Float64) ExtractBy(s []float64, f func(a float64) bool) ([]float64, []float64) {
	extracted := []float64{}
	remaining := []float64{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Float64) ExtractFirstBy(s []float64, f func(a float64) bool) ([]float64, []float64) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []float64{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []float64{}, s
}

func (_ *_Float64) ExtractLastBy(s []float64, f func(a float64) bool) ([]float64, []float64) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []float64{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []float64{}, s
}

func (_ *_Float64) Remove(s []float64, i int) []float64 {
	_, ns := Float64.Extract(s, i)
	return ns
}

func (_ *_Float64) RemoveBy(s []float64, f func(a float64) bool) []float64 {
	_, remaining := Float64.ExtractBy(s, f)
	return remaining
}

func (_ *_Float64) RemoveFirstBy(s []float64, f func(a float64) bool) []float64 {
	_, remaining := Float64.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Float64) RemoveLastBy(s []float64, f func(a float64) bool) []float64 {
	_, remaining := Float64.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Float64) Push(s []float64, n float64) []float64 {
	return append(s, n)
}

func (_ *_Float64) Pop(s []float64) (float64, []float64) {
	last := len(s) - 1
	return s[last], s[:last]
}

func (_ *_Float64) Enqueue(s []float64, n float64) []float64 {
	return append([]float64{n}, s...)
}

func (_ *_Float64) Dequeue(s []float64) (float64, []float64) {
	return s[0], s[1:]
}
