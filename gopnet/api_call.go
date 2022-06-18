package gopnet

type APICall[T any] struct {
	Method string
	URL    string
	Body   interface{}
	Opts   []RequestOption
}

func (c APICall[T]) Call() (NetRes[T], error) {
	return Call[T](c.Method, c.URL, c.Body, c.Opts...)
}
