package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/quick", func(ctx iris.Context) {
		ctx.Write(parseAndSort([]byte("[" + ctx.URLParam("array") + "]")))
	})
	app.Post("/quick", func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		ctx.Write(parseAndSort(body))
	})
	app.Run(iris.Addr(":80"))
}

func parseAndSort(bytes []byte) []byte {
	var array []int
	json.Unmarshal(bytes, &array)

	b, _ := json.Marshal(map[string][]int{"result": sort(array)})

	return b
}

func sort(list []int) []int {
	innerSort(list, 0, len(list) - 1)

	return list
}

func innerSort(array []int, from int, to int) {
	if from < to {
		pi := partition(array, from, to)

		innerSort(array, from, pi - 1)
		innerSort(array, pi + 1, to)
	}
}

func partition(array []int, from int, to int) int {
	pivot := array[to]
	i := from -1

	for j := from; j <= to - 1; j++ {
		if array[j] < pivot {
			i++
			swap(array, i, j)
		}
	}

	swap(array, i + 1, to)
	return i + 1
}

func swap(array []int, i int, j int) {
	buf := array[i]
	array[i] = array[j]
	array[j] = buf
}