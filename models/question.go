package models

import (
	"github.com/atedesch1/leet-tui/api"
	tea "github.com/charmbracelet/bubbletea"
)

type questionModel struct {
	titleSlug string
	question  api.Question
}

func newQuestionModel() *questionModel {
	return &questionModel{
		question: api.Question{},
	}
}

func (m *questionModel) getQuestion() tea.Msg {
	question, _ := api.GetFullQuestion(m.titleSlug)
	return question
}

func (m *questionModel) Init() tea.Cmd {
	return m.getQuestion
}

func (m *questionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case SelectedQuestionMsg:
		m.titleSlug = msg.question.TitleSlug
		return m, m.getQuestion

	case api.Question:
		m.question = msg
	}

	return m, nil
}

func (m *questionModel) View() string {
	return m.question.TitleSlug
}
