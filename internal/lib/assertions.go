package lib

import "reflect"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustBeTrue(b bool) {
	if !b {
		panic("false is not true")
	}
}

func MustBeEqual[T comparable](t1, t2 T) {
	if !reflect.DeepEqual(t1, t2) {
		Panicf("%v != %v", t1, t2)
	}
}
