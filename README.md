# Slices for Go

`slices` is a small library for common slice manipulations.
The aim of this library is just __improving readabiliy.__

## Example
```go
import (
	//...
	github.com/kevinkassimo/slices
)

intSlice := []int{1, 2, 3}
intSlice = slices.Int.Append(intSlice, 4, 5) // [1, 2, 3, 4, 5]
newSlice, intSlice := slices.Int.ExtractBy(intSlice, func (a int) bool {
	return a < 3
}) // [1, 2], [3, 4, 5]
val, intSlice := slices.Int.Dequeue(intSlice) // 3, [4, 5]
```