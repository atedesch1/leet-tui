package models

import (
	"github.com/atedesch1/leet-tui/api"
	tea "github.com/charmbracelet/bubbletea"
)

type Question view

type QuestionModel struct {
	titleSlug string
	question  api.Question
}

func newQuestionModel() *QuestionModel {
	return &QuestionModel{
		question: api.Question{},
	}
}

func (m *QuestionModel) getQuestion() tea.Msg {
	question, _ := api.GetFullQuestion(m.titleSlug)
	return question
}

func (m *QuestionModel) Init() tea.Cmd {
	return m.getQuestion
}

func (m *QuestionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case SelectedQuestionMsg:
		m.titleSlug = msg.question.TitleSlug
		return m, m.getQuestion

	case api.Question:
		m.question = msg
	}

	return m, nil
}

func (m *QuestionModel) View() string {
	return m.question.TitleSlug
}
