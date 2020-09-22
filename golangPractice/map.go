package main

import (
	"fmt"
	"log"
)

func mapTest() {
	m := make(map[string]string)

	m["Go"] = "Beego"
	m["Python"] = "Django"
	m["Ruby"] = "Rails"
	m["PHP"] = "Laravel"

	//檢查key/value是否存在
	v, ok := m["Java"]
	fmt.Println(v)
	fmt.Println(ok)

	//移除
	delete(m, "PHP")
	fmt.Println(m)

	//遍歷map
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%s: %s", k, v))
	}

	if !ok {
		log.Fatal("It should be true")
	}

	if !(v == "Beego") {
		log.Fatal("Wrong value")
	}

	_, ok = m["Java"]
	if !(ok == false) {
		log.Fatal("It should be false")
	}

}
