package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("doSomething")
	fmt.Println(ctx.Value("key"))
}

func tryContext() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")

	doSomething(ctx)
}
