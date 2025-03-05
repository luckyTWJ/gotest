package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "电影",
		Year:   2019,
		Price:  100,
		Actors: []string{"tom", "jack"},
	}
	//编码json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json err", err)
		return
	}
	fmt.Printf("json str->%s\n", jsonStr)
	//解码json 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json err", err)
		return
	}
	fmt.Printf("myMovie->%+v\n", myMovie)
	fmt.Printf("myMovie->%v\n", myMovie)

}
