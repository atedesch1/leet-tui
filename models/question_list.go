package models

import (
	"fmt"

	"github.com/atedesch1/leet-tui/api"
	tea "github.com/charmbracelet/bubbletea"
)

type QuestionList view

type QuestionListModel struct {
	cursor       int
	questions    []api.QuestionInfo
	skip         int
	limit        int
	categorySlug string

	selectedQuestion api.QuestionInfo
}

func newQuestionListModel() *QuestionListModel {
	return &QuestionListModel{
		cursor:       0,
		questions:    make([]api.QuestionInfo, 0),
		skip:         0,
		limit:        10,
		categorySlug: "",
	}
}

func (m *QuestionListModel) getQuestionInfoList() tea.Msg {
	questions, _ := api.GetQuestionInfoList(m.categorySlug, m.skip, m.limit)
	return questions
}

type SelectedQuestionMsg struct {
	question api.QuestionInfo
}

func (m *QuestionListModel) selectQuestionCmd() tea.Msg {
	return SelectedQuestionMsg{
		question: m.selectedQuestion,
	}
}

func (m *QuestionListModel) Init() tea.Cmd {
	return m.getQuestionInfoList
}

func (m *QuestionListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case []api.QuestionInfo:
		m.questions = msg

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.questions)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selectedQuestion = m.questions[m.cursor]
			return m, m.selectQuestionCmd
		}

	}
	return m, nil
}

func (m *QuestionListModel) View() string {
	s := "Choose the question!\n\n"

	for i, choice := range m.questions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}
	return s
}
