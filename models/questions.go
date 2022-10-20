package models

type TopicTag struct {
	name string
	id   string
	slug string
}

type Question struct {
	acRate             float32
	difficulty         string
	freqBar            interface{}
	frontendQuestionId string
	isFavor            bool
	paidOnly           bool
	status             string
	title              string
	titleSlug          string
	topicTags          []TopicTag
	hasSolution        bool
	hasVideoSolution   bool
}
