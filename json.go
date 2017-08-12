package main

import "fmt"
import "encoding/json"

type response1 struct {
	Id    int //The json package only accesses the exported fields of struct types (those that begin with an uppercase letter). Therefore only the the exported fields of a struct will be present in the JSON output.
	Names []string
}

type response2 struct {
	Id    int      `json:"idNew"` // this will be used when marshel/unmarshal
	Names []string `json:"nameNew"`
}

func main() {
	b, _ := json.Marshal("hello") // ignore error as the second return value
	fmt.Println(b)
	fmt.Println(string(b)) // "hello" with quotes

	s, _ := json.Marshal(2.467)
	fmt.Println(string(s)) // 2.467

	a, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println(string(a))

	m, _ := json.Marshal(map[string]int{"a": 1, "b": 2, "c": 3})
	fmt.Println(string(m)) // will not print the map, just bunch of ascii values

	r1 := response1{Id: 1, Names: []string{"a", "n", "m"}}
	fmt.Println(r1)
	r2, _ := json.Marshal(r1)
	fmt.Println(string(r2))

	r3, _ := json.Marshal(response2{Id: 2, Names: []string{"b", "y", "p"}})
	fmt.Println(string(r3))

	byt := []byte(`{
		"num":5.6,"strs":["a","b"]
	}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"].(float64))
	fmt.Println(dat["strs"].([]interface{})[0].(string))

	str := `{"idNew" :2, "nameNew":["r","t","s"]}`
	res2 := response2{}
	json.Unmarshal([]byte(str), &res2)
	fmt.Println(res2)
	fmt.Println(res2.Names[2])
}
