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

type LeetQuestionInfo struct {
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

type QuestionInfo struct {
	AcRate     float32
	Difficulty string
	Status     string
	Title      string
	TitleSlug  string
	TopicTags  []TopicTag
}

type LeetProblemsetQuestionList struct {
	Total     int
	Questions []LeetQuestionInfo
}

type ProblemsetQuestionList struct {
	Total     int
	Questions []QuestionInfo
}

func GetProblemsetQuestionList(categorySlug string, skip int, limit int) (ProblemsetQuestionList, error) {
	var qlist ProblemsetQuestionList

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
	req.Var("filters", struct{}{})

	var data map[string]interface{}
	if err := getClient().Run(context.Background(), req, &data); err != nil {
		return qlist, err
	}

	jsonBody, err := json.Marshal(data["problemsetQuestionList"])
	if err != nil {
		return qlist, err
	}

	ll := LeetProblemsetQuestionList{}
	if err := json.Unmarshal(jsonBody, &ll); err != nil {
		return qlist, err
	}

	qlist.Total = ll.Total
	for _, lqi := range ll.Questions {
		qlist.Questions = append(qlist.Questions, QuestionInfo{
			AcRate:     lqi.AcRate,
			Difficulty: lqi.Difficulty,
			Status:     lqi.Status,
			Title:      lqi.Title,
			TitleSlug:  lqi.TitleSlug,
			TopicTags:  lqi.TopicTags,
		})
	}

	return qlist, nil
}

type LeetQuestion struct {
	QuestionId         string
	QuestionFrontendId string
	BoundTopicId       string
	Title              string
	TitleSlug          string
	Content            string
	TranslatedTitle    string
	TranslatedContent  string
	IsPaidOnly         bool
	CanSeeQuestion     bool
	Difficulty         string
	Likes              int
	Dislikes           int
	IsLiked            bool
	SimilarQuestions   string
	// struct {
	// 	Title           string
	// 	TitleSlug       string
	// 	Difficulty      string
	// 	TranslatedTitle string
	// }
	ExampleTestcases string
	CategoryTitle    string
	// contributors {
	// 	username
	// 	profileUrl
	// 	avatarUrl
	// 	__typename
	// }
	TopicTags []TopicTag
	// companyTagStats
	CodeSnippets []struct {
		Lang     string
		LangSlug string
		Code     string
	}
	Stats    string
	Hints    []string
	Solution struct {
		Id           string
		CanSeeDetail bool
		// PaidOnly bool
		// HasVideoSolution bool
		// PaidOnlyVideo bool
	}
	Status         string
	SampleTestCase string
	MetaData       string
	// JudgerAvailable bool
	// JudgeType string
	// MysqlSchemas
	// EnableRunCode
	// EnableTestMode
	// EnableDebugger
	// EnvInfo
	// LibraryUrl string
	// AdminUrl string
	// ChallengeQuestion struct {
	// 	Id int64
	// 	Date string
	// 	IncompleteChallengeCount
	// 	StreakCount
	// }
}

type Question struct {
	QuestionId       string
	Title            string
	TitleSlug        string
	Content          string
	Difficulty       string
	Likes            int
	Dislikes         int
	IsLiked          bool
	SimilarQuestions []struct {
		Title      string
		TitleSlug  string
		Difficulty string
	}
	ExampleTestcases string
	CategoryTitle    string
	TopicTags        []TopicTag
	CodeSnippets     []struct {
		Lang     string
		LangSlug string
		Code     string
	}
	Stats struct {
		TotalAccepted      string
		TotalSubmission    string
		TotalAcceptedRaw   int
		TotalSubmissionRaw int
		AcRate             string
	}
	Hints    []string
	Solution struct {
		Id           string
		CanSeeDetail bool
	}
	Status         string
	SampleTestCase string
	// MetaData       string
}

func GetFullQuestion(titleSlug string) (Question, error) {
	question := Question{}

	req := graphql.NewRequest(`
	query questionData($titleSlug: String!) {
		question(titleSlug: $titleSlug) {
			questionId
			questionFrontendId
			boundTopicId
			title
			titleSlug
			content
			translatedTitle
			translatedContent
			isPaidOnly
			canSeeQuestion
			difficulty
			likes
			dislikes
			isLiked
			similarQuestions
			exampleTestcases
			categoryTitle
			contributors {
				username
				profileUrl
				avatarUrl
				__typename
			}
			topicTags {
				name
				slug
				translatedName
				__typename
			}
			companyTagStats
			codeSnippets {
				lang
				langSlug
				code
				__typename
			}
			stats
			hints
			solution {
				id
				canSeeDetail
				paidOnly
				hasVideoSolution
				paidOnlyVideo
				__typename
			}
			status
			sampleTestCase
			metaData
			judgerAvailable
			judgeType
			mysqlSchemas
			enableRunCode
			enableTestMode
			enableDebugger
			envInfo
			libraryUrl
			adminUrl
			challengeQuestion {
				id
				date
				incompleteChallengeCount
				streakCount
				type
				__typename
			}
			__typename
		}
	}	
`)

	req.Var("titleSlug", titleSlug)

	var data map[string]interface{}
	if err := getClient().Run(context.Background(), req, &data); err != nil {
		return question, err
	}

	jsonBody, err := json.Marshal(data["question"])
	if err != nil {
		return question, err
	}

	lq := LeetQuestion{}

	if err := json.Unmarshal(jsonBody, &lq); err != nil {
		return question, err
	}

	var similarQuestions []struct {
		Title      string
		TitleSlug  string
		Difficulty string
	}

	if err := json.Unmarshal([]byte(lq.SimilarQuestions), &similarQuestions); err != nil {
		return question, err
	}

	var stats struct {
		TotalAccepted      string
		TotalSubmission    string
		TotalAcceptedRaw   int
		TotalSubmissionRaw int
		AcRate             string
	}

	if err := json.Unmarshal([]byte(lq.Stats), &stats); err != nil {
		return question, err
	}

	question = Question{
		QuestionId:       lq.QuestionId,
		Title:            lq.Title,
		TitleSlug:        lq.TitleSlug,
		Content:          lq.Content,
		Difficulty:       lq.Difficulty,
		Likes:            lq.Likes,
		Dislikes:         lq.Dislikes,
		IsLiked:          lq.IsLiked,
		SimilarQuestions: similarQuestions,
		ExampleTestcases: lq.ExampleTestcases,
		CategoryTitle:    lq.CategoryTitle,
		TopicTags:        lq.TopicTags,
		CodeSnippets:     lq.CodeSnippets,
		Stats:            stats,
		Hints:            lq.Hints,
		Solution:         lq.Solution,
		Status:           lq.Status,
		SampleTestCase:   lq.SampleTestCase,
	}

	return question, nil
}
