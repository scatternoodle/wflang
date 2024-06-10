package testhelper

type TH interface {
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Error(args ...any)
	Fatal(args ...any)
	FailNow()
	Helper()
}
