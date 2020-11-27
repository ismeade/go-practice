package basic

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main1() {

	var a map[string]int
	fmt.Println(a == nil)

	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	a["key1"] = 100
	a["key2"] = 200

	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)

	b := map[int]bool{
		1: true,
		2: false,
	}

	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T\n", b)

	// 是否存在
	val, ok := b[3]
	if ok {
		fmt.Println("has key", val)
	} else {
		fmt.Println("no")
	}

	// 遍历
	for k, v := range b {
		fmt.Println(k, v)
	}

	for k := range b {
		fmt.Println(k)
	}

	for _, v := range b {
		fmt.Println(v)
	}

	// 删除
	delete(b, 1)
	fmt.Println(b)

}

func main2() {
	var scoreMap = make(map[string]int, 8)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

func main3() {
	var mapSlice = make([]map[string]int, 8, 8)

	for _, value := range mapSlice {
		value = make(map[string]int, 8)
		value["key-0-1"] = 1
	}

	fmt.Println(mapSlice)

}

func homeWork1(s string) {
	var wordCount = make(map[string]int, 8)
	words := strings.Split(s, " ")
	for _, value := range words {
		count, ok := wordCount[value]
		if ok {
			wordCount[value] = count + 1
		} else {
			wordCount[value] = 1
		}
	}
	fmt.Println(wordCount)
	a := make(map[string]int)
	fmt.Println(a)
}

type Person struct {
}
