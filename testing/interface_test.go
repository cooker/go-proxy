package testing

import (
	"reflect"
	"testing"
)

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
	var a = new(A)
	a.Hello()
	var b = new(B)
	b.Hello()

	b2, _ := interface{}(b).(Say)
	b2.Hello()

	var say Say = new(A)
	say.Hello()
	(say).(Say).Hello()
	elem := reflect.TypeOf((*Say)(nil)).Elem()
	println(reflect.TypeOf(b).Implements(elem))
}

func TestAfield(t *testing.T) {
	//var a = A{}
	//value := reflect.ValueOf(a)

}
