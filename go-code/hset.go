package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) ToJson() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func WorkHSET() {
	rdb := GetClient()
	//n, err := rdb.HSet(context.TODO(), "sample", "k1", "v1", "k2", "v2", "k3", "v3").Result()
	//fmt.Println(n, err)
	//v1, err := rdb.HGetAll(context.Background(), "sample").Result()
	//fmt.Println(v1, err)
	//p := Person{Name: "Ichigo Hoshimiya", Age: 22}
	//_, err := rdb.HSet(context.Background(), "sample", "ichigo", p.ToJson()).Result()
	//fmt.Println(err)
	s, _ := rdb.HGet(context.TODO(), "sample", "ichigo").Result()
	var p Person
	json.Unmarshal([]byte(s), &p)
	fmt.Println(p)
}
