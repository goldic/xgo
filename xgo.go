package xgo

import (
	"fmt"
	"sync"
)

// If returns a when f is true, otherwise returns b.
func If[T any](f bool, a, b T) T {
	if f {
		return a
	}
	return b
}

// OK panics if err is not null.
func OK(err error) {
	if err != nil {
		panic(err)
	}
}

// NoErr panics if err is not null. Synonym of OK(err)
func NoErr(err error) {
	OK(err)
}

// Val returns v or panics if err is not null.
func Val[T any](v T, err error) T {
	OK(err)
	return v
}

// Val2 returns v1, v2 or panics if err is not null.
func Val2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	OK(err)
	return v1, v2
}

// Val3 returns v1, v2, v3 or panics if err is not null.
func Val3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	OK(err)
	return v1, v2, v3
}

// SafeVal returns v and ignores error.
func SafeVal[T any](v T, err error) T {
	// ignore error
	return v
}

// SafeVal2 returns v1, v2 and ignores error.
func SafeVal2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	// ignore error
	return v1, v2
}

// SafeVal3 returns v1, v2, v3 and ignores error.
func SafeVal3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	// ignore error
	return v1, v2, v3
}

// Require panics if statement is false.
func Require(statement bool, err any) {
	if !statement {
		if _, ok := err.(error); ok {
			panic(err)
		}
		panic(fmt.Errorf("%v", err))
	}
}

// Recover recovers and returns error by argument pointer.
func Recover(err *error) {
	if r := recover(); r != nil && err != nil && *err == nil {
		if e, ok := r.(error); ok {
			*err = e
		} else {
			*err = fmt.Errorf("%v", r)
		}
	}
}

// Mute mutes panic-error.
func Mute() {
	recover()
}

// Call runs the function safely, recovers panic-error.
func Call(fn func()) (err error) {
	defer Recover(&err)
	fn()
	return
}

// Go runs the function safely.
func Go(fn func()) {
	go Call(fn)
}

// Async asynchronously runs several functions and waits for them to complete, returns an error in case of panic.
func Async(fn ...func()) (err error) {
	var wg sync.WaitGroup
	wg.Add(len(fn))
	for _, f := range fn {
		go func(fn func()) {
			defer wg.Done()
			defer Recover(&err)
			fn()
		}(f)
	}
	wg.Wait()
	return
}

// In reports whether v is present in ...value.
func In[T comparable](v T, value ...T) bool {
	for _, v2 := range value {
		if v == v2 {
			return true
		}
	}
	return false
}

// Or returns the first non-empty value.
func Or[T comparable](values ...T) (v0 T) {
	for _, v := range values {
		if v != v0 {
			return v
		}
	}
	return
}

// FilterFunc returns v if v satisfies filter(v).
func FilterFunc[T any](v T, filter func(T) bool) (_ T) {
	if filter(v) {
		return v
	}
	return
}

// Exclude returns v if it is not present in vv.
func Exclude[T comparable](v T, vv ...T) (result T) {
	if In(v, vv...) {
		return
	}
	return v
}

// ExcludeFunc returns v if v doesnt satisfies filter(v).
func ExcludeFunc[T any](v T, exclude func(T) bool) (_ T) {
	if exclude(v) {
		return
	}
	return v
}
