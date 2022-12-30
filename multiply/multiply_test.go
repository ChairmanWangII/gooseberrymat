package multiply

import (
	"fmt"
	"reflect"
	"testing"
)

type Num interface {
	int32 | int64
}

func add[N Num](a, b N) N {
	fmt.Println(reflect.TypeOf(a).Name())
	return a + b
}
func TestReflect(t *testing.T) {
	add[int32](12, 21)
}
