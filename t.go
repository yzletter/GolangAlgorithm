package main

import (
	"fmt"
	"github.com/yzletter/go-kit/setx"
)

func main() {
	s := setx.NewSet[string](10)
	for i := 0; i < 10; i++ {
		ss := fmt.Sprintf("%d", i)
		fmt.Println(ss)
		s.Insert(ss)
	}
	arr := s.Keys()
	fmt.Println(len(arr))
}
