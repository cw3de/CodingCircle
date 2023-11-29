package lisp

type PairFunc func(int, int) int

func Cons(a, b int) func(PairFunc) int {
	return func(f PairFunc) int {
		return f(a, b)
	}
}

func Car(f func(PairFunc) int) int {
	return f(func(a, b int) int {
		return a
	})
}

func Cdr(f func(PairFunc) int) int {
	return f(func(a, b int) int {
		return b
	})
}
