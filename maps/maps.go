package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println("original map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")

	fmt.Println("map: ", m)

	_, present := m["k2"]
	fmt.Println("k2 present? ", present)

	n := map[string]string{"ok": "guys", "welcome": "proba"}
	fmt.Println("new map: ", n)

}
