package main

import (
	"fmt"

	"github.com/samber/lo"
)

type Item struct {
	Name string
	Type string
}

func main() {
	// 奇数のみを抽出
	results := lo.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(x int, _ int) bool {
		return x%2 != 0
	})
	fmt.Println(results) // [1, 3, 5, 7, 9]

	// オブジェクトの抽出
	items := []Item{{Name: "item1", Type: "foo"}, {Name: "item2", Type: "bar"}, {Name: "item3", Type: "foo"}}
	filteredItem := lo.Filter(items, func(e Item, _ int) bool { // 第二引数の関数の二番目の引数はindexです
		return e.Type == "foo"
	})
	fmt.Println(filteredItem) // [{item1 foo} {item3 foo}]
}
