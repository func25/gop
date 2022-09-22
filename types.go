package gop

func Ptr[T any](i T) *T {
	return &i
}
