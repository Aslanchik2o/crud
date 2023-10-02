package cache

import (
	 
	"fmt"
	"github.com/Aslanchik2o/cache"
)

func main() {
	cache := cache.New
	cache().Set("asas", 17)
	userId := cache
	

	fmt.Println(userId)
	cache().Delete("userId")
	userId := cache

	fmt.Println(userId)
}