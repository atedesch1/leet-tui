package main

import (
	"fmt"

	"github.com/atedesch1/leet-tui/api"
)

func main() {
	api.Init()
	list, _ := api.GetQuestionInfoList("", 0, 50, struct{}{})
	question, err := api.GetFullQuestion(list[0].TitleSlug)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(list[0].TitleSlug)
	fmt.Println(question.Stats)
}
