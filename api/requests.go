package api

import (
	"context"
	"encoding/json"

	"github.com/machinebox/graphql"
)

type TopicTag struct {
	Name string
	Id   string
	Slug string
}

type QuestionInfo struct {
	AcRate             float32
	Difficulty         string
	FreqBar            interface{}
	FrontendQuestionId string
	IsFavor            bool
	PaidOnly           bool
	Status             string
	Title              string
	TitleSlug          string
	TopicTags          []TopicTag
	HasSolution        bool
	HasVideoSolution   bool
}

type ProblemsetQuestionList struct {
	Total     int32
	Questions []QuestionInfo
}

func GetProblemsetQuestionList(categorySlug string, skip int, limit int, filters struct{}) (ProblemsetQuestionList, error) {
	list := ProblemsetQuestionList{}

	req := graphql.NewRequest(`
	query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
		problemsetQuestionList: questionList(
			categorySlug: $categorySlug
			limit: $limit
			skip: $skip
			filters: $filters
		) {
			total: totalNum
			questions: data {
				acRate
				difficulty
				freqBar
				frontendQuestionId: questionFrontendId
				isFavor
				paidOnly: isPaidOnly
				status
				title
				titleSlug
				topicTags {
					name
					id
					slug
				}
				hasSolution
				hasVideoSolution
			}
		}
	}
`)

	req.Var("categorySlug", categorySlug)
	req.Var("skip", skip)
	req.Var("limit", limit)
	req.Var("filters", filters)

	ctx := context.Background()

	var data map[string]interface{}
	if err := GetClient().Run(ctx, req, &data); err != nil {
		return list, err
	}

	jsonBody, err := json.Marshal(data["problemsetQuestionList"])
	if err != nil {
		return list, err
	}

	if err := json.Unmarshal(jsonBody, &list); err != nil {
		return list, err
	}

	return list, nil
}
