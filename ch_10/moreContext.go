package main

import (
	"context"
	"fmt"
)

type aKey string

func searchKey(ctx context.Context, k aKey) {
	// Извлечь значение из контекста.
	v := ctx.Value(k)
	// Проверить, существует ли значение.
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("mySecretValue")
	// context.WithValue() предоставляет способ связывания значения с объектом Context.
	ctx := context.WithValue(context.Background(), myKey, "mySecretValue")
	searchKey(ctx, myKey)

	searchKey(ctx, aKey("notThere"))
	// В данном случае мы заявляем, что, несмотря на то что мы намерены использовать
	// контекст операции, мы все еще не уверены в этом и потому используем функцию
	// context.TODO()
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}
