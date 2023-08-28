package sqarray

type SqArray[T Primitive] struct {
	rowLen int
	colLen int
	array  []T
}

type Primitive interface {
	bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128 | string
}

func Create[T Primitive](row, col int) SqArray[T] {
	return SqArray[T]{rowLen: col, colLen: row, array: make([]T, row*col)}
}

func (s SqArray[T]) At(row, col int) *T {
	return &s.array[row*s.rowLen+col]
}

func (s SqArray[T]) RowLen() int {
	return s.rowLen
}

func (s SqArray[T]) ColLen() int {
	return s.colLen
}
