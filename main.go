package main

import (
	"fmt"

	"github.com/Aslanchik2o/cache"
)

func main() {
	cache := cache.New()
	cache.Set("userId", 17)

	userId := cache.Get("userId")

	fmt.Println(userId)


	cache.Delete("userId")

	userId = cache.Get("userId")

	fmt.Println(userId)

}