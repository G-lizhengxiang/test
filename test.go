package main

import "fmt"

type t_classmates map[string]int

func main() {

	domitory := make(map[string]t_classmates)
	class1 := make(t_classmates)
	class1["zhangsan"] = 23
	class1["lisi"] = 24

	domitory["308"] = class1
	domitory["309"] = t_classmates{"wangwu": 25, "zhaoliu": 26}
	fmt.Println(domitory)
}
