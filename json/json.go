package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var p = fmt.Println

func main() {
	bolB, _ := json.Marshal(true)
	p(string(bolB))

	intB, _ := json.Marshal(1)
	p(string(intB))

	fltB, _ := json.Marshal(2.34)
	p(string(fltB))

	strB, _ := json.Marshal("welcome")
	p(string(strB))

	slcD := []string{"applie", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	p(string(slcB))

	mapD := map[string]int{"apple": 1, "lettuce": 5}
	mapB, _ := json.Marshal(mapD)
	p(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "pear", "peach"},
	}
	res1B, _ := json.Marshal(res1D)
	p(string(res1B))

	res2D := &response2{
		Page:   2,
		Fruits: []string{"grape", "tomato", "cherry"},
	}
	res2B, _ := json.Marshal(res2D)
	p(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
