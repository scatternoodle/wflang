package testhelper

type TH interface {
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Error(args ...any)
	Fatal(args ...any)
	FailNow()
	Helper()
}

// Takes a type parameter T and an interface x. Calls Fatalf if x does not match T,
// or returns x as a T.
func AssertType[T any](t TH, x any) T {
	y, ok := x.(T)
	if !ok {
		t.Fatalf("type = %t, want %s", x, *new(T))
	}
	return y
}
