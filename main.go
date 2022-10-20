package main

import (
	"fmt"

	"github.com/atedesch1/leet-tui/api"
)

func main() {
	api.Init()
	list, _ := api.GetProblemsetQuestionList("", 0, 50, struct{}{})

	fmt.Println(list.Questions[1])
}
