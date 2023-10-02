package cache

import (
	 
	"fmt"
)

func main() {
	cache := cache.New
	cache.Set("userId", 42)
	userId := cache.Get("userId")
	

	fmt.Println(userId)
	cache.Deleate("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}