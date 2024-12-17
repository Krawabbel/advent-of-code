package lib

func Abs[T Number](t T) T {
	if t > 0 {
		return t
	}
	return -t
}

func Min[T Number](ts ...T) T {
	t := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] < t {
			t = ts[i]
		}
	}
	return t
}

func Max[T Number](ts ...T) T {
	t := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] > t {
			t = ts[i]
		}
	}
	return t
}

func Transform[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}
	return us
}

func IntPow(base, exponent int) int {
	if exponent < 0 {
		panic("negative exponent")
	}
	return powImpl(base, exponent)
}

func powImpl(base, exponent int) int {

	if exponent == 0 {
		return 1
	} else if exponent == 1 {
		return base
	}

	if (exponent % 2) == 0 {
		sqrt := powImpl(base, exponent/2)
		return sqrt * sqrt
	}

	return base * powImpl(base, exponent-1)
}
