package slices

import "strconv"

type _String struct {}
var String _String

func (_ *_String) Append(s []string, args ...string) []string {
	return append(s, args...)
}

func (_ *_String) AppendSlice(s []string, s2 []string) []string {
	return append(s, s2...)
}

func (_ *_String) Insert(s []string, i int, args ...string) []string {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_String) InsertSlice(s []string, i int, s2 []string) []string {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_String) Cut(s []string, i, j int) ([]string, []string) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = "" // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_String) Extract(s []string, i int) (string, []string) {
	if len(s) <= i {
		crash("cannot remove at index " + strconv.Itoa(i) +  ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = ""
	s = s[:len(s)-1]

	return r, s
}

func (_ *_String) ExtractBy(s []string, f func(a string) bool) ([]string, []string) {
	extracted := []string{}
	remaining := []string{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_String) ExtractFirstBy(s []string, f func(a string) bool) ([]string, []string) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []string{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []string{}, s
}

func (_ *_String) ExtractLastBy(s []string, f func(a string) bool) ([]string, []string) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []string{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []string{}, s
}

func (_ *_String) Remove(s []string, i int) []string {
	_, ns := String.Extract(s, i)
	return ns
}

func (_ *_String) RemoveBy(s []string, f func(a string) bool) []string {
	_, remaining := String.ExtractBy(s, f)
	return remaining
}

func (_ *_String) RemoveFirstBy(s []string, f func(a string) bool) []string {
	_, remaining := String.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_String) RemoveLastBy(s []string, f func(a string) bool) []string {
	_, remaining := String.ExtractLastBy(s, f)
	return remaining
}

func (_ *_String) Push(s []string, n string) []string {
	return append(s, n)
}

func (_ *_String) Pop(s []string) (string, []string) {
	last := len(s)-1
	return s[last], s[:last]
}

func (_ *_String) Enqueue(s []string, n string) []string {
	return append([]string{n}, s...)
}

func (_ *_String) Dequeue(s []string) (string, []string) {
	return s[0], s[1:]
}
