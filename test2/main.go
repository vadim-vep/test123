package main

type String string

func (s *String) Set(key string) {
	*s = String(key)
}

type settable[T any] interface {
	*T
	Set(string)
}

func calSettable[T any, S settable[T]]() T {
	var v T

	S(&v).Set("Hello world!")

	return v
}

func main() {
	s := calSettable[String]()
	//   ^^^^^^^^^^^^^^^^^^^^^
	//   Here it shows "Cannot use *T as the type settable[T]"
	println(s)
}
sdfdsf