package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Price int `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{Title:"喜剧之王", Year: 2012, Price: 37, Actors: []string{"Jay Zhou", "张柏芝"}}
	// 将结构体转为 json string
	movieStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("marshal json failed")
		return
	}

	//fmt.Println(movieStr)
	fmt.Printf("json stirng: %s \n", movieStr) //  {"title":"喜剧之王","year":2012,"price":37,"actors":["Jay Zhou","张柏芝"]}

	var v Movie
	err = json.Unmarshal(movieStr, &v)
	if err != nil {
		fmt.Println("unmarshal json failed")
		return
	}

	fmt.Println(v)
}
