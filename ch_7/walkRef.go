package main

import (
	"fmt"
	"github.com/mitchellh/reflectwalk"
	"reflect"
)

type Values struct {
	Extra map[string]string
}

type WalkMap struct {
	MapVal reflect.Value
	Keys   map[string]bool
	Values map[string]bool
}

func (t *WalkMap) Map(m reflect.Value) error {
	t.MapVal = m
	return nil
}

func (t *WalkMap) MapElem(m, k, v reflect.Value) error {
	if t.Keys == nil {
		t.Keys = make(map[string]bool)
		t.Values = make(map[string]bool)
	}

	t.Keys[k.Interface().(string)] = true
	t.Values[v.Interface().(string)] = true
	return nil
}

func main() {
	w := new(WalkMap)

	type S struct {
		Map map[string]string
	}
	// Определить новую переменную data, которая содержит хэш-таблицу
	data := &S{
		Map: map[string]string{
			"V1": "v1v",
			"V2": "v2v",
			"V3": "v3v",
			"V4": "v4v",
		},
	}

	// Вызвать функцию reflectwalk.Walk чтобы исследовать переменную data
	err := reflectwalk.Walk(data, w)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := w.MapVal
	fmt.Println("MapVal:", r)
	rType := r.Type()
	fmt.Printf("Type of r: %s\n", rType)

	// Метод MapKeys возвращает срез reflect.Values, каждое значение которого
	// содержит один из ключей хэш-таблицы.
	for _, key := range r.MapKeys() {
		// Метод MapIndex позволяет вывести на экран значение ключа
		fmt.Println("key:", key, "value:", r.MapIndex(key))
	}

	// MapVal: map[V1:v1v V2:v2v V3:v3v V4:v4v]
	// Type of r: map[string]string
	// key: V3 value: v3v
	// key: V4 value: v4v
	// key: V1 value: v1v
	// key: V2 value: v2v
}
