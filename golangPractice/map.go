package main

import "log"

func mapTest() {
	m := make(map[string]string)

	m["Go"] = "Beego"
	m["Python"] = "Django"
	m["Ruby"] = "Rails"
	m["PHP"] = "Laravel"

	// Call the value by the key.
	if !(m["Go"] == "Beego") {
		log.Fatal("Wrong value")
	}
}
