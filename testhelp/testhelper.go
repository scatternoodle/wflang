package testhelp

// TH allows us to share helper functions between *testing.T and other testing objects
// such as *testing.B
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
		t.Fatalf("type = %t, want %T", x, *new(T))
	}
	return y
}
