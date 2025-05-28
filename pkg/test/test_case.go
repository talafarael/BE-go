package test_case

type TestCase[T any, Y any] struct {
	Name    string
	Data    T
	Res     Y
	Err     error
	WantErr bool
}
