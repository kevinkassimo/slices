package slices

import "strconv"

type _Interface struct {}
var Interface _Interface

func (_ *_Interface) Append(s []interface{}, args ...interface{}) []interface{} {
	return append(s, args...)
}

func (_ *_Interface) AppendSlice(s []interface{}, s2 []interface{}) []interface{} {
	return append(s, s2...)
}

func (_ *_Interface) Insert(s []interface{}, i int, args ...interface{}) []interface{} {
	s = append(s, args...)
	copy(s[i+len(args):], s[i:])

	for index, val := range args {
		s[i+index] = val
	}
	return s
}

func (_ *_Interface) InsertSlice(s []interface{}, i int, s2 []interface{}) []interface{} {
	s = append(s, s2...)
	copy(s[i+len(s2):], s[i:])

	for index, val := range s2 {
		s[i+index] = val
	}
	return s
}

func (_ *_Interface) Cut(s []interface{}, i, j int) ([]interface{}, []interface{}) {
	cutSlice := s[i:j]

	copy(s[i:], s[j:])
	for k, n := len(s)-j+i, len(s); k < n; k++ {
		s[k] = 0 // or the zero value of T
	}
	s = s[:len(s)-j+i]

	return cutSlice, s
}

func (_ *_Interface) Extract(s []interface{}, i int) (interface{}, []interface{}) {
	if len(s) <= i {
		crash("cannot remove at index " + strconv.Itoa(i) +  ": slice too short")
	}

	r := s[i]

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	return r, s
}

func (_ *_Interface) ExtractBy(s []interface{}, f func(a interface{}) bool) ([]interface{}, []interface{}) {
	extracted := []interface{}{}
	remaining := []interface{}{}

	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			extracted = append(extracted, s[i])
		} else {
			remaining = append(remaining, s[i])
		}
	}
	return extracted, remaining
}

func (_ *_Interface) ExtractFirstBy(s []interface{}, f func(a interface{}) bool) ([]interface{}, []interface{}) {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return []interface{}{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []interface{}{}, s
}

func (_ *_Interface) ExtractLastBy(s []interface{}, f func(a interface{}) bool) ([]interface{}, []interface{}) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return []interface{}{s[i]}, append(s[:i], s[i+1:]...)
		}
	}
	return []interface{}{}, s
}

func (_ *_Interface) Remove(s []interface{}, i int) []interface{} {
	_, ns := Interface.Extract(s, i)
	return ns
}

func (_ *_Interface) RemoveBy(s []interface{}, f func(a interface{}) bool) []interface{} {
	_, remaining := Interface.ExtractBy(s, f)
	return remaining
}

func (_ *_Interface) RemoveFirstBy(s []interface{}, f func(a interface{}) bool) []interface{} {
	_, remaining := Interface.ExtractFirstBy(s, f)
	return remaining
}

func (_ *_Interface) RemoveLastBy(s []interface{}, f func(a interface{}) bool) []interface{} {
	_, remaining := Interface.ExtractLastBy(s, f)
	return remaining
}

func (_ *_Interface) Push(s []interface{}, n interface{}) []interface{} {
	return append(s, n)
}

func (_ *_Interface) Pop(s []interface{}) (interface{}, []interface{}) {
	last := len(s)-1
	return s[last], s[:last]
}

func (_ *_Interface) Enqueue(s []interface{}, n interface{}) []interface{} {
	return append([]interface{}{n}, s...)
}

func (_ *_Interface) Dequeue(s []interface{}) (interface{}, []interface{}) {
	return s[0], s[1:]
}
