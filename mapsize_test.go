package mapsize

import (
	"fmt"
	_ "net/http/pprof"
)

func ExampleSize() {
	m := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		m[i] = struct{}{}
	}
	fmt.Println(Size(m)) // 1416
}
