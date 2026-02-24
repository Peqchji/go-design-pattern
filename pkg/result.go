package pkg

type Result[T any] struct {
	Result T
	Error  error
}