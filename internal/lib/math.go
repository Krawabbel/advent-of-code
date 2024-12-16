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
