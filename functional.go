package functionalgo

type Tuple[A any, B any] struct {
	fst A
	snd B
}

func Take[A ~[]any](src A, num int) A {
	if num > len(src) {
		return src
	}
	return src[:num]
}

func Drop[A ~[]any](src A, num int) A {
	if num > len(src) {
		return A{}
	}
	return src[num+1:]
}
func Head[A any, B ~[]A](src B) A {
	if len(src) == 0 {
		panic("cannot take head of empty list")
	}
	return src[0]
}

func Tail[A any, B ~[]A](src B) B {
	if len(src) == 0 {
		panic("cannot take tail of empty list")
	}
	return src[1:]
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
	if len(src) == 0 {
		return acc
	}
	return Foldl(fn, fn(acc, Head(src)), Tail(src))
}

func Foldr[A any, B any](fn func(A, B) B, acc B, src []A) B {
	if len(src) == 1 {
		return fn(Head(src), acc)
	}
	return fn(Head(src), Foldr(fn, acc, Tail(src)))
}

func ZipWith[A any, B any, C any](fn func(A, B) C, srcA []A, srcB []B) (result []C) {
	if len(srcA) == 0 || len(srcB) == 0 {
		return nil
	}
	return append(append(result, fn(Head(srcA), Head(srcB))), ZipWith(fn, Tail(srcA), Tail(srcB))...)
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
	if len(src) == 0 {
		panic("cant get last element of empty slice")
	}
	return src[len(src)-1]
}

func Any[A any](src []A, fn func(A) bool) bool {
	if len(src) == 0 {
		return false
	}
	if fn(Head(src)) {
		return true
	}
	return Any(Tail(src), fn)
}

func All[A any](src []A, fn func(A) bool) bool {
	if len(src) == 0 {
		return true
	}
	if !fn(Head(src)) {
		return false
	}
	return All(Tail(src), fn)
}
