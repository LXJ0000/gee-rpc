package main

import (
	"fmt"
	"reflect"
	"testing"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Println(reflect.ValueOf(v))
}

func TestReflect(t *testing.T) {

	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}
