package xgo

import (
	"fmt"
	"sync"
)

// If returns the second argument if the first argument is true, otherwise returns the third argument.
func If[T any](f bool, a, b T) T {
	if f {
		return a
	}
	return b
}

// OK panics if argument is not null error.
func OK(err error) {
	if err != nil {
		panic(err)
	}
}

// NoErr panics if argument is not null error.
func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Val returns argument and panics on error.
func Val[T any](v T, err error) T {
	OK(err)
	return v
}

// Val2 returns arguments and panics on error.
func Val2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	OK(err)
	return v1, v2
}

// Val3 returns arguments and panics on error.
func Val3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	OK(err)
	return v1, v2, v3
}

// SafeVal returns argument, ignores error.
func SafeVal[T any](v T, err error) T {
	// ignore error
	return v
}

// SafeVal2 returns argument2, ignores error.
func SafeVal2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	// ignore error
	return v1, v2
}

// SafeVal3 returns argument2, ignores error.
func SafeVal3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	// ignore error
	return v1, v2, v3
}

// Require panics if the statement is false.
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

// In returns true if the value is included in the list of values.
func In[T comparable](v T, values ...T) bool {
	for _, v2 := range values {
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

// FilterFn returns v if it is present by filter function.
func FilterFn[T any](v T, filter func(T) bool) (_ T) {
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

// ExcludeFn returns v if it is not present by exclude function.
func ExcludeFn[T any](v T, exclude func(T) bool) (_ T) {
	if exclude(v) {
		return
	}
	return v
}
