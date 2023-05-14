package generics

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type MyNumber interface {
	Integer | Float
}

func Sum[T MyNumber](numbers ...T) T {
	var result T

	for _, number := range numbers {
		result += number
	}

	return result
}
