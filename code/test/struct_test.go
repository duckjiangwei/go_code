package structTest

import "testing"

type User struct{}

func TestStruct(t *testing.T) {
	s := User{}
	t.Log(s)
}
