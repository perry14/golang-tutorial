package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	kelu := Person{"kelu", 25}
	t := reflect.TypeOf(kelu)
	n := t.NumField()
	for i := 0; i < n; i++ {
		fmt.Println(t.Field(i).Name)
		fmt.Println(t.Field(i).PkgPath)
		fmt.Println(t.Field(i).Type)
		fmt.Println(t.Field(i).Tag)
		fmt.Println(t.Field(i).Offset)
		fmt.Println(t.Field(i).Index)
		fmt.Println(t.Field(i).Anonymous)
	}
}
