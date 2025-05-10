package functionalgo

type (
	Tuple[A any, B any] struct {
		fst A
		snd B
	}
	GuardS[T any] struct {
		cond bool
		fn   func() T
	}
	GeneratorS struct {
	}
)

func Guard[T any](cond bool, fn func() T) GuardS[T] {
	return GuardS[T]{
		cond: cond,
		fn:   fn,
	}
}

func Guards[T any](guards ...GuardS[T]) T {
	for _, v := range guards {
		if v.cond {
			return v.fn()
		}
	}
	panic("not exhaustive guards")
}

func Take[A ~[]any](src A, num int) A {
	return Guards(
		Guard(num > len(src), func() A { return src }),
		Guard(true, func() A { return src[:num] }),
	)
}

func Drop[A ~[]any](src A, num int) A {
	return Guards(
		Guard(num > len(src), func() A { return src[:0] }),
		Guard(true, func() A { return src[num+1:] }),
	)
}

func Head[A any, B ~[]A](src B) A {
	return Guards(
		Guard(len(src) == 0, func() A { panic("cannot take head of empty list") }),
		Guard(true, func() A { return src[0] }),
	)
}

func Tail[A any, B ~[]A](src B) B {
	return Guards(
		Guard(len(src) == 0, func() B { panic("cannot take head of empty list") }),
		Guard(true, func() B { return src[1:] }),
	)
}

func Map[A any, B any](fn func(A) B, src []A) (result []B) {
	return Foldl(func(b []B, a A) []B {
		return append(b, fn(a))
	}, []B{}, src)
}

func Filter[A any](fn func(A) bool, src []A) (result []A) {
	return Foldl(func(b []A, a A) []A {
		if fn(a) {
			return append(b, a)
		} else {
			return b
		}
	}, []A{}, src)
}

func Compose[A any, B any, C any](fnB func(B) C, fnA func(A) B) func(A) C {
	return func(a A) C {
		return fnB(fnA(a))
	}
}

func Foldl[A any, B any](fn func(B, A) B, acc B, src []A) B {
	return Guards(
		Guard(len(src) == 0, func() B { return acc }),
		Guard(true, func() B { return Foldl(fn, fn(acc, Head(src)), Tail(src)) }),
	)
}

func Foldr[A any, B any](fn func(A, B) B, acc B, src []A) B {
	return Guards(
		Guard(len(src) == 0, func() B { return acc }),
		Guard(len(src) == 1, func() B { return fn(Head(src), acc) }),
		Guard(true, func() B { return fn(Head(src), Foldr(fn, acc, Tail(src))) }),
	)
}

func ZipWith[A any, B any, C any](fn func(A, B) C, srcA []A, srcB []B) (result []C) {
	return Guards(
		Guard(len(srcA) == 0 || len(srcB) == 0, func() []C { return nil }),
		Guard(true, func() []C {
			return append(append(result, fn(Head(srcA), Head(srcB))), ZipWith(fn, Tail(srcA), Tail(srcB))...)
		}),
	)
}

func Zip[A any, B any](srcA []A, srcB []B) (result []Tuple[A, B]) {
	return ZipWith(func(a A, b B) Tuple[A, B] {
		return Tuple[A, B]{fst: a, snd: b}
	}, srcA, srcB)
}

func Fst[A any, B any](t Tuple[A, B]) A {
	return t.fst
}

func Snd[A any, B any](t Tuple[A, B]) B {
	return t.snd
}

func Last[A any](src []A) A {
	return Guards(
		Guard(len(src) == 0, func() A { panic("cant get last element of empty slice") }),
		Guard(true, func() A { return src[len(src)-1] }),
	)
}

func Any[A any](fn func(A) bool, src []A) bool {
	return Guards(
		Guard(len(src) == 0, func() bool { return false }),
		Guard(len(src) > 0 && fn(Head(src)), func() bool { return true }),
		Guard(len(src) > 0, func() bool { return Any(fn, Tail(src)) }),
		Guard(true, func() bool { return false }), // Fallback, sollte nie erreicht werden
	)
}

func All[A any](fn func(A) bool, src []A) bool {
	return Guards(
		Guard(len(src) == 0, func() bool { return true }),
		Guard(len(src) > 0 && !fn(Head(src)), func() bool { return false }),
		Guard(len(src) > 0, func() bool { return All(fn, Tail(src)) }),
		Guard(true, func() bool { return true }), // Fallback, sollte nie erreicht werden
	)
}

func Sum[A int8 | int16 | int32 | int64 | int | float32 | float64](src []A) A {
	return Foldr(func(acc A, x A) A {
		return acc + x
	}, A(0), src)
}

func Product[A int8 | int16 | int32 | int64 | int | float32 | float64](src []A) A {
	return Foldr(func(acc A, x A) A {
		return acc * x
	}, A(0), src)
}

func FlattenWith[A comparable, B any, C any](fn func(A, B) C, src map[A]B) (result []C) {
	for k, v := range src {
		result = append(result, fn(k, v))
	}
	return result
}

func Flatten[A comparable, B any](src map[A]B) (result []Tuple[A, B]) {
	return FlattenWith(func(a A, b B) Tuple[A, B] {
		return Tuple[A, B]{fst: a, snd: b}
	}, src)
}

func Maximum[A comparable](src []A) A {
	return Guards(
		Guard(len(src) == 0, func() A { panic("called maximum on empty list") }),
		Guard(true, func() A {
			return Foldl(func(acc A, x A) A {
				if Compare(acc, x) == GT {
					return acc
				} else {
					return x
				}
			}, src[0], src)
		}),
	)

}

func Minimum[A comparable](src []A) A {
	return Guards(
		Guard(len(src) == 0, func() A { panic("called minimum on empty list") }),
		Guard(true, func() A {
			return Foldl(func(acc A, x A) A {
				if Compare(acc, x) == LT {
					return acc
				} else {
					return x
				}
			}, src[0], src)
		}),
	)
}
