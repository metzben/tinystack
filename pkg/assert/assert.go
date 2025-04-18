package assert

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/exp/constraints"
)

func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}

func NotEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected == actual {
		t.Errorf("Received %v (type %v), expected value other than %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}

func NotNil(t *testing.T, actual any) {
	t.Helper()
	if actual == nil {
		t.Error("Expected a non-nil value, but got nil")
	}
}

func Nil(t *testing.T, actual any) {
	t.Helper()
	if actual != nil {
		t.Errorf("Expected nil, but got %v (type %v)", actual, reflect.TypeOf(actual))
	}
}

func DeepEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}

// ErrorContains checks if the error message in out contains the text in
// want. It returns true if it does, false otherwise.
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

// Between checks if a value is within the specified range (inclusive of min and max)
func Between[T constraints.Ordered](t *testing.T, min, max, actual T) {
	t.Helper()
	if actual < min || actual > max {
		t.Errorf("Expected value between %v and %v (inclusive), but got %v", min, max, actual)
	}
}
