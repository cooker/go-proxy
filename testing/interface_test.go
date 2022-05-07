package testing

import "testing"

type Say interface {
	Hello()
}

type A struct {
}

func (*A) Hello() {
	println("sasasa")
}

type B struct {
	Say
}

func (*B) Hello() {
	println("bbbb")
}

func TestA(t *testing.T) {
	var a = A{}
	a.Hello()
	var b = B{}
	b.Hello()

	var say Say = new(A)
	say.Hello()
}
