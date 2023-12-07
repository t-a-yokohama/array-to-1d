package nomsquare

type NomSquare[T Primitive] struct {
	rowLen int
	colLen int
	array  [][]T
}

type Primitive interface {
	bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128 | string
}

func Create[T Primitive](row, col int) NomSquare[T] {
	arr := make([][]T, row)
	for i := range arr {
		arr[i] = make([]T, col)
	}
	return NomSquare[T]{rowLen: col, colLen: row, array: arr}
}

func (s NomSquare[T]) At(row, col int) *T {
	return &s.array[row][col]
}

func (s NomSquare[T]) RowLen() int {
	return s.rowLen
}

func (s NomSquare[T]) ColLen() int {
	return s.colLen
}
