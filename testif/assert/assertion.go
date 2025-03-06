package assert

type TestingT interface {
	Errorf(format string, args ...interface{})
}

// ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
// for table driven tests.
type ComparisonAssertionFunc func(TestingT, interface{}, interface{}, ...interface{}) bool

// ValueAssertionFunc is a common function prototype when validating a single value.  Can be useful
// for table driven tests.
type ValueAssertionFunc func(TestingT, interface{}, ...interface{}) bool

// BoolAssertionFunc is a common function prototype when validating a bool value.  Can be useful
// for table driven tests.
type BoolAssertionFunc func(TestingT, bool, ...interface{}) bool

// ErrorAssertionFunc is a common function prototype when validating an error value.  Can be useful
// for table driven tests.
type ErrorAssertionFunc func(TestingT, error, ...interface{}) bool

// PanicAssertionFunc is a common function prototype when validating a panic value.  Can be useful
// for table driven tests.
type PanicAssertionFunc = func(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool

// Comparison is a custom function that returns true on success and false on failure
type Comparison func() (success bool)

type tHelper = interface {
	Helper()
}
