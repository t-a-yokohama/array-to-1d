package cubarray

type CubArray[T Primitive] struct {
	firstLen  int
	secondLen int
	thirdLen  int
	array     []T
}

type Primitive interface {
	bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128 | string
}

func Create[T Primitive](fst, snd, trd int) CubArray[T] {
	return CubArray[T]{firstLen: fst, secondLen: snd, thirdLen: trd, array: make([]T, fst*snd*trd)}
}

func (c CubArray[T]) At(fst, snd, trd int) *T {
	return &c.array[fst*c.secondLen*c.thirdLen+snd*c.thirdLen+trd]
}

func (c CubArray[T]) FirstLen() int {
	return c.firstLen
}

func (c CubArray[T]) SecondLen() int {
	return c.secondLen
}

func (c CubArray[T]) ThirdLen() int {
	return c.thirdLen
}
